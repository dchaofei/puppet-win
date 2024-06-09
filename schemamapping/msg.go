package schemamapping

import (
	"github.com/dchaofei/puppet-win/schemamapping/msgparse"
	"github.com/dchaofei/puppet-win/winapi"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"strconv"
	"time"
)

func MsgToWechaty(data *winapi.MsgEventData) *schemas.MessagePayload {
	// TODO 支持不同类型消息解析
	payload := &schemas.MessagePayload{
		MessagePayloadBase: schemas.MessagePayloadBase{
			Id:            strconv.FormatInt(data.MsgSvrID, 10),
			MentionIdList: nil, // TODO 解析@人？
			FileName:      "",  // TODO ?
			Text:          data.Content,
			Timestamp:     time.Unix(int64(data.Createtime), 0),
			Type:          msgparse.WinMsgType2WechatyMsgType[data.Type],
		},
		MessagePayloadRoom: schemas.MessagePayloadRoom{
			ListenerId: data.To,
		},
	}

	if data.IsChatroomMsg == 1 {
		payload.RoomId = data.From
		payload.TalkerId = data.ChatroomMemberInfo.UserName // TODO 企业微信需要解析xml拿到username
	} else {
		payload.TalkerId = data.From
	}
	return payload
}
