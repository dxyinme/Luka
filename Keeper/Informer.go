package Keeper

import (
	"github.com/dxyinme/Luka/chatMsg"
	MSA "github.com/dxyinme/Luka/proto/MasterServerApi"
	"github.com/dxyinme/Luka/util"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

)

var (
	updateMessage *chan chatMsg.Msg
)


func InitInformer(msgChan *chan chatMsg.Msg) error {
	updateMessage = msgChan
	return util.NewTimeTask( "*/5 * * * * ?" , func() {
		now := pack()
		conn, err := grpc.Dial(util.MasterUrl, grpc.WithInsecure())
		if err != nil {
			glog.Error(err)
		}
		defer conn.Close()
		client := MSA.NewMasterServiceApiClient(conn)
		var packMsg [][]byte
		for i := 0; i < len(now); i ++ {
			nowBytes,err := now[i].Marshal()
			if err != nil {
				glog.Errorf("No. %d , msg is : %v", i, now[i])
				continue
			}
			packMsg = append(packMsg, nowBytes)
		}
		resp, err := client.KeeperSync(context.Background(), &MSA.KeeperSyncReq{
			PackMsg: packMsg,
		})
		if err != nil {
			glog.Error(err)
		}
		glog.Info(resp)
	})
}

func pack() []chatMsg.Msg {
	nowlen := len(*updateMessage)
	var upSendPack []chatMsg.Msg
	for i := 0; i < nowlen; i++ {
		msg, ok := <-*updateMessage
		if ok {
			upSendPack = append(upSendPack, msg)
		} else {
			glog.Info("updateMessageChan is closed")
		}
	}
	return upSendPack
}