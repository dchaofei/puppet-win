package msgparse

import (
	"github.com/dchaofei/puppet-win/winapi"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
)

var WinMsgType2WechatyMsgType = map[winapi.MsgType]schemas.MessageType{
	winapi.MsgTypeText:     schemas.MessageTypeText,
	winapi.MsgTypeImage:    schemas.MessageTypeImage,
	winapi.MsgTypeAudio:    schemas.MessageTypeAudio,
	winapi.MsgTypeEmoticon: schemas.MessageTypeEmoticon,
	winapi.MsgTypeLocation: schemas.MessageTypeLocation,
	winapi.MsgTypeXML:      schemas.MessageTypeAttachment,
}
