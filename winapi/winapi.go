package winapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	logger "github.com/wechaty/go-wechaty/wechaty-puppet/log"
	"io/ioutil"
	"net/http"
	"time"
)

var log = logger.L.WithField("module", "winapi")

type WinApi struct {
	server string
	client *http.Client
}

func NewWinApi(server string) *WinApi {
	w := &WinApi{server: server}
	w.client = &http.Client{
		Timeout: 10 * time.Second,
	}
	return w
}

func (w *WinApi) GetLoginQrCode() (*GetLoginQrCodeResp, error) {
	log.Tracef("WinApi GetLoginQrCode()")
	resp := &GetLoginQrCodeResp{}
	err := w.request(map[string]interface{}{
		"type": 0,
	}, resp)
	if err != nil {
		return nil, errors.Wrap(err, "GetLoginQrCode")
	}

	return resp, nil
}

func (w *WinApi) AddMsgHandler(protocolType MsgHandlerProtocolType, url string) error {
	log.Tracef("WinApi AddMsgHandler(%v,%v)", protocolType, url)
	resp := &CommonResp{}
	err := w.request(map[string]interface{}{
		"type":     1001,
		"protocol": protocolType,
		"url":      url,
	}, resp)
	if err != nil {
		return errors.Wrap(err, "AddMsgHandler")
	}

	return nil
}

type respInterface interface {
	IsSuccess() bool
	Error() error
}

func (w *WinApi) request(params map[string]interface{}, respData respInterface) error {
	reqData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	start := time.Now()
	resp, err := w.client.Post(w.server, "application/json", bytes.NewReader(reqData))
	if err != nil {
		return fmt.Errorf("%s  cost:%s", err.Error(), time.Now().Sub(start))
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("statusCode:%d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, respData)
	if err != nil {
		return fmt.Errorf("%w, raw:[%v]", err, string(body))
	}

	if !respData.IsSuccess() {
		return respData.Error()
	}

	return nil
}
