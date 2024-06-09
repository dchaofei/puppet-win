package puppet_win

import (
	"github.com/dchaofei/puppet-win/winapi"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"strings"
)

func user2Contact(user *winapi.User) *schemas.ContactPayload {
	payload := &schemas.ContactPayload{
		Id:       user.UserName,
		Gender:   schemas.ContactGender(user.Sex),
		Name:     user.NickName,
		Avatar:   user.Alias,
		Alias:    user.Alias,
		City:     user.City,
		Province: user.Province,
		Friend:   user.IsFriend,
	}
	if strings.HasPrefix(user.UserName, "gh_") {
		payload.Type = schemas.ContactTypeOfficial
	} else {
		payload.Type = schemas.ContactTypePersonal
	}

	return payload
}
