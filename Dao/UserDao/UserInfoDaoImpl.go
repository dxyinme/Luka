package UserDao

import (
	"database/sql"
	"flag"
	"github.com/dxyinme/LukaComm/Auth"
)

var (
	userInfoRedisHost = flag.String("userInfoRedisHost", "", "userInfoRedisHost")
)

type UserInfoDaoImpl struct {
	sqlDB *sql.DB
}

func (u *UserInfoDaoImpl) RegisterNewUser(info *Auth.UserInfo) error {
	panic("implement me")
}

func (u *UserInfoDaoImpl) ChangeUserInfo(info *Auth.UserInfo, oldInfo *Auth.UserInfo) error {
	panic("implement me")
}

func (u *UserInfoDaoImpl) Login(code *Auth.UserPassword) error {
	panic("implement me")
}

func (u *UserInfoDaoImpl) Initial() {

}