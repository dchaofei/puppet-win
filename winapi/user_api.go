package winapi

import "fmt"

func (w *WinApi) GetUser(userIDs []string) (*GetUserResp, error) {
	log.Tracef("WinApi GetUser(%v ... %v)", userIDs[0], len(userIDs))
	resp := &GetUserResp{}
	err := w.request(map[string]interface{}{
		"type":     10015,
		"userName": userIDs,
	}, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WinApi) BatchGetUser(userIDs []string) (*GetUserResp, error) {
	log.Tracef("WinApi BatchGetUser(%v ... %v)", userIDs[0], len(userIDs))
	resp := &GetUserResp{}
	err := w.request(map[string]interface{}{
		"type":      21,
		"userNames": userIDs,
	}, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WinApi) GetSavedFriendUsers() ([]*User, error) {
	userNames, err := w.GetSavedFriendUserNames()
	if err != nil {
		return nil, err
	}

	var (
		users []*User
	)
	userNameBatch := w.splitBatch(userNames, 100)
	batchCount := len(userNameBatch)
	for i, userNames := range userNameBatch {
		log.Tracef("GetSavedFriendUsers %d/%d userNameLen:%d", i+1, batchCount, len(userNames))
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
			user.IsFriend = true
			users = append(users, user)
		}
	}

	return users, nil
}

func (w *WinApi) splitBatch(userNames []string, num int) [][]string {
	max := len(userNames)
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]string{userNames}
	}
	//获取应该数组分割为多少份
	var quantity int
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]string, 0)
	//声明分割数组的截止下标
	var start, end, i int
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, userNames[start:end])
		} else {
			segments = append(segments, userNames[start:])
		}
		start = i * num
	}
	return segments
}

// GetSavedFriendUserNames 获取通讯录好友
func (w *WinApi) GetSavedFriendUserNames() ([]string, error) {
	sql := "SELECT t1.UserName FROM Contact t1 WHERE t1.VerifyFlag = 0 AND (t1.Type = 3 OR t1.Type > 50) and t1.Type != 2050 AND t1.UserName not like '%@chatroom' and t1.UserName NOT IN ('qmessage', 'tmessage') and t1.DelFlag!=1;"
	resp, err := w.getContactUserNameResp(sql)
	if err != nil {
		return nil, err
	}

	return w.users2UserNames(resp.Data.Data), nil
}

// GetSavedRoomUserNames 获取通讯录群聊
func (w *WinApi) GetSavedRoomUserNames() ([]string, error) {
	sql := "SELECT t1.UserName FROM Contact t1 WHERE t1.VerifyFlag = 0 AND (t1.Type = 3 OR t1.Type > 50) AND t1.UserName like '%@chatroom' and t1.DelFlag!=1;"
	resp, err := w.getContactUserNameResp(sql)
	if err != nil {
		return nil, err
	}

	return w.users2UserNames(resp.Data.Data), nil
}

func (w *WinApi) getContactUserNameResp(sql string) (*GetContactUserNameResp, error) {
	resp := &GetContactUserNameResp{}
	err := w.request(map[string]interface{}{
		"type":   10058,
		"dbName": "MicroMsg.db",
		"sql":    sql,
	}, resp)
	if err != nil {
		return nil, err
	}

	if resp.Data.Status != 0 {
		return nil, fmt.Errorf("status:%d desc:%s", resp.Data.Status, resp.Data.Desc)
	}
	return resp, nil
}

func (w *WinApi) users2UserNames(users []*User) []string {
	var usernames []string
	for _, user := range users {
		usernames = append(usernames, user.UserName)
	}
	return usernames
}

// SearchUser 好友、群、群成员
func (w *WinApi) SearchUser(userIDs []string) (*SearchResp, error) {
	log.Tracef("WinApi SearchUser(%v ... %v)", userIDs[0], len(userIDs))
	resp := &SearchResp{}
	err := w.request(map[string]interface{}{
		"type":      10112,
		"userNames": userIDs,
	}, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
