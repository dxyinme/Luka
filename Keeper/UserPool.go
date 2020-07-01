package Keeper

import (
	"fmt"
	"sync"

	"github.com/dxyinme/Luka/chatMsg"
	"github.com/golang/glog"
)

const (
	MsgChannelSize = 1500
)

type UserPool struct {
	mp        	map[string]*User
	MsgCh 		*chan chatMsg.Msg
	closeSign 	*chan byte
	isClosed  	bool
	mutex     	sync.Mutex
}

var keepUserPool *UserPool

func InitUserPool() *UserPool {
	tmp1 := make(chan chatMsg.Msg, MsgChannelSize)
	tmp2 := make(chan byte, 1)
	keepUserPool = &UserPool{
		mp:        	map[string]*User{},
		MsgCh: 		&tmp1,
		closeSign: 	&tmp2,
		isClosed:  	false,
	}
	go reSend()
	glog.Info("UserPool initial finished")
	return keepUserPool
}

// 增加/更新 用户连接
func AddUser(user *User) {
	keepUserPool.mp[user.name] = user
	glog.Infof("user %s is login", user.name)
}

// 用户断开连接
func DeleteUser(name string) error {
	if keepUserPool.mp[name] == nil {
		return fmt.Errorf("%s is not connected", name)
	}
	keepUserPool.mp[name] = nil
	glog.Infof("delete %s success.", name)
	return nil
}

// 获取name的对应用户
func GetUser(name string) *User {
	return keepUserPool.mp[name]
}

// 消息转发器
func reSend() {
	var (
		textData chatMsg.Msg
	)
	for {
		select {
		case textData = <-(*keepUserPool.MsgCh):
			{
				glog.Info(textData)
				if target, ok := keepUserPool.mp[textData.GetTarget()]; ok && target != nil {
					errAdd := target.AddMessage(textData)
					if errAdd != nil {
						glog.Infof("reSend error:%v\n", errAdd)
					}
				} else {
					glog.Infof("user %s is not in this keeper , message update\n", textData.GetTarget())
				}
			}
		case <-*keepUserPool.closeSign:
			{
				goto ERROR
			}
		}
	}
ERROR:
	UserPoolClose()
}

func UserPoolClose() error {
	var err error = nil
	if keepUserPool == nil {
		return fmt.Errorf("keepUserPool is <nil>")
	}
	keepUserPool.mutex.Lock()
	if !keepUserPool.isClosed {
		keepUserPool.isClosed = true
		close(*keepUserPool.closeSign)
	}
	keepUserPool.mutex.Unlock()
	return err
}
