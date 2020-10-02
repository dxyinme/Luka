package broadcast

import (
	"fmt"
	clusterConfig "github.com/dxyinme/Luka/cluster/config"
	CynicUClient "github.com/dxyinme/LukaComm/CynicU/Client"
	"github.com/dxyinme/LukaComm/chatMsg"
	"time"
)

type BroadcasterForPull struct {
	clients 	[]*CynicUClient.Client
	respItems 	[]*chatMsg.MsgPack
	errorItems 	[]error
	targetIs 	string
	finishChan  []chan bool
}

func (b *BroadcasterForPull) Initial() error {
	var (
		err error
	)
	if clusterConfig.AllHosts == nil {
		return fmt.Errorf("No Cluster Hosts")
	}
	for i := 0 ; i < len(clusterConfig.AllHosts) ; i ++ {
		if clusterConfig.AllHosts[i] == clusterConfig.Host {
			continue
		}
		nowClient := &CynicUClient.Client{}
		err = nowClient.Initial(clusterConfig.AllHosts[i], time.Second)
		if err != nil {
			return err
		}
		b.clients = append(b.clients, nowClient)
		b.respItems = append(b.respItems, nil)
		b.errorItems = append(b.errorItems, nil)
		b.finishChan = append(b.finishChan, make(chan bool, 1))
	}
	return nil
}

func (b *BroadcasterForPull) outBroad(i int) bool {
	return !(0 <= i && i < len(b.clients))
}

func (b *BroadcasterForPull) GetResp(i int) (*chatMsg.MsgPack,error) {
	if b.outBroad(i) {
		return nil,nil
	} else {
		select {
			case <-b.finishChan[i] :
		}
		return b.respItems[i],b.errorItems[i]
	}
}

func (b *BroadcasterForPull) pullItem(i int) {
	if !b.outBroad(i) {
		b.respItems[i], b.errorItems[i] = b.clients[i].Pull(&chatMsg.PullReq{From: b.targetIs})
		close(b.finishChan[i])
	}
}

func (b *BroadcasterForPull) SetTarget(targetIs string) {
	b.targetIs = targetIs
}

func (b *BroadcasterForPull) Size() int {
	return len(b.clients)
}

func (b *BroadcasterForPull) Do() {
	for i := 0 ; i < len(b.clients); i ++ {
		go b.pullItem(i)
	}
}