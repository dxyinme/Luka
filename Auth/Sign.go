package Auth

import "github.com/dxyinme/LukaComm/Auth"

type Sign interface {
	SignUp(user *Auth.UserInfo) (err error)
	Login(uid, password string) (err error)
}