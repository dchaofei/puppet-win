package schemamapping

import (
	"github.com/dchaofei/puppet-win/schemamapping/msg"
	"github.com/dchaofei/puppet-win/winapi"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"strconv"
	"time"
)

func MsgToWechaty(data *winapi.MsgEventData) *schemas.MessagePayload {
	payload := &schemas.MessagePayload{
		MessagePayloadBase: schemas.MessagePayloadBase{
			Id:            strconv.FormatInt(data.MsgSvrID, 10),
			MentionIdList: nil, // TODO 解析@人？
			FileName:      "",  // TODO ?
			Text:          data.Content,
			Timestamp:     time.Unix(int64(data.Createtime), 0),
		},
		MessagePayloadRoom: schemas.MessagePayloadRoom{
			ListenerId: data.To,
			TalkerId:   data.From,
		},
	}

	return msg.ExecuteMsgParsers(data, payload)
}
