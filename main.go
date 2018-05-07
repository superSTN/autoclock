package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"strconv"
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
	token := login()
	clock(token, checkin)
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
		fmt.Println("打卡成功")
	}
}