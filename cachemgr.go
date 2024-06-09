package puppet_win

import (
	"github.com/dchaofei/puppet-win/winapi"
	lru "github.com/hashicorp/golang-lru/v2"
)

type cacheMgr struct {
	msgCache        *lru.Cache[string, *winapi.MsgEventData]
	contactCache    *lru.Cache[string, *winapi.User]
	roomCache       *lru.Cache[string, *winapi.User]
	roomMemberCache *lru.Cache[string, map[string]*winapi.User]
}

func (c *cacheMgr) init() error {
	var err error
	c.msgCache, err = lru.New[string, *winapi.MsgEventData](1024)
	if err != nil {
		return err
	}
	c.contactCache, err = lru.New[string, *winapi.User](1024)
	if err != nil {
		return err
	}
	c.roomCache, err = lru.New[string, *winapi.User](1024)
	if err != nil {
		return err
	}
	c.roomMemberCache, err = lru.New[string, map[string]*winapi.User](1024)
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheMgr) SetRoomMember(roomId string, roomMemberMap map[string]*winapi.User) {
	c.roomMemberCache.Add(roomId, roomMemberMap)
}

func (c *cacheMgr) hasContact(userID string) bool {
	return c.contactCache.Contains(userID)
}
