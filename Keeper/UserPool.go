package Keeper

import (
	"github.com/dxyinme/Luka/vio"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/golang/glog"
)

const (
	MsgChannelSize = 1500
)

type UserPool struct {

}

var keepUserPool *vio.PoolVio

func InitUserPool() *chan chatMsg.Msg {
	keepUserPool = vio.InitPool(MsgChannelSize)
	return keepUserPool.GetUploadChan()
	//fmt.Println("init!")
}

// 增加/更新 用户连接
func AddUser(user *User) {
	err := keepUserPool.AddItem(user)
	if err != nil {
		glog.Errorf("user %s add/update error, because %v", user.Name(), err)
		return
	}
	glog.Infof("user %s is login", user.Name())
}


// 用户断开连接
func DeleteUser(name string) error {
	err := keepUserPool.DeleteItem(name)
	if err != nil {
		glog.Errorf("delete %s error. because %v", name, err)
		return err
	}
	glog.Infof("delete %s success.", name)
	return nil
}
