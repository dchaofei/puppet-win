package winapi

import (
	"fmt"
	"os"
	"testing"
)

var apiInstance = NewWinApi(os.Getenv("puppetWinServer"))

func TestWinApi_GetLoginQrCode(t *testing.T) {
	qrCode, err := apiInstance.GetLoginQrCode()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(qrCode)
}

func TestWinApi_AddMsgHandler(t *testing.T) {
	err := apiInstance.AddMsgHandler(WebsocketProtocolType, "ws://xxxxx:8232/echo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestWinApi_GetUser(t *testing.T) {
	resp, err := apiInstance.GetUser([]string{"filehelper"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.Data.Data.Profiles[0])
}
