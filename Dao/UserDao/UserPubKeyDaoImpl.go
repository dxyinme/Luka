package UserDao

import (
	"flag"
	"github.com/dxyinme/LukaComm/KvDao"
	"github.com/gomodule/redigo/redis"
)

var (
	UserPubKeyRedisHost = flag.String("UserPubKeyRedisHost",
		"127.0.0.1:6379", "redis host for user-pubKey-info")
)

type UserPubKeyDaoImpl struct {
	dao *KvDao.RedisDao
}

func NewUserPubKeyDaoImpl() *UserPubKeyDaoImpl {
	return &UserPubKeyDaoImpl{
		dao: KvDao.NewRedisKv(*UserPubKeyRedisHost),
	}
}

func (u *UserPubKeyDaoImpl) SetUserPubKey(uid string, pubKey []byte) (err error) {
	err = u.dao.Set(uid, pubKey)
	return
}

func (u *UserPubKeyDaoImpl) GetUserPubKey(uid string) (pubKey []byte, err error) {
	pubKey, err = redis.Bytes(u.dao.Get(uid))
	return
}
