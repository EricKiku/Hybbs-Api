package controllers

import (
	"Hybbs-API/Pojo"
	"Hybbs-API/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
)

// 客户端结构体
type Client struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
}

// 映射关系
var clientMap map[int64]*Client = make(map[int64]*Client, 0)

var rwLocker sync.RWMutex

func WsController(ctx *gin.Context) {
	sendUId := ctx.DefaultQuery("sendUId", "")
	sendUIdInt64, _ := strconv.ParseInt(sendUId, 10, 64)
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 根据鉴权的方式来处理,如果不想鉴权的就直接返回true,如果需要鉴权就要根据判断来返回true，或者false
			return true
		},
	}

	// 升级Ws
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("ws连接错误")
		ctx.JSON(http.StatusOK, Pojo.Res{Status: "201", Msg: "服务器连接错误"})
		return
	}

	client := &Client{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
	}
	rwLocker.Lock()
	clientMap[sendUIdInt64] = client
	rwLocker.Unlock()

	// 发送数据给客户端
	go senProc(client)
	// 接收客户端的数据
	go recvProc(client)
}

// 将数据推送到管道中
func sendMsg(userId int64, message []byte) {
	fmt.Printf("接受者ID：%v", userId)
	fmt.Printf("内容：%v", message)
	rwLocker.RLock()
	client, isOk := clientMap[userId]
	rwLocker.RUnlock()
	if isOk {
		client.DataQueue <- message
	}
}

// 从管道中获取数据发送出去
func senProc(node *Client) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("发送消息失败")
				return
			}
		}
	}
}

// 接收客户端数据
func recvProc(client *Client) {
	for {
		_, data, err := client.Conn.ReadMessage()
		if err != nil {
			fmt.Println("接收数据失败", err)
			return
		}
		fmt.Printf("接收的客户端数据：")
		fmt.Println(data)
		// 将数据处理转发给对应的人
		dispatch(data)
	}
}

// 分发数据
func dispatch(data []byte) {
	fmt.Println("接收到的数据", string(data))
	// 解析出数据
	var reqMap map[string]interface{}
	err := json.Unmarshal(data, &reqMap)
	var sId, _ = reqMap["sendId"].(float64)
	var rId, _ = reqMap["receId"].(float64)
	sId64 := int64(sId)
	rId64 := int64(rId)
	var msg = reqMap["message"].(string)
	var date = reqMap["date"].(string)
	if err != nil {
		fmt.Println("解析数据失败:", err.Error())
		return
	} else {
		// 把聊天记录存储到数据库
		err := dao.SaveChatListDao(sId64, rId64, msg, date)
		if err != nil {
			return
		} else {
			// 发送数据给对方
			sendMsg(rId64, data)
			// 给自己也发一个
			sendMsg(sId64, data)
		}
	}

}
