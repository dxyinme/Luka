package chatMsg

import "github.com/dxyinme/Luka/util"

const (
	checkPoint = 0
	fromBegin = checkPoint + 1
	targetBegin = fromBegin + 32
	msgTypeBegin = targetBegin + 32
	msgContentTypeBegin = msgTypeBegin + 2
	contentBegin = msgContentTypeBegin + 2
	checkPointNone = 0
	checkPointYes = 1
)


// 统一化的LukaMsg 直接使用 bytes 保存内容
type LukaMsg struct {
	s []byte
}

func (l LukaMsg) GetContent() string {
	return string(l.s[contentBegin:])
}

func (l LukaMsg) GetFrom() string {
	return util.ByteToString(l.s[fromBegin: targetBegin])
}

func (l LukaMsg) GetTarget() string {
	return util.ByteToString(l.s[targetBegin: msgTypeBegin])
}

func (l LukaMsg) GetMsgType() MsgTypeEnum {
	MsgTypeByte := l.s[msgTypeBegin: msgContentTypeBegin]
	return MsgTypeEnum(util.ByteToInt16(MsgTypeByte))
}

func (l LukaMsg) GetMsgContentType() MsgContentTypeEnum {
	MsgContentTypeByte := l.s[msgContentTypeBegin: contentBegin]
	return MsgContentTypeEnum(util.ByteToInt16(MsgContentTypeByte))
}

func (l LukaMsg) Marshal() []byte {
	return l.s
}

func NewLukaMsg(From,Target string,
	msgType MsgTypeEnum, msgContentType MsgContentTypeEnum,
	content []byte, isContentLuka bool) *LukaMsg {
	var nows []byte
	if !isContentLuka {
		nows = append(nows, byte(0))
	} else {
		nows = append(nows, byte(1))
	}
	nows = append(nows, util.StringToByteStaticLength(From, 32) ...)
	nows = append(nows, util.StringToByteStaticLength(Target, 32) ...)
	nows = append(nows, util.Int16ToByte(int16(msgType)) ...)
	nows = append(nows, util.Int16ToByte(int16(msgContentType)) ...)
	nows = append(nows, content ...)
	return &LukaMsg{nows}
}

func NewLukaMsgByte(s_ []byte) *LukaMsg {
	return &LukaMsg{s: s_}
}