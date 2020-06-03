package Keeper

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)


type Connector struct {
	userPool *UserPool
	upgrade websocket.Upgrader
}

func NewConnector(checkOrigin func(r *http.Request) bool) *Connector{
	defer log.Println("NewConnector build finished")
	return &Connector{
		userPool: NewUserPool(),
		upgrade:  websocket.Upgrader{
			CheckOrigin: checkOrigin,
		},
	}
}

// http request 登录处理，我们将其升级成为 websocket
func (cot *Connector) ConnectIt(w http.ResponseWriter, r *http.Request) {
	var (
		conn *websocket.Conn
		err error
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
	user = NewUser(name,conn)
	cot.userPool.AddUser(user)
	defer cot.userPool.DeleteUser(name)
	if err = user.Serve(); err != nil {
		log.Println(err)
	}
	log.Println("close websocket :", user.Close())
}