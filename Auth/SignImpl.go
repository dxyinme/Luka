package Auth

import (
	"github.com/dxyinme/Luka/User"
)

type SignImpl struct {

}

func (s SignImpl) SignUp(user *User.User) (err error) {
	panic("implement me")
}

func (s SignImpl) Login(uid, password string) (err error) {
	panic("implement me")
}

