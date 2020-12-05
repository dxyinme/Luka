package Auth

import "github.com/dxyinme/Luka/User"

type Sign interface {
	SignUp(user User.User) (err error)
	Login(uid, password string) (err error)
}