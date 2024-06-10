package msg

import (
	"github.com/dchaofei/puppet-win/winapi"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
)

func typeParser(wMsg *winapi.MsgEventData, ret *schemas.MessagePayload, context *MsgParserContext) *schemas.MessagePayload {
	ret.Type = TypeMappings[wMsg.Type]
	return ret
}

var TypeMappings = map[winapi.MsgType]schemas.MessageType{
	winapi.MsgTypeText:     schemas.MessageTypeText,
	winapi.MsgTypeImage:    schemas.MessageTypeImage,
	winapi.MsgTypeAudio:    schemas.MessageTypeAudio,
	winapi.MsgTypeEmoticon: schemas.MessageTypeEmoticon,
	winapi.MsgTypeLocation: schemas.MessageTypeLocation,
	winapi.MsgTypeXML:      schemas.MessageTypeAttachment,
	winapi.MsgTypeSystem:   schemas.MessageTypeUnknown,
}
