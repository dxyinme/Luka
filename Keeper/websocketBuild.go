package Keeper

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/websocket"
)

type Connector struct {
	keeperName 	string
	userPool 	*UserPool
	upgrade  	websocket.Upgrader
}

// 用于初始化Keeper的表 userPool 以及他们的连接
func NewConnector(keeperNameNow string, checkOrigin func(r *http.Request) bool) *Connector {
	defer glog.Info("NewConnector build finished")
	glog.Infof("this keeper's name is %s \n", keeperNameNow)

	return &Connector{
		keeperName: keeperNameNow,
		userPool: InitUserPool(),
		upgrade: websocket.Upgrader{
			CheckOrigin: checkOrigin,
		},
	}
}

// http request 登录处理，我们将其升级成为 websocket
func (cot *Connector) ConnectIt(w http.ResponseWriter, r *http.Request) {
	var (
		conn *websocket.Conn
		err  error
		user *User
		name string
		// data []byte
	)
	err = r.ParseForm()
	name = r.Form.Get("name")
	// upgrade to websocket
	if conn, err = cot.upgrade.Upgrade(w, r, nil); err != nil {
		return
	}
	user = NewUser(name, conn)
	AddUser(user)
	defer DeleteUser(name)
	if err = user.Serve(); err != nil {
		glog.Errorln("serve error:", err)
	}
	glog.Infof("%s close websocket error : %v", name, user.Close())
}
