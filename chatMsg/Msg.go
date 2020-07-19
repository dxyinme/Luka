package chatMsg

// 消息类型(1.群聊/2.单聊/)
type MsgTypeEnum int
// 消息内容类型(1.文本/2.图片/)
type MsgContentTypeEnum int

const (
	Group  MsgTypeEnum = 1
	Single MsgTypeEnum = 2
)

const (
	Text MsgContentTypeEnum = 1
	Img	 MsgContentTypeEnum = 2
)

// chatMsg接口
type Msg interface {

	// 获取 传输内容的byte转string(如果能转)
	GetContent() 	string

	// 获取发送者
	GetFrom()		string

	// 获取传送目标
	GetTarget()		string

	// 获取传送消息类型
	GetMsgType()	MsgTypeEnum

	// 获取传送消息内容类型
	GetMsgContentType()	MsgContentTypeEnum

	// 获取转换之后的[]byte 即将发送给客户端
	Marshal() ([]byte, error)

}

// 重复字段
type commonField struct {
	From 			string
	Target 			string
	MsgTime			string
	MsgType 		MsgTypeEnum
	MsgContentType 	MsgContentTypeEnum
}
