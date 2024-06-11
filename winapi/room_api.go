package winapi

import "github.com/pkg/errors"

func (w *WinApi) GetSavedRoomUsers() ([]*User, error) {
	userNames, err := w.GetSavedRoomUserNames()
	if err != nil {
		return nil, err
	}

	var (
		users []*User
	)
	userNameBatch := w.splitBatch(userNames, 100)
	batchCount := len(userNameBatch)
	for i, userNames := range userNameBatch {
		log.Tracef("GetSavedRoomUsers %d/%d userNameLen:%d", i+1, batchCount, len(userNames))
		if len(userNames) == 0 {
			continue
		}
		resp, err := w.BatchGetUser(userNames)
		if err != nil {
			return nil, err
		}
		for _, user := range resp.Data.Data.Profiles {
			if user.ErrorCode != 0 {
				continue
			}
			users = append(users, user)
		}
	}

	return users, nil
}

func (w *WinApi) GetRoomMember(roomID string) (*GetRoomMemberResp, error) {
	log.Tracef("WinApi GetRoomMember(%s)", roomID)
	resp := &GetRoomMemberResp{}
	err := w.request(map[string]interface{}{
		"type":             31,
		"chatroomUserName": roomID,
	}, resp)
	if err != nil {
		return nil, errors.Wrap(err, "GetRoomMember")
	}

	return resp, nil
}
