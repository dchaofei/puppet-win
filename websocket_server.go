package puppet_win

import (
	"encoding/json"
	"github.com/dchaofei/puppet-win/winapi"
	"github.com/gorilla/websocket"
	"github.com/tidwall/pretty"
	logger "github.com/wechaty/go-wechaty/wechaty-puppet/log"
	"net/http"
	"os"
)

// TODO debug
var outfile *os.File

func init() {
	// TODO debug
	return
	var err error
	outfile, err = os.OpenFile("out.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(outfile)
	}
}

func writeFile(b []byte) {
	// TODO debug
	return
	msg := pretty.Pretty(b)
	_, err := outfile.Write(msg)
	if err != nil {
		panic(err)
	}
	err = outfile.Sync()
	if err != nil {
		panic(err)
	}
}

var log = logger.L.WithField("module", "wechaty-puppet-win")
var upgrader = websocket.Upgrader{}

type webSocketServer struct {
	Port            string
	handleEventFunc func(msg *winapi.EventMsg)
	httpServer      *http.Server
}

func newWebsocketServer(port string, handleEventFunc func(msg *winapi.EventMsg)) *webSocketServer {
	ws := &webSocketServer{
		Port:            port,
		handleEventFunc: handleEventFunc,
	}

	return ws
}

func (ws *webSocketServer) start() error {
	server := &http.Server{Addr: ":" + ws.Port, Handler: http.HandlerFunc(ws.handler)}
	ws.httpServer = server

	return server.ListenAndServe()
}

func (ws *webSocketServer) handler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("webSocketServer.handler upgrade:", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Error("webSocketServer.handler read:", err)
			break
		}
		if mt != websocket.TextMessage {
			log.Errorf("webSocketServer.handler unknown type: %d, data=[%s]", mt, message)
			continue
		}
		if len(message) == 5 && string(message) == "hello" {
			log.Tracef("webSocketServer.handler %s", string(message))
			continue
		}

		writeFile(message)
		eventMsg := &winapi.EventMsg{}
		err = json.Unmarshal(message, eventMsg)
		if err != nil {
			log.Errorf("webSocketServer.handler Unmarshal err:%s, data=[%s]", err.Error(), message)
			continue
		}

		//log.Tracef("websocket_event data=[%s]", message)
		eventMsg.Raw = message
		go ws.handleEventFunc(eventMsg)
	}
}
