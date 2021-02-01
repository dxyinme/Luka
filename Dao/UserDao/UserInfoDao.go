package UserDao

import (
	"github.com/dxyinme/LukaComm/Auth"
)

type UserInfoDao interface {
	RegisterNewUser(info *Auth.UserInfo) error
	ChangeUserInfo(info *Auth.UserInfo, oldInfo *Auth.UserInfo) error
	Login(code Auth.UserPassword) error
}