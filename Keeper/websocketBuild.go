package Keeper

import (
	MSA "github.com/dxyinme/Luka/proto/MasterServerApi"
	"github.com/dxyinme/Luka/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/websocket"
)

type Connector struct {
	keeperName      string
	keeperUrl       string
	upgrade         websocket.Upgrader
	registerSuccess bool
}

// 用于初始化Keeper的表 userPool 以及他们的连接
func NewConnector(keeperNameNow string, keeperUrlNow string, checkOrigin func(r *http.Request) bool) *Connector {
	defer glog.Info("NewConnector build finished")
	glog.Infof("this keeper's name is %s", keeperNameNow)
	uploadChan := InitUserPool()
	err := InitInformer(uploadChan)
	if err != nil {
		glog.Errorln(err)
	}
	return &Connector{
		keeperName: keeperNameNow,
		keeperUrl:  keeperUrlNow,
		upgrade: websocket.Upgrader{
			CheckOrigin: checkOrigin,
		},
		registerSuccess: false,
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

// 将 keeper 注册到master-server上 url 为 master-server 的url
func (cot *Connector) Register(url string) bool {
	var (
		err  error
		conn *grpc.ClientConn
	)
	conn, err = grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := MSA.NewMasterServiceApiClient(conn)
	res, err := client.KeeperAdd(context.Background(),
		&MSA.KeeperAddReq{
			Name:      cot.keeperName,
			KeeperUrl: cot.keeperUrl,
		})
	if res == nil {
		glog.Warningln("rpc response is nil")
		return false
	}
	if err != nil {
		glog.Warningf("rpc error : %v", err)
		return false
	}

	return res.GetStatus() == util.OK
}
