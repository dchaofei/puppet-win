package msg

import (
	"github.com/dchaofei/puppet-win/msg"
	"github.com/dchaofei/puppet-win/winapi"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
)

var msgParserList []MessageParser

func init() {
	addMsgParser(typeParser)
	addMsgParser(roomParser)
	// TODO app/refere/sys
}

type MsgParserContext struct {
	AppMsgPayload *msg.AppMsgPayload
}

type MessageParser func(wMsg *winapi.MsgEventData, ret *schemas.MessagePayload, context *MsgParserContext) *schemas.MessagePayload

func addMsgParser(parser MessageParser) {
	msgParserList = append(msgParserList, parser)
}

func ExecuteMsgParsers(wMsg *winapi.MsgEventData, ret *schemas.MessagePayload) *schemas.MessagePayload {
	context := &MsgParserContext{AppMsgPayload: nil}

	for _, parser := range msgParserList {
		ret = parser(wMsg, ret, context)
	}

	return ret
}
