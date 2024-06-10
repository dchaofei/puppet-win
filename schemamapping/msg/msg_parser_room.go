package msg

import (
	"github.com/dchaofei/puppet-win/utils"
	"github.com/dchaofei/puppet-win/winapi"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"strings"
)

func roomParser(wMsg *winapi.MsgEventData, ret *schemas.MessagePayload, context *MsgParserContext) *schemas.MessagePayload {
	if wMsg.IsChatroomMsg != 1 {
		return ret
	}

	ret.RoomId = wMsg.From
	ret.TalkerId = wMsg.ChatroomMemberInfo.UserName

	// text:    "wxid_xxxx:\nnihao"
	// appmsg:  "wxid_xxxx:\n<?xml version="1.0"?><msg><appmsg appid="" sdkver="0">..."
	// pat:     "19850419xxx@chatroom:\n<sysmsg type="pat"><pat><fromusername>xxx</fromusername><chatusername>19850419xxx@chatroom</chatusername><pattedusername>wxid_xxx</pattedusername>...<template><![CDATA["${vagase}" 拍了拍我]]></template></pat></sysmsg>"
	parts := strings.Split(wMsg.Content, ":\n")
	if len(parts) > 1 {
		ret.Text = strings.Join(parts[1:], "")
		if ret.TalkerId == "" && utils.IsIMContactId(parts[0]) {
			ret.TalkerId = parts[0]
		}
	}

	// TODO parser MentionIdList
	return ret
}
