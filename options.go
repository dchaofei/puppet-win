package puppet_win

type Options struct {
	WinApiServer  string // winApi 的地址："http:/127.0.0.1:8888/api"
	RunRemote     bool   // 是否和 winapi 运行在同一台机器
	WebsocketHost string // 主要是用来注册给 winApi 进行消息通知，所以这个地址必须是 winApi 机器能访问通的: "127.0.0.1"
	WebsocketPort string // 本地需要监听websocket端口，会注册给 winApi 进行消息通知: "8888"
}
