package AuthServer

import (
	"context"
	"github.com/dxyinme/Luka/Dao/UserDao"
	"github.com/dxyinme/LukaComm/Auth"
	"github.com/golang/glog"
)

type Server struct {
	userPubKeyDao UserDao.UserPubKeyDao
}

func (s *Server) ChangeInfo(context.Context, *Auth.ChangeInfoReq) (*Auth.ChangeInfoRsp, error) {
	panic("implement me")
}

func (s *Server) SetAuthPubKey(ctx context.Context,req *Auth.SetAuthPubKeyReq) (*Auth.SetAuthPubKeyRsp, error) {
	glog.Infof("SetAuthPubKey: %s, %v", req.Uid, req.AuthRsaPubKey)
	err := s.userPubKeyDao.SetUserPubKey(req.Uid, req.AuthRsaPubKey)
	rsp := &Auth.SetAuthPubKeyRsp{}
	if err != nil {
		rsp.ErrorMsg = err.Error()
	}
	return rsp, nil
}

func (s *Server) Login(ctx context.Context, code *Auth.UserPassword) (*Auth.LoginRsp, error) {
	panic("implement me")
}

func (s *Server) SignUp(ctx context.Context, userInfo *Auth.UserInfo) (*Auth.SignUpRsp, error) {
	panic("implement me")
}

// Get someone PubKey
func (s *Server) GetAuthPubKey(ctx context.Context, req *Auth.GetAuthPubKeyReq) (*Auth.GetAuthPubKeyRsp, error) {
	var (
		rsp = &Auth.GetAuthPubKeyRsp{}
		err error
	)
	rsp.AuthRsaPubKey, err = s.userPubKeyDao.GetUserPubKey(req.Uid)
	glog.Infof("GetAuthPubKey: %s, %v", req.Uid, rsp.AuthRsaPubKey)
	if err != nil {
		rsp.ErrorMsg = err.Error()
	}
	return rsp, nil
}

func (s *Server) Initial() {
	s.userPubKeyDao = UserDao.NewUserPubKeyDaoImpl()
}