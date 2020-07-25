package vio

import (
	"fmt"
	"github.com/dxyinme/Luka/chatMsg"
	"github.com/golang/glog"
	"sync"
)

type PoolVio struct {
	mp        	map[string] ChVio
	MsgCh 		*chan chatMsg.Msg
	closeSign 	*chan byte
	isClosed  	bool
	mutex     	sync.Mutex
}


// PoolVio 单个系统里面允许出现的vioPool数量最多不能超过1个
var vioPool *PoolVio

func InitPool(Size int) *PoolVio {
	tmp1 := make(chan chatMsg.Msg, Size)
	tmp2 := make(chan byte, 1)
	vioPool = &PoolVio{
		mp:        	make(map[string] ChVio),
		MsgCh: 		&tmp1,
		closeSign: 	&tmp2,
		isClosed:  	false,
	}
	go reSend()
	glog.Info("Pool initial finished")
	return vioPool
}

// 删除一个在线连接
func (vp *PoolVio) DeleteItem(name string) error {
	if vp.mp[name] == nil {
		return fmt.Errorf("%s is not connected", name)
	}
	vp.mp[name] = nil
	glog.Infof("delete %s success.", name)
	return nil
}

// 增加/更新 一个连接
func (vp *PoolVio) AddItem(item ChVio) error {
	vp.mp[item.Name()] = item
	return nil
}

// 获取name的对应Item
func (vp *PoolVio) GetItem(name string) ChVio {
	return vp.mp[name]
}

// ChVio上传 Msg 到转发池
func (vp *PoolVio) Upload(msg chatMsg.Msg) error {
	select {
	case *vp.MsgCh <- msg:
		{
			return nil
		}
	case <-*vp.closeSign:
		{
			return fmt.Errorf("vioPool is closed")
		}
	}
}


// 消息转发器
func reSend() {
	var (
		textData chatMsg.Msg
	)
	for {
		select {
		case textData = <-(*vioPool.MsgCh):
			{
				glog.Info(textData)
				if target, ok := vioPool.mp[textData.GetTarget()]; ok && target != nil {
					errAdd := target.AddMessage(textData)
					if errAdd != nil {
						glog.Infof("reSend error:%v\n", errAdd)
					}
				} else {
					glog.Infof("user %s is not in this keeper , message update\n", textData.GetTarget())
				}
			}
		case <-*vioPool.closeSign:
			{
				goto ERROR
			}
		}
	}
ERROR:
	PoolClose()
}

func PoolClose() error {
	glog.Infof("Pool Close")
	var err error = nil
	if vioPool == nil {
		return fmt.Errorf(" Pool is <nil> ")
	}
	vioPool.mutex.Lock()
	if !vioPool.isClosed {
		vioPool.isClosed = true
		close(*vioPool.closeSign)
	}
	vioPool.mutex.Unlock()
	return err
}