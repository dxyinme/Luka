package chatMsg

import (
	"github.com/dxyinme/Luka/util"
	"github.com/golang/glog"
	"time"
)


// 消息类型
type TmpMsg struct {
	CommonField commonField
	Content 	string
}


func (i TmpMsg) Marshal() ([]byte, error) {
	var (
		err error
		res []byte
	)
	res,err = util.IJson.Marshal(i)
	if err != nil {
		return nil,err
	}
	return res,nil
}

func (i TmpMsg) GetContent() string {
	return i.Content
}

func (i TmpMsg) GetTime() string {
	return i.CommonField.MsgTime
}

func (i TmpMsg) GetFrom() string {
	return i.CommonField.From
}

func (i TmpMsg) GetTarget() string {
	return i.CommonField.Target
}

func (i TmpMsg) GetMsgType() MsgTypeEnum {
	return i.CommonField.MsgType
}

func (i TmpMsg) GetMsgContentType() MsgContentTypeEnum {
	return i.CommonField.MsgContentType
}

func (i *TmpMsg) SetTime(timeNow string) {
	i.CommonField.MsgTime = timeNow
}

// create a new TmpMsg by []byte
func NewTmpMsgUnmarshal(b []byte) *TmpMsg {
	var (
		err error
	)
	var tmp TmpMsg
	err = util.IJson.Unmarshal(b,&tmp)
	if err != nil {
		glog.Warningf("[TmpMsg] Unmarshal error : %v",err)
		return nil
	}
	now := time.Now()
	tmp.SetTime(now.Format(util.TIMELAYOUT))
	return &tmp
}