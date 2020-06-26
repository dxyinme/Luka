package chatMsg

type ImgMsg struct {
	CommonField commonField
	Content 	[]byte
}

func (i ImgMsg) GetContent() []byte {
	return i.Content
}

func (i ImgMsg) GetTime() string {
	return i.CommonField.MsgTime
}

func (i ImgMsg) GetFrom() string {
	return i.CommonField.From
}

func (i ImgMsg) GetTarget() string {
	return i.CommonField.Target
}

func (i ImgMsg) GetMsgType() MsgTypeEnum {
	return i.CommonField.MsgType
}

func (i ImgMsg) GetMsgContentType() MsgContentTypeEnum {
	return i.CommonField.MsgContentType
}


// a new ImgMsg
func NewImgMsg(from string, target string,
	targetType MsgTypeEnum, targetContentType MsgContentTypeEnum, content []byte) *ImgMsg {
	return &ImgMsg{
		CommonField: NewCommonField(from,target,targetType,targetContentType),
		Content:     content,
	}
}

