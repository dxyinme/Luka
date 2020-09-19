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
}

func (b *BroadcasterForPull) Initial() error {
	var (
		err error
	)
	if clusterConfig.AllHosts == nil {
		return fmt.Errorf("%s","No Cluster Hosts")
	}
	for i := 1 ; i < len(clusterConfig.AllHosts) ; i ++ {
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
	}
	return nil
}

func (b *BroadcasterForPull) pullItem(i int) {
	b.respItems[i], b.errorItems[i] = b.clients[i].Pull(&chatMsg.Ack{From: b.targetIs})
}

func (b *BroadcasterForPull) Do() {
	for i := 0 ; i < len(b.clients); i ++ {
		go b.pullItem(i)
	}
}