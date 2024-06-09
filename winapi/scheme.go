package winapi

import (
	"fmt"
)

const (
	TcpProtocolType       MsgHandlerProtocolType = 1
	HttpProtocolType      MsgHandlerProtocolType = 2
	WebsocketProtocolType MsgHandlerProtocolType = 3
)

const (
	MsgTypeText               MsgType = 1
	MsgTypeImage              MsgType = 3
	MsgTypeAudio              MsgType = 34
	MsgTypeFriendVerification MsgType = 37
	MsgTypeFriendRecommend    MsgType = 42
	MsgTypeEmoticon           MsgType = 47
	MsgTypeLocation           MsgType = 48
	MsgTypeXML                MsgType = 49
	MsgTypeTelephone          MsgType = 50
	MsgTypeMobileOperate      MsgType = 51
	MsgTypeSystem             MsgType = 10000
	MsgTypeRecall             MsgType = 10002
)

type CommonResp struct {
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
	Robot       *Robot `json:"robot"`
	Type        int    `json:"type"`
}

func (c *CommonResp) IsSuccess() bool {
	if c.ErrorCode != 10000 && c.ErrorCode != 0 {
		return false
	}
	return true
}

func (c *CommonResp) Error() error {
	return fmt.Errorf("errorCode:%d Description:%s", c.ErrorCode, c.Description)
}

type GetLoginQrCodeResp struct {
	Data GetLoginQrCodeData `json:"data"`
	CommonResp
}

func (g *GetLoginQrCodeResp) IsLogin() bool {
	return g.Data.Status == -2
}

type GetLoginQrCodeData struct {
	Desc   string `json:"desc"`
	Qrcode []int  `json:"qrcode"`
	Status int    `json:"status"`
	Uuid   string `json:"uuid"`
}

type GetUserResp struct {
	Data struct {
		Data   *GetUserRespData `json:"data"`
		Desc   string           `json:"desc"`
		Status int              `json:"status"`
	} `json:"data"`
	CommonResp
}

type GetUserRespData struct {
	Count    int     `json:"count"`
	Profiles []*User `json:"profiles"`
}

type User struct {
	Alias           string `json:"alias"`
	BigHeadImgUrl   string `json:"bigHeadImgUrl"`
	CertFlag        int    `json:"certFlag"`
	City            string `json:"city"`
	ErrorCode       int    `json:"errorCode"`
	Fullpy          string `json:"fullpy"`
	Nation          string `json:"nation"`
	NickName        string `json:"nickName"`
	Province        string `json:"province"`
	Remark          string `json:"remark"`
	Reserved1       int    `json:"reserved1"`
	Sex             int    `json:"sex"`
	Simplepy        string `json:"simplepy"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl"`
	UserFlag        int    `json:"userFlag"`
	UserName        string `json:"userName"`

	ChatroomAccessType int `json:"chatroomAccessType"`
	ChatroomMaxCount   int `json:"chatroomMaxCount"`
	ChatroomNotify     int `json:"chatroomNotify"`

	IsFriend   bool     `json:"-"`
	MemberNum  int      `json:"-"`
	OwnerID    string   `json:"-"`
	MemberList []*User  `json:"-"`
	AdminIDs   []string `json:"-"`
}

func (u *User) CloneContact() *User {
	member := *u
	member.reset()
	return &member
}

func (u *User) reset() {
	u.MemberNum = 0
	u.AdminIDs = nil
	u.OwnerID = ""
	u.MemberList = nil
	u.AdminIDs = nil
}

type GetContactUserNameResp struct {
	Data struct {
		Data   []*User `json:"data"`
		Desc   string  `json:"desc"`
		Status int     `json:"status"`
	} `json:"data"`
	CommonResp
}

type MsgHandlerProtocolType int

type MsgType int

type Robot struct {
	Alias           string `json:"alias"`
	IsLogin         bool   `json:"isLogin"`
	NickName        string `json:"nickName"`
	Pid             int    `json:"pid"`
	Port            int    `json:"port"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl"`
	UserName        string `json:"userName"`
}

type GetRoomMemberResp struct {
	Data struct {
		Data   *GetRoomMemberRespData `json:"data"`
		Desc   string                 `json:"desc"`
		Status int                    `json:"status"`
	} `json:"data"`
	CommonResp
}

type GetRoomMemberRespData struct {
	ChatroomAdminUserNames   []string      `json:"chatroomAdminUserNames"`
	ChatroomMemberInfoVerion int           `json:"chatroomMemberInfoVerion"`
	ChatroomUserName         string        `json:"chatroomUserName"`
	Count                    int           `json:"count"`
	Members                  []*RoomMember `json:"members"`
	OwnerUserName            string        `json:"ownerUserName"`
}

type RoomMember struct {
	BigHeadImgUrl   string `json:"bigHeadImgUrl"`
	InviterUserName string `json:"inviterUserName"`
	IsAdmin         bool   `json:"isAdmin"`
	MemberTraceBack string `json:"memberTraceBack"`
	NickName        string `json:"nickName"`
	Permission      int    `json:"permission"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl"`
	UserName        string `json:"userName"`
}

type SearchResp struct {
	Data struct {
		Data   map[string]*User `json:"data"`
		Desc   string           `json:"desc"`
		Status int              `json:"status"`
	} `json:"data"`
	CommonResp
}

type SendTextResp struct {
	Data struct {
		Desc     string `json:"desc"`
		MsgSvrID int64  `json:"msgSvrID"`
		Status   int    `json:"status"`
	} `json:"data"`
	CommonResp
}
