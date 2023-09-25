package Tools

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
)

//// MailConfig 邮箱发送配置
//type MailConfig struct {
//	Account string
//	// QQ邮箱填写授权码
//	Password string
//	// QQ：POP/SMTP 587
//	Port int
//	// QQ：smtp.qq.com
//	Host string
//}

// SendMail 发送邮件
// from 发送者别名，mailTo 发送对象，subject主题，body 内容
func SendMail(email string) string {
	message := gomail.NewMessage()

	// 发送人
	message.SetHeader("From", "2966678301@qq.com")
	//接收人
	message.SetHeader("To", email)
	//抄送人

	// 主题
	message.SetHeader("Subject", "HYBBS--验证码五分钟内有效")
	// 生成随机数
	yzmStr := ""
	for i := 0; i < 4; i++ {
		num := rand.Intn(10)
		yzmStr += strconv.Itoa(num)
	}
	fmt.Println(yzmStr)
	// 内容
	body := "<h1 style=\"text-align: center;\">您的验证码是：</h1><div style=\"padding: 10px;background-color: rgb(72, 82, 82);color: white;text-align: center;width: 50%;margin: 0 auto;border-radius: 5px;letter-spacing: 5px;font-size: 18px;\">" + yzmStr + "</div>"
	message.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.qq.com", 587, "2966678301@qq.com", "nuyvghutvvladffb")

	if err := d.DialAndSend(message); err != nil {
		fmt.Printf("DialAndSend err %v", err)
		panic(err)
	}
	fmt.Printf("send email success\n")

	return yzmStr
}
