package UserDao

import (
	"flag"
	"github.com/dxyinme/LukaComm/KvDao"
	"github.com/gomodule/redigo/redis"
)

var (
	UserGroupRedisHost = flag.String("UserGroupRedisHost",
		"127.0.0.1:6379", "redis host for user-group-info")
)

type UGDImpl struct {
	dao *KvDao.RedisDao
}

func NewUGDImpl() *UGDImpl {
	return &UGDImpl{dao: KvDao.NewRedisKv(*UserGroupRedisHost)}
}

func (U *UGDImpl) JoinGroupDao(uid string, groupName string) error {
	return U.dao.Insert(uid, groupName)
}

func (U *UGDImpl) LeaveGroupDao(uid string, groupName string) error {
	return U.dao.Remove(uid, groupName)
}

func (U *UGDImpl) CreateGroupDao(uid string, groupName string) error {
	return U.dao.Insert(uid, groupName)
}

func (U *UGDImpl) DeleteGroupDao(uid string, groupName string) error {
	return U.dao.Remove(uid, groupName)
}

func (U *UGDImpl) GetGroupNameList(uid string) ([]string, error) {
	return redis.Strings(U.dao.GetMembers(uid))
}
