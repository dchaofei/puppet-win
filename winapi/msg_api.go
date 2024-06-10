package winapi

import (
	"fmt"
	"github.com/pkg/errors"
)

func (w *WinApi) SendText(userName string, text string, atUserList []string) (*SendTextResp, error) {
	log.Tracef("WinApi SendText(%v,%v,%v)", userName, text, atUserList)
	resp := &SendTextResp{}
	err := w.request(map[string]interface{}{
		"type":             10009,
		"userName":         userName,
		"msgContent":       text,
		"atUserList":       atUserList,
		"insertToDatabase": true,
	}, resp)
	if err != nil {
		return nil, errors.Wrap(err, "SendText")
	}

	if resp.Data.Status != 0 {
		return nil, fmt.Errorf("status:%d desc:%s", resp.Data.Status, resp.Data.Desc)
	}
	return resp, nil
}

func (w *WinApi) SendURL(userName string, content string) (*SendTextResp, error) {
	log.Tracef("WinApi SendURL(%v,%v)", userName, "...")

	resp := &SendTextResp{}
	err := w.request(map[string]interface{}{
		"type":             10092,
		"userName":         userName,
		"content":          content,
		"insertToDatabase": true,
	}, resp)
	if err != nil {
		return nil, errors.Wrap(err, "SendURL")
	}

	if resp.Data.Status != 0 {
		return nil, fmt.Errorf("status:%d desc:%s", resp.Data.Status, resp.Data.Desc)
	}
	return resp, nil
}
