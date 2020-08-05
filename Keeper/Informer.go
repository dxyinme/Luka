package Keeper

import (
	"github.com/dxyinme/Luka/chatMsg"
	"github.com/dxyinme/Luka/util"
	"sync"
)

const (
	MessageLength = 200
)

var (
	updateMessage chan chatMsg.Msg
	mutexChan sync.Mutex
)


func InitInformer() error {
	updateMessage = make(chan chatMsg.Msg, MessageLength)
	return util.NewTimeTask( "*/5 * * * * ?" , func() {
		//fmt.Println("informer 5s")
		// todo Inform message to Master
	})
}

func pack() []chatMsg.Msg {
	mutexChan.Lock()
	nowlen := len(updateMessage)
	var upSendPack []chatMsg.Msg
	for i := 0; i < nowlen; i++ {
		upSendPack = append(upSendPack, <-updateMessage)
	}
	mutexChan.Unlock()
	return upSendPack
}


func ReceiveUpload(msg chatMsg.Msg) {
	mutexChan.Lock()
	updateMessage <- msg
	mutexChan.Unlock()
}