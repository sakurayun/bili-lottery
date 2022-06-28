package main

import (
	"fmt"
	"github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
	"time"
)

func loginStart() {
	headers := map[string]string{
		"dnt":        "1",
		"user-agent": ua,
		"origin":     "https://www.bilibili.com",
		"refer":      "https://www.bilibili.com/",
	}

	login.SetHeaders(headers)

	if config.Cookies.Sessdata == "" {
		webLogin()
	}

	cookies := []*http.Cookie{
		{Name: "SESSDATA", Value: config.Cookies.Sessdata},
		{Name: "bili_jct", Value: config.Cookies.BiliJct},
		{Name: "DedeUserID", Value: config.Cookies.DedeUserID},
		{Name: "DedeUserID__ckMd5", Value: config.Cookies.DedeUserIDCkMd5},
	}

	client.SetCookies(cookies)

Loop:
	if !verifyLogin() {
		webLogin()
		goto Loop
	}
}

func webLogin() {
	var mode int

	log.Println("暂未检测到 SESSDATA, 需要进行扫码登录，请选择登陆模式 (输入数字).")
	fmt.Println("\t1. 终端中生成二维码.")
	fmt.Println("\t2. 当前目录下生成二维码图片.")
	fmt.Println("\t3. APP 打开 URL 登陆.")

Loop:
	_, err := fmt.Scanf("%v", &mode)
	checkErr(err)

	qrCodeUrl, OauthKey := getLoginUrl()

	switch mode {
	case 1:
		obj := qrcodeTerminal.New()
		obj.Get(qrCodeUrl).Print()
	case 2:
		err = qrcode.WriteFile(qrCodeUrl, qrcode.Medium, 256, "./login.png")
		log.Println("已在当前目录下生成二维码，请查看.")
		checkErr(err)
	case 3:
		log.Println("请将此 URL 在 APP 中打开:")
		fmt.Println(qrCodeUrl)
	default:
		fmt.Printf("请重新输入: ")
		goto Loop
	}

	getLoginInfo(OauthKey)
}

// 获取二维码以及token
func getLoginUrl() (url, key string) {
	g := &GetLoginUrl{}

	_, err := login.R().
		SetResult(g).
		Get("/qrcode/getLoginUrl")

	checkErr(err)

	qrCodeUrl := g.Data.Url
	oauthKey := g.Data.OauthKey

	return qrCodeUrl, oauthKey
}

// 获取二维码状态
func getLoginInfo(key string) {
	var cookies []*http.Cookie

	for {
		task := time.NewTimer(3 * time.Second)
		data := map[string]string{
			"oauthKey": key,
		}

		g := &GetLoginInfo{}
		r, err := login.R().
			SetFormData(data).
			SetResult(g).
			Post("/qrcode/getLoginInfo")

		checkErr(err)

		if g.Status == true {
			cookies = r.Cookies()
			for _, cookie := range cookies {
				switch cookie.Name {
				case "SESSDATA":
					config.Cookies.Sessdata = cookie.Value
				case "bili_jct":
					config.Cookies.BiliJct = cookie.Value
				case "DedeUserID":
					config.Cookies.DedeUserID = cookie.Value
				case "DedeUserID__ckMd5":
					config.Cookies.DedeUserIDCkMd5 = cookie.Value
				}
			}

			err = writeConf(filePath)
			checkErr(err)
			break
		}
		<-task.C
	}

}
