package chatMsg

import (
	"encoding/json"
	"log"
	"time"
)

type targetTypeEnum int

const (
	Group targetTypeEnum = 1
	Single targetTypeEnum = 2
)

// 消息发送的格式
type TextMsg struct {
	// textMsg 是由哪个User发送的，From就是他的name,
	// 用于唯一定位一个user,确定消息的来源。
	From  		string

	// 目标类型 group/single
	TargetType  targetTypeEnum

	// 消息发送的目标，如果是群组则是群id，否则是user的name
	Target 		string

	// 消息内容，以utf-8编码发送
	Content 	string

	// 消息发出时间 Unix Time
	MsgTime 	int64
}

// 请保证调用这个函数的Keeper机器的时间必须与master-server一致
// create a new chatMsg
func NewTextMsg(from string, targetType targetTypeEnum, target string, content string) *TextMsg {
	return &TextMsg{From: from, TargetType: targetType,
		Target: target, Content: content, MsgTime: time.Now().Unix()}
}


// create a new chatMsg by []byte
func NewTextMsgUnmarshal(b []byte) *TextMsg {
	ret := &TextMsg{}
	errJson := json.Unmarshal(b,ret)
	if errJson != nil {
		log.Println(errJson)
		return nil
	}
	return ret
}