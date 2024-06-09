package winapi

import (
	"fmt"
	"github.com/pkg/errors"
)

func (w *WinApi) SendText(userName string, text string, atUserList []string) (*SendTextResp, error) {
	log.Tracef("WinApi SendText(%v,%v,%v)", userName, text, atUserList)
	resp := &SendTextResp{}
	err := w.request(map[string]interface{}{
		"type":       10009,
		"userName":   userName,
		"msgContent": text,
		"atUserList": atUserList,
	}, resp)
	if err != nil {
		return nil, errors.Wrap(err, "SendText")
	}

	if resp.Data.Status != 0 {
		return nil, fmt.Errorf("status:%d desc:%s", resp.Data.Status, resp.Data.Desc)
	}
	return resp, nil
}

var urlXml = "<appmsg appid=\"\" sdkver=\"0\">\n\t\t<title>%s</title>\n\t\t<des>%s</des>\n\t\t<type>5</type>\n\t\t<url>%s</url>\n\t\t<appattach>\n\t\t\t<cdnthumbaeskey />\n\t\t\t<aeskey />\n\t\t</appattach>\n\t\t<thumburl>%s</thumburl>\n\t</appmsg>\n\t<scene>0</scene>\n\t<appinfo>\n\t\t<version>1</version>\n\t\t<appname></appname>\n\t</appinfo>\n\t<commenturl></commenturl>"

func (w *WinApi) SendURL(userName string, content string) (*SendTextResp, error) {
	log.Tracef("WinApi SendURL(%v,%v)", userName, "...")

	resp := &SendTextResp{}
	err := w.request(map[string]interface{}{
		"type":     10092,
		"userName": userName,
		"content":  content,
	}, resp)
	if err != nil {
		return nil, errors.Wrap(err, "SendURL")
	}

	if resp.Data.Status != 0 {
		return nil, fmt.Errorf("status:%d desc:%s", resp.Data.Status, resp.Data.Desc)
	}
	return resp, nil
}
