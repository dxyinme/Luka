package Auth

import (
	"github.com/dxyinme/LukaComm/Auth"
)

type SignImpl struct {

}

func (s SignImpl) SignUp(user *Auth.UserInfo) (err error) {
	panic("implement me")
}

func (s SignImpl) Login(uid, password string) (err error) {
	panic("implement me")
}

