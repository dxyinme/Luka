package assignCli

import (
	"fmt"
	"github.com/dxyinme/Luka/assigneerServer/AssignUtil"
	"github.com/dxyinme/LukaComm/Assigneer"
	"github.com/dxyinme/LukaComm/util"
	"log"
)

func PrintGetAllKeeperInfoRsp(o *Assigneer.ClusterRsp) {
	if o == nil {
		return
	}
	var now = &AssignUtil.KeeperList{}
	err := util.IJson.Unmarshal(o.RspInfo, now)
	if err != nil {
		log.Println(err)
		return
	}
	// title.
	fmt.Printf("%s\t%s\t%s\t%s\t\n",
		"KeeperId", "MsgRecv", "MsgSend", "MsgNotLocal")
	for i := 0 ; i < len(now.KeeperId); i ++ {
		fmt.Printf("%d\t%d\t%d\t%d\t\n",
			now.KeeperId[i], now.Lis[i].MsgRecv, now.Lis[i].MsgSend, now.Lis[i].MsgNotLocal)
	}
}
