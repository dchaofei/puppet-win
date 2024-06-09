package schemamapping

import (
	"github.com/dchaofei/puppet-win/winapi"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
)

func RoomToWecahty(user *winapi.User) *schemas.RoomPayload {
	return &schemas.RoomPayload{
		Id:           user.UserName,
		Topic:        user.NickName,
		Avatar:       user.SmallHeadImgUrl,
		MemberIdList: members2MemberIDs(user.MemberList),
		OwnerId:      user.OwnerID,
		AdminIdList:  user.AdminIDs,
	}
}

func RoomMemberToWechaty(user *winapi.User) *schemas.RoomMemberPayload {
	return &schemas.RoomMemberPayload{
		Id:        user.UserName,
		RoomAlias: user.Remark,
		InviterId: "",
		Avatar:    user.SmallHeadImgUrl,
		Name:      user.NickName,
	}
}

func members2MemberIDs(members []*winapi.User) []string {
	var res []string
	for _, m := range members {
		res = append(res, m.UserName)
	}
	return res
}
