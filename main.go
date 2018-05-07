package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
	"github.com/superSTN/autoclock/tool"
)

const (
	loginURL = "https://passport.escience.cn/oauth2/authorize?client_id=58861&redirect_uri=http%3A%2F%2F159.226.29.10%2FCnicCheck%2Ftesttoken&response_type=code&theme=simple&pageinfo=userinfo"
	clockURL = "http://159.226.29.10/CnicCheck/CheckServlet"
	username = "xiaoxj@cnic.cn"
	password = "xiao3285158528"
	weidu = 39.97799936612442
	jingdu = 116.32936686204675
	checkin = "checkin"       // 上班
	checkout = "checkout"     // 下班
	filePath = "log.txt"
)

type User struct {
	Token			 string     `json:"token"`
	Uname            string     `json:"uname"`
	Uemail           string     `json:"uemail"`
	RefreshToken     string     `json:"refreshToken"`
}

type ResultRSP struct {
	ErrorMessage     string     `json:"errorMessage"`
	Success          string     `json:"success"`
}


func main()  {
	ticker := time.NewTicker(1 * time.Hour)
	for range ticker.C {
		tm := time.Now().Hour()
		if tm == 7 {
			token := login()
			clock(token, checkin)
		} else if tm == 18 {
			token := login()
			clock(token, checkout)
		}
	}

}

// 登陆返回token函数
func login() string {
	loginURI := loginURL
	loginURI += "&userName=" + username
	loginURI += "&password=" + password

	rsp, err := http.Get(loginURI)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	var user User
	json.Unmarshal(body, &user)
	fmt.Printf("user : %#v\n", user)
	return user.Token
}

// 打卡函数
func clock(token string, clockType string)  {
	clockURI := clockURL
	clockURI += "?weidu=" + strconv.FormatFloat(weidu + (rand.Float64() - 0.5) / 1000, 'f', -1, 64)
	clockURI += "&jingdu=" + strconv.FormatFloat(jingdu + (rand.Float64() - 0.5) / 1000, 'f', -1, 64)
	clockURI += "&type=" + clockType
	clockURI += "&token=" + token

	rsp, err := http.Get(clockURI)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	var resultRSP ResultRSP
	json.Unmarshal(body, &resultRSP)
	fmt.Printf("result : %#v\n", resultRSP)
	if resultRSP.Success == "true" {
		if clockType == checkin {
			fmt.Println("上班打卡成功")
			tool.SendCheckIn()
		} else if clockType == checkout {
			fmt.Println("下班打卡成功")
			tool.SendCheckOut()
		}
	} else {
		fmt.Printf("打卡失败, type: %v, err: %v", clockType, resultRSP.ErrorMessage)
	}
}