package GroupDao

import (
	"errors"
	"flag"
	"github.com/dxyinme/LukaComm/KvDao"
	"github.com/gomodule/redigo/redis"
	"sync"
)

var (
	MainInfoInGroupHost = flag.String("MainInfoInGroupHost",
		"127.0.0.1:6409", "MainInfoInGroupHost")
)

type MIIGDImpl struct {
	dao *KvDao.RedisDao
	mu sync.RWMutex
}

func (M *MIIGDImpl) CreateGroup(groupName string, uid string) (err error) {
	M.mu.Lock()
	defer M.mu.Unlock()
	var isExisted bool
	isExisted, err = M.dao.Exists(groupName)
	if isExisted {
		err = errors.New("groupName is existed")
		return
	} else {
		err = M.dao.Set(groupName, uid)
	}
	return
}

func (M *MIIGDImpl) DeleteGroup(groupName string, uid string) (err error) {
	M.mu.Lock()
	defer M.mu.Unlock()
	var isExisted bool
	isExisted, err = M.dao.Exists(groupName)
	if !isExisted {
		err = errors.New("groupName is not existed")
		return
	} else {
		err = M.dao.Delete(groupName)
	}
	return

}

// only be used in new keeper initial
func (M *MIIGDImpl) GetAllGroup() (groupNameList,uidList []string, err error) {
	M.mu.RLock()
	defer M.mu.RUnlock()
	var uidListInterface []interface{}
	groupNameList , uidListInterface , err = M.dao.Keys()
	if err != nil {
		return nil, nil, err
	}
	uidList, err = redis.Strings(uidListInterface, err)
	return
}

func (M *MIIGDImpl) GroupGetAllUser(groupName string) (uidList []string, err error) {
	M.mu.RLock()
	defer M.mu.RUnlock()
	var uidListInterface []interface{}
	uidListInterface, err = M.dao.GetMembers(groupName)
	if err != nil {
		return nil, err
	}
	uidList, err = redis.Strings(uidListInterface, err)
	return
}

func NewMIIGDImpl() *MIIGDImpl {
	return &MIIGDImpl{dao: KvDao.NewRedisKv(*MainInfoInGroupHost)}
}

