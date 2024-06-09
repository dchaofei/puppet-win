package winapi

import (
	"encoding/json"
	"fmt"
)

const (
	AsyncRespPushType     = 0
	SyncMsgPushType       = 1
	LoginMsgPushType      = 2
	LogoutMsgPushType     = 3
	MsgHookPushType       = 4
	RoomChangeMsgPushType = 5
)

const (
	WaitScanLoginState  = 0
	ScannedLoginState   = 1
	ConfirmedLoginState = 2
)

const (
	ScanLoginStep = 0
	LoggedStep    = 1
)

var pushType2EventData = map[int]func() interface{}{
	LoginMsgPushType: func() interface{} {
		return &LoginEventData{}
	},
	LogoutMsgPushType: func() interface{} {
		return &LogoutEventData{}
	},
	SyncMsgPushType: func() interface{} {
		return &MsgEventData{}
	},
}

type EventMsg struct {
	Data     json.RawMessage `json:"data"`
	PushTime int             `json:"pushTime"`
	PushType int             `json:"pushType"`
	Robot    *Robot          `json:"robot"`

	Raw []byte `json:"-"`
}

func (c *EventMsg) Parse() (interface{}, error) {
	getEventDataFunc, ok := pushType2EventData[c.PushType]
	if !ok {
		return nil, fmt.Errorf("unknown push type %d", c.PushType)
	}

	data := getEventDataFunc()
	err := json.Unmarshal(c.Data, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type LoginEventData struct {
	Alias           string `json:"alias"`
	BigHeadImgUrl   string `json:"bigHeadImgUrl"`
	City            string `json:"city"`
	Nation          string `json:"nation"`
	NickName        string `json:"nickName"`
	Phone           string `json:"phone"`
	Province        string `json:"province"`
	Sex             int    `json:"sex"`
	Signature       string `json:"signature"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl"`
	Step            int    `json:"step"`
	State           int    `json:"state"`
	Uuid            string `json:"uuid"`
	Uin             int    `json:"uin"`
	UserName        string `json:"userName"`
}

type LogoutEventData struct {
	Time     int    `json:"time"`
	UserName string `json:"userName"`
}

type MsgEventData struct {
	ChatroomMemberInfo struct {
		Alias                         string `json:"alias"`
		BelongChatroomNickName        string `json:"belongChatroomNickName"`
		BelongChatroomSmallHeadImgUrl string `json:"belongChatroomSmallHeadImgUrl"`
		BelongChatroomUserName        string `json:"belongChatroomUserName"`
		ChatroomDisplayName           string `json:"chatroomDisplayName"`
		ChatroomUserFlag              int    `json:"chatroomUserFlag"`
		IsChatroomAdmin               bool   `json:"isChatroomAdmin"`
		IsChatroomOwner               bool   `json:"isChatroomOwner"`
		NickName                      string `json:"nickName"`
		Remark                        string `json:"remark"`
		SmallHeadImgUrl               string `json:"smallHeadImgUrl"`
		Type                          string `json:"type"`
		UserName                      string `json:"userName"`
		VerifyFlag                    string `json:"verifyFlag"`
	} `json:"chatroomMemberInfo"`
	Content        string `json:"content"`
	Createtime     int    `json:"createtime"`
	From           string `json:"from"`
	IsChatroomMsg  int    `json:"isChatroomMsg"`
	IsSender       int    `json:"isSender"`
	MsgSvrID       int64  `json:"msgSvrID"`
	Reserved1      string `json:"reserved1"`
	SyncFromMobile bool   `json:"syncFromMobile"`
	TalkerInfo     struct {
		Alias           string `json:"alias"`
		NickName        string `json:"nickName"`
		Remark          string `json:"remark"`
		SmallHeadImgUrl string `json:"smallHeadImgUrl"`
		Type            string `json:"type"`
		UserName        string `json:"userName"`
		VerifyFlag      string `json:"verifyFlag"`
	} `json:"talkerInfo"`
	To   string  `json:"to"`
	Type MsgType `json:"type"`
}
