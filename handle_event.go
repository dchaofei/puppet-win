package puppet_win

import (
	"fmt"
	"github.com/dchaofei/puppet-win/winapi"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"strconv"
)

func (p *PuppetWin) handleEvent(msg *winapi.EventMsg) {
	data, err := msg.Parse()
	if err != nil {
		log.Errorf("handleEvent msg.Parse: %s, rawMsg:%s", err, msg.Raw)
		return
	}
	switch msg.PushType {
	case winapi.LoginMsgPushType:
		p.handleLoginEvent(msg, data.(*winapi.LoginEventData))
	case winapi.LogoutMsgPushType:
		p.Emit(schemas.PuppetEventNameLogout, &schemas.EventLogoutPayload{
			ContactId: data.(*winapi.LogoutEventData).UserName,
		})
	case winapi.SyncMsgPushType:
		p.handleSyncMsgEvent(msg, data.(*winapi.MsgEventData))
	}
}

func (p *PuppetWin) handleLoginEvent(msg *winapi.EventMsg, data *winapi.LoginEventData) {
	if data.Step == winapi.ScanLoginStep {
		if msg.Robot.IsLogin {
			return
		}

		p.handleScanEvent(msg, data)
		return
	}

	if data.Step != winapi.LoggedStep {
		log.Errorf("handleLoginEvent unkown step rawMsg:%s", msg.Raw)
		return
	}
	p.login(msg.Robot)
	go p.ready()
}

func (p *PuppetWin) handleScanEvent(msg *winapi.EventMsg, data *winapi.LoginEventData) {
	qrCode := fmt.Sprintf("http://weixin.qq.com/x/%s", data.Uuid)
	payload := &schemas.EventScanPayload{
		QrCode: qrCode,
	}
	switch data.State {
	case winapi.WaitScanLoginState:
		payload.Status = schemas.ScanStatusWaiting
	case winapi.ScannedLoginState:
		payload.Status = schemas.ScanStatusScanned
	case winapi.ConfirmedLoginState:
		payload.Status = schemas.ScanStatusConfirmed
	default:
		log.Errorf("handleScanEvent unkown state rawMsg:%s", msg.Raw)
	}

	p.Emit(schemas.PuppetEventNameScan, payload)
}

func (p *PuppetWin) handleSyncMsgEvent(msg *winapi.EventMsg, data *winapi.MsgEventData) {
	msgID := strconv.FormatInt(data.MsgSvrID, 10)
	if p.cacheMgr.msgCache.Contains(msgID) {
		return
	}
	p.cacheMgr.msgCache.Add(msgID, data)

	switch data.Type {
	// TODO 支持更多消息解析
	//case winapi.MsgTypeFriendVerification:
	//	p.Emit(schemas.PuppetEventNameFriendship, &schemas.EventFriendshipPayload{
	//		FriendshipID: msgID,
	//	})
	default:
		p.Emit(schemas.PuppetEventNameMessage, &schemas.EventMessagePayload{
			MessageId: msgID,
		})
	}

}
