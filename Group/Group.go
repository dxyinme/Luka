package Group

import (
	"fmt"
	"github.com/golang/glog"
	"sync"
)

type Group interface {
	Join(uid string) error
	Leave(uid string) error
	GetMaster() string
	SetMaster(uid string) error
	GetGroupName() string
}

type Impl struct {
	groupName	string
	masterUid 	string

	RWmu 		sync.RWMutex
	Members  	map[string]bool

	mu 			sync.Mutex
}

func (i *Impl) Join(uid string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	glog.Info("user [%s] join in group [%s]", uid, i.groupName)
	if _, ok := i.Members[uid]; ok {
		return fmt.Errorf("user [%s] has join in group [%s]", uid, i.groupName)
	}
	i.Members[uid] = true
	return nil
}

func (i *Impl) Leave(uid string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if _, ok := i.Members[uid]; !ok {
		return fmt.Errorf("user [%s] hasn't join in group [%s]", uid, i.groupName)
	}
	delete(i.Members, uid)
	return nil
}

func (i *Impl) GetMaster() string {
	i.RWmu.RLock()
	defer i.RWmu.RUnlock()
	return i.masterUid
}

func (i *Impl) SetMaster(uid string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	if _, ok := i.Members[uid]; !ok {
		return fmt.Errorf("user [%s] has not join in group [%s]", uid, i.groupName)
	}
	if uid == i.masterUid {
		return fmt.Errorf("user [%s] has been master of group [%s]", uid, i.groupName)
	}
	i.masterUid = uid
	return nil
}

func (i *Impl) GetGroupName() string {
	i.RWmu.RLock()
	defer i.RWmu.RUnlock()
	return i.groupName
}

func New(groupName, masterUid string) *Impl {
	members := make(map[string]bool)
	return &Impl{
		groupName: groupName,
		masterUid: masterUid,
		Members:   members,
	}
}