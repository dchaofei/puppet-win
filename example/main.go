package main

import (
	"fmt"
	puppet_win "github.com/dchaofei/puppet-win"
	"github.com/mdp/qrterminal/v3"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

func main() {
	puppetWin, err := puppet_win.NewPuppetWin(puppet_win.Options{
		WinApiServer:  "http://126.xxx.xx.xxx:8888/api/", // windows 机器注入 dll 启动的服务
		WebsocketHost: "127.0.0.1",                       // 本程序会启动 websocket 服务用于接收消息消息回调，确保本地址能被 windows 机器访问
		WebsocketPort: "25465",                           // 本程序会启动 websocket 服务用于接收消息消息回调，确保本端口能被 windows 机器访问
	})
	if err != nil {
		panic(err)
	}

	var bot = wechaty.NewWechaty(wechaty.WithPuppet(puppetWin))

	bot.OnScan(onScan).OnLogin(func(ctx *wechaty.Context, user *user.ContactSelf) {
		fmt.Printf("User %s logined\n", user.Name())
	}).OnMessage(onMessage).OnLogout(func(ctx *wechaty.Context, user *user.ContactSelf, reason string) {
		fmt.Printf("User %s logouted: %s\n", user, reason)
	})

	bot.DaemonStart()
}

func onMessage(ctx *wechaty.Context, message *user.Message) {
	log.Println(message)

	if message.Self() {
		log.Println("Message discarded because its outgoing")
	}

	if message.Age() > 2*60*time.Second {
		log.Println("Message discarded because its TOO OLD(than 2 minutes)")
	}

	if message.Type() != schemas.MessageTypeText || message.Text() != "#ding" {
		log.Println("Message discarded because it does not match 'ding'")
		return
	}

	// 1. reply 'dong'
	_, err := message.Say("dong")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("REPLY: dong")
}

func onScan(ctx *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
	fmt.Printf("onScan: %s\n", status)
	if status == schemas.ScanStatusWaiting || status == schemas.ScanStatusTimeout {
		qrterminal.GenerateHalfBlock(qrCode, qrterminal.L, os.Stdout)

		qrcodeImageUrl := fmt.Sprintf("https://wechaty.js.org/qrcode/%s", url.QueryEscape(qrCode))
		fmt.Printf("onScan: %s - %s\n", status, qrcodeImageUrl)
		return
	}
}
