package chatMsg

import (
	"github.com/dxyinme/Luka/util"
	"github.com/golang/glog"
)

type UserMsg struct {
	From 			string
	Target 			string
	MsgTime			string
	Content         string
	MsgType 		MsgTypeEnum
	MsgContentType 	MsgContentTypeEnum
}

// 你会拿到的是base64加密后的content内容
func (l UserMsg) GetContent() string {
	return l.Content
}

func (l UserMsg) GetFrom() string {
	return l.From
}

func (l UserMsg) GetTarget() string {
	return l.Target
}

func (l UserMsg) GetMsgType() MsgTypeEnum {
	return l.MsgType
}

func (l UserMsg) GetMsgContentType() MsgContentTypeEnum {
	return l.MsgContentType
}

func (l UserMsg) Marshal() ([]byte,error) {
	return util.IJson.Marshal(l)
}

func NewUserMsgByte(s []byte) *UserMsg {
	var msg *UserMsg
	err := util.IJson.Unmarshal(s, &msg)
	if err != nil {
		glog.Errorf("New UserMsg create failed , because %v", err)
		return nil
	}
	return msg
}

