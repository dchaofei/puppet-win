package msg

type AppMsgType int

const (
	AppMsgTypeText                  AppMsgType = 1
	AppMsgTypeImg                   AppMsgType = 2
	AppMsgTypeAudio                 AppMsgType = 3
	AppMsgTypeVideo                 AppMsgType = 4
	AppMsgTypeUrl                   AppMsgType = 5
	AppMsgTypeAttach                AppMsgType = 6
	AppMsgTypeOpen                  AppMsgType = 7
	AppMsgTypeEmoji                 AppMsgType = 8
	AppMsgTypeVoiceRemind           AppMsgType = 9
	AppMsgTypeScanGood              AppMsgType = 10
	AppMsgTypeGood                  AppMsgType = 13
	AppMsgTypeEmotion               AppMsgType = 15
	AppMsgTypeCardTicket            AppMsgType = 16
	AppMsgTypeRealtimeShareLocation AppMsgType = 17
	AppMsgTypeChatHistory           AppMsgType = 19
	AppMsgTypeMiniProgram           AppMsgType = 33
	AppMsgTypeMiniProgramApp        AppMsgType = 36
	AppMsgGroupNote                 AppMsgType = 53
	AppMsgReferMsg                  AppMsgType = 57
	AppMsgTypeTransfers             AppMsgType = 2000
	AppMsgTypeRedEnvelopes          AppMsgType = 2001
	AppMsgTypeReaderType            AppMsgType = 100001
)

type AppMsgPayload struct {
	Des          string
	Thumburl     string
	Title        string
	Url          string
	Appattach    AppAttachPayload
	Type         AppMsgType
	Md5          string
	Fromusername string
	Recorditem   string
	Refermsg     ReferMsgPayload
}

type AppAttachPayload struct {
	Totallen       int
	Attachid       string
	Emoticonmd5    string
	Fileext        string
	Cdnattachurl   string
	Aeskey         string
	Cdnthumbaeskey string
	Encryver       int
	Islargefilemsg int
}

type ReferMsgPayload struct {
	Type        string
	Svrid       string
	Fromusr     string
	Chatusr     string
	Displayname string
	Content     string
}
