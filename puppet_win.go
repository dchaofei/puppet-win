package puppet_win

import (
	"errors"
	"fmt"
	"github.com/dchaofei/puppet-win/schemamapping"
	"github.com/dchaofei/puppet-win/utils"
	"github.com/dchaofei/puppet-win/winapi"
	wechatyPuppet "github.com/wechaty/go-wechaty/wechaty-puppet"
	"github.com/wechaty/go-wechaty/wechaty-puppet/filebox"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"strconv"
	"time"
)

type PuppetWin struct {
	*wechatyPuppet.Puppet

	webSocketServer *webSocketServer
	winApi          *winapi.WinApi
	option          Options
	webSocketHost   string

	cacheMgr *cacheMgr
}

func NewPuppetWin(o Options) (*PuppetWin, error) {
	puppetAbstract, err := wechatyPuppet.NewPuppet(wechatyPuppet.Option{})
	if err != nil {
		return nil, err
	}
	puppetWin := &PuppetWin{
		Puppet: puppetAbstract,
	}
	puppetAbstract.SetPuppetImplementation(puppetWin)

	o.WebsocketPort, err = envWebsocketPort(o.WebsocketPort)
	if err != nil {
		return nil, err
	}
	webSocketServer := newWebsocketServer(o.WebsocketPort, puppetWin.handleEvent)
	puppetWin.webSocketServer = webSocketServer

	o.WebsocketHost, err = envWebsocketHost(o.WebsocketHost)
	if err != nil {
		return nil, err
	}

	winApiServer, err := envWinAPIServer(o.WinApiServer)
	if err != nil {
		return nil, err
	}
	puppetWin.winApi = winapi.NewWinApi(winApiServer)
	puppetWin.option = o

	return puppetWin, nil
}

func (p *PuppetWin) Start() (err error) {
	qrCodeResp, err := p.winApi.GetLoginQrCode()
	if err != nil {
		return err
	}
	if qrCodeResp.IsLogin() {
		p.login(qrCodeResp.CommonResp.Robot)
	}

	webSocketUrl := fmt.Sprintf("ws://%s:%s/", p.option.WebsocketHost, p.option.WebsocketPort)
	err = p.winApi.AddMsgHandler(winapi.WebsocketProtocolType, webSocketUrl)
	if err != nil {
		return err
	}

	go func() {
		err := p.webSocketServer.start()
		if err != nil {
			log.Errorf("p.webSocketServer.start() %s", err.Error())
		}
	}()

	return nil
}

func (p *PuppetWin) Stop() {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) login(robot *winapi.Robot) {
	p.cacheMgr = &cacheMgr{}
	err := p.cacheMgr.init()
	if err != nil {
		log.Errorf("login cacheMgr.init err:%s", err.Error())
		return
	}

	p.SetID(robot.UserName)
	// TODO 从缓存管理器获取联系人信息，缓存管理器支持保存到文件

	go p.Emit(schemas.PuppetEventNameLogin, &schemas.EventLoginPayload{
		ContactId: robot.UserName,
	})

	go p.ready()
}

func (p *PuppetWin) MessageImage(messageID string, imageType schemas.ImageType) (*filebox.FileBox, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) Logout() error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) Ding(data string) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) SetContactAlias(contactID string, alias string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) ContactAlias(contactID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) ContactList() ([]string, error) {
	return p.cacheMgr.contactCache.Keys(), nil
}

func (p *PuppetWin) ContactQRCode(contactID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) SetContactAvatar(contactID string, fileBox *filebox.FileBox) error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) ContactAvatar(contactID string) (*filebox.FileBox, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) ContactRawPayload(contactID string) (*schemas.ContactPayload, error) {
	contact, _ := p.cacheMgr.contactCache.Get(contactID)
	if contact != nil {
		return user2Contact(contact), nil
	}

	contact, err := p._refreshContact(contactID, false)
	if err != nil {
		return nil, err
	}
	if contact == nil {
		return nil, errors.New("contact not found")
	}
	return user2Contact(contact), nil
}

func (p *PuppetWin) SetContactSelfName(name string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) ContactSelfQRCode() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) SetContactSelfSignature(signature string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) MessageContact(messageID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) MessageSendMiniProgram(conversationID string, miniProgramPayload *schemas.MiniProgramPayload) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) MessageRecall(messageID string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) MessageFile(id string) (*filebox.FileBox, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) MessageLocation(messageID string) (*schemas.LocationPayload, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) MessageSendLocation(conversationID string, payload *schemas.LocationPayload) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) MessageRawPayload(id string) (*schemas.MessagePayload, error) {
	msg, ok := p.cacheMgr.msgCache.Get(id)
	if !ok {
		return nil, fmt.Errorf("msg(%s) not found", id)
	}

	return schemamapping.MsgToWechaty(msg), nil
}

func (p *PuppetWin) MessageSendText(conversationID string, text string, mentionIDList ...string) (string, error) {
	resp, err := p.winApi.SendText(conversationID, text, mentionIDList)
	if err != nil {
		return "", err
	}

	msgPayload := &winapi.MsgEventData{}
	msgPayload.Type = winapi.MsgTypeText
	msgPayload.Content = text
	msgPayload.TalkerInfo.UserName = p.SelfID()
	msgPayload.To = conversationID
	p._onSendMessage(msgPayload, resp.Data.MsgSvrID)

	return strconv.FormatInt(resp.Data.MsgSvrID, 10), nil
}

func (p *PuppetWin) _onSendMessage(msg *winapi.MsgEventData, msgID int64) {
	msg.MsgSvrID = msgID
	msg.Createtime = int(time.Now().Unix())

	p.cacheMgr.msgCache.Add(strconv.FormatInt(msgID, 10), msg)
	// TODO revoke info 参考 padlocal
}

func (p *PuppetWin) MessageSendFile(conversationID string, fileBox *filebox.FileBox) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) MessageSendContact(conversationID string, contactID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) MessageSendURL(conversationID string, urlLinkPayload *schemas.UrlLinkPayload) (string, error) {
	content := utils.BuildUrlXml(urlLinkPayload.Title, urlLinkPayload.Description, urlLinkPayload.Url, urlLinkPayload.ThumbnailUrl)
	resp, err := p.winApi.SendURL(conversationID, content)
	if err != nil {
		return "", err
	}

	msgPayload := &winapi.MsgEventData{}
	msgPayload.Type = winapi.MsgTypeXML
	msgPayload.Content = content
	msgPayload.TalkerInfo.UserName = p.SelfID()
	msgPayload.To = conversationID
	p._onSendMessage(msgPayload, resp.Data.MsgSvrID)

	return strconv.FormatInt(resp.Data.MsgSvrID, 10), nil
}

func (p *PuppetWin) MessageURL(messageID string) (*schemas.UrlLinkPayload, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomRawPayload(id string) (*schemas.RoomPayload, error) {
	log.Tracef("RoomRawPayload(%s)", id)

	room, _ := p.cacheMgr.roomCache.Get(id)
	if room == nil {
		var err error
		room, err = p._refreshContact(id, false)
		if err != nil {
			return nil, err
		}
	}
	if room == nil {
		return nil, fmt.Errorf("room(%s) not found", id)
	}
	return schemamapping.RoomToWecahty(room), nil
}

func (p *PuppetWin) RoomList() ([]string, error) {
	return p.cacheMgr.roomCache.Keys(), nil
}

func (p *PuppetWin) RoomDel(roomID, contactID string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomAvatar(roomID string) (*filebox.FileBox, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomAdd(roomID, contactID string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) SetRoomTopic(roomID string, topic string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomTopic(roomID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomCreate(contactIDList []string, topic string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomQuit(roomID string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomQRCode(roomID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomMemberList(roomID string) ([]string, error) {
	members, _, err := p._getRoomMemberList(roomID, false)
	if err != nil {
		return nil, err
	}

	var memberIDs []string
	for _, member := range members {
		memberIDs = append(memberIDs, member.UserName)
	}
	return memberIDs, nil
}

func (p *PuppetWin) RoomMemberRawPayload(roomID string, contactID string) (*schemas.RoomMemberPayload, error) {
	members, ok := p.cacheMgr.roomMemberCache.Get(roomID)
	if !ok {
		return nil, fmt.Errorf("RoomMemberRawPayload(%s, %s) room not found", roomID, contactID)
	}

	member := members[contactID]
	if member == nil {
		return nil, fmt.Errorf("RoomMemberRawPayload(%s, %s) roomMember notfound", roomID, contactID)
	}

	return schemamapping.RoomMemberToWechaty(member), nil
}

func (p *PuppetWin) SetRoomAnnounce(roomID, text string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomAnnounce(roomID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomInvitationAccept(roomInvitationID string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) RoomInvitationRawPayload(id string) (*schemas.RoomInvitationPayload, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) FriendshipSearchPhone(phone string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) FriendshipSearchWeixin(weixin string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) FriendshipRawPayload(id string) (*schemas.FriendshipPayload, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) FriendshipAdd(contactID, hello string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) FriendshipAccept(friendshipID string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) TagContactAdd(id, contactID string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) TagContactRemove(id, contactID string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) TagContactDelete(id string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) TagContactList(contactID string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) MessageRawMiniProgramPayload(messageID string) (*schemas.MiniProgramPayload, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PuppetWin) ready() {
	log.Tracef("PuppetWin readying ...")
	users, err := p.winApi.GetSavedFriendUsers()
	if err != nil {
		log.Errorf("ready GetSavedFriendUsers err: %s", err.Error())
		return
	}
	for _, user := range users {
		p._onPushContact(user)
	}

	rooms, err := p.winApi.GetSavedRoomUsers()
	if err != nil {
		log.Errorf("ready GetSavedRoomUsers err: %s", err.Error())
		return
	}
	for _, room := range rooms {
		p._onPushContact(room)
	}

	// TODO 加载公众号

	log.Tracef("PuppetWin ready")
	p.Emit(schemas.PuppetEventNameReady, &schemas.EventReadyPayload{})
}

func (p *PuppetWin) _onPushContact(user *winapi.User) {
	p._updateContactCache(user)
}

func (p *PuppetWin) _updateContactCache(user *winapi.User) {
	if user.UserName == "" {
		return
	}

	if utils.IsRoomID(user.UserName) {
		p._updateRoomCache(user)
		return
	}

	p.cacheMgr.contactCache.Add(user.UserName, user)
	_ = p.DirtyPayload(schemas.PayloadTypeContact, user.UserName)
}

func (p *PuppetWin) _updateRoomCache(user *winapi.User) {
	members, roomInfo, err := p._getRoomMemberList(user.UserName, true)
	if err != nil {
		log.Errorf("_updateRoomCache _getRoomMemberList err: %s", err.Error())
		return
	}
	user.MemberNum = len(members)
	user.MemberList = p._memberMap2List(members)
	if roomInfo != nil {
		user.OwnerID = roomInfo.OwnerUserName
		user.AdminIDs = roomInfo.ChatroomAdminUserNames
	}
	p.cacheMgr.roomCache.Add(user.UserName, user)
	_ = p.DirtyPayload(schemas.PayloadTypeRoom, user.UserName)
	p._updateRoomMember(user.UserName, nil)
}

func (p *PuppetWin) _updateRoomMember(roomID string, roomMemberMap map[string]*winapi.User) {
	if roomMemberMap != nil {
		p.cacheMgr.SetRoomMember(roomID, roomMemberMap)
	} else {
		p.cacheMgr.roomMemberCache.Remove(roomID)
	}
	_ = p.DirtyPayload(schemas.PayloadTypeRoomMember, roomID)
}

func (p *PuppetWin) _getRoomMemberList(roomID string, force bool) (map[string]*winapi.User, *winapi.GetRoomMemberRespData, error) {
	ret, _ := p.cacheMgr.roomMemberCache.Get(roomID)
	if ret != nil && !force {
		return ret, nil, nil
	}

	resp, err := p.winApi.GetRoomMember(roomID)
	if err != nil {
		return nil, nil, err
	}

	var pendingIDs []string

	membersMap := make(map[string]*winapi.User)
	for _, member := range resp.Data.Data.Members {
		contact, _ := p.cacheMgr.contactCache.Get(member.UserName)
		hasContact := contact != nil
		if !hasContact {
			pendingIDs = append(pendingIDs, member.UserName)
		} else {
			contact = contact.CloneContact()
			contact.Remark = member.MemberTraceBack
			membersMap[member.UserName] = contact
		}
	}

	if len(pendingIDs) > 0 {
		usersResp, err := p.winApi.SearchUser(pendingIDs)
		if err != nil {
			return nil, nil, err
		}

		for _, user := range usersResp.Data.Data {
			p.cacheMgr.contactCache.Add(user.UserName, user)
			membersMap[user.UserName] = user
		}
	}

	p._updateRoomMember(roomID, membersMap)
	return membersMap, resp.Data.Data, nil
}

func (p *PuppetWin) _searchContact(id string) (*winapi.User, error) {
	log.Tracef("PuppetWin _searchContact(%v)", id)
	resp, err := p.winApi.SearchUser([]string{id})
	if err != nil {
		return nil, err
	}

	user := resp.Data.Data[id]
	return user, nil
}

func (p *PuppetWin) _refreshContact(id string, isDelete bool) (*winapi.User, error) {
	contact, err := p._searchContact(id)
	if err != nil {
		return nil, err
	}
	if contact == nil {
		return nil, nil
	}
	if isDelete {
		contact.IsFriend = false
	}
	p._updateContactCache(contact)
	return contact, nil
}

func (p *PuppetWin) _memberMap2List(members map[string]*winapi.User) []*winapi.User {
	var users []*winapi.User
	for _, u := range members {
		users = append(users, u)
	}
	return users
}
