package vio

import "github.com/dxyinme/Luka/chatMsg"

type ChVio interface {

	// 往chVio中加入一段消息
	AddMessage(s chatMsg.Msg) error

	// 从chVio中获取一个消息
	GetMessage() ([]byte,error)

	// 连接名字
	Name() string
}