package Keeper

import (
	"Luka/chatMsg"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

const (
	TextChannelSize = 1500
)

type UserPool struct {
	mp map[string] *User
	TextMsgCh *chan chatMsg.TextMsg
	closeSign *chan byte
	isClosed bool
	mutex sync.Mutex
}

var keepUserPool *UserPool

func InitUserPool() *UserPool {
	tmp1 := make(chan chatMsg.TextMsg,TextChannelSize)
	tmp2 := make(chan byte,1)
	keepUserPool = &UserPool{
		mp:			map[string]*User{},
		TextMsgCh:	&tmp1,
		closeSign:	&tmp2,
		isClosed:   false,
	}
	go keepUserPool.reSend()
	log.Println("UserPool initial finished")
	return keepUserPool
}

// 增加/更新 用户连接
func (kup *UserPool)AddUser(user *User) {
	kup.mp[user.name] = user
	log.Printf("user %s is login", user.name)
}

// 用户断开连接
func (kup *UserPool)DeleteUser(name string) error {
	if kup.mp[name] == nil {
		return fmt.Errorf("%s is not connected" , name)
	}
	kup.mp[name] = nil
	return nil
}

// 获取name的对应用户
func (kup *UserPool)GetUser(name string) *User {
	return kup.mp[name]
}


// 消息转发器
func (kup *UserPool) reSend() {
	var (
		textData chatMsg.TextMsg
	)
	for {
		select {
		case  textData = <- (*kup.TextMsgCh) : {
			log.Println(textData)
			textByte,errJson := json.Marshal(textData)
			if errJson != nil {
				log.Printf("[msg]:%s",errJson)
			}
			if target,ok := kup.mp[textData.Target]; ok && target != nil {
				errAdd := target.AddMessage(textByte)
				if errAdd != nil {
					log.Printf("[reSend error] %v\n",errAdd)
				}
			} else {
				log.Printf("user %s has logout\n",textData.Target)
			}
		}
		case <- *kup.closeSign: {
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