package tool

import (
	"strings"
	"net/smtp"
	"fmt"
	"math/rand"
	"time"
)

const (
	user    = "372776076@qq.com"   //你开通smtp邮箱的 邮箱的地址
	pwd     = "brgabdgxuyakbgid"   //这里填你自己的授权码
	host    = "smtp.qq.com:25"
	to      = "372776076@qq.com"   //目标地址
	body = `
      <html>
      <body>
      <h3>
      "打卡成功"
      </h3>
      </body>
      </html>
      `
)

//授权码 brgabdgxuyakbgid

func SendToMail(user, pwd, host, to, subject, body, mailtype string) error{
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, pwd, hp[0])
	var content_type string

	if mailtype == "html" {

		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"

	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"


	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " +subject+ "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func SendCheckIn()  {
	num := rand.Int31n(5 * 60)
	time.Sleep(time.Duration(num) * time.Second)
	subject := "实验室上班自动打卡成功" + string(num)
	fmt.Println("SEND EMAIL ...")
	err :=SendToMail(user, pwd, host, to, subject, body,"html")

	if err != nil{
		fmt.Println("发送邮件错误！")
		fmt.Println(err)
	}else{
		fmt.Println("发送成功")
	}
}

func SendCheckOut()  {
	num := rand.Int31n(5 * 60)
	time.Sleep(time.Duration(num) * time.Second)
	subject := "实验室下班自动打卡成功" + string(num)
	fmt.Println("SEND EMAIL ...")
	err :=SendToMail(user, pwd, host, to, subject, body,"html")

	if err != nil{
		fmt.Println("发送邮件错误！")
		fmt.Println(err)
	}else{
		fmt.Println("发送成功")
	}
}

func SendStart()  {
	subject := "自动打卡启动v2.0"
	fmt.Println("SEND EMAIL ...")
	err :=SendToMail(user, pwd, host, to, subject, body,"html")

	if err != nil{
		fmt.Println("发送邮件错误！")
		fmt.Println(err)
	}else{
		fmt.Println("发送成功")
	}
}