package Keeper

import (
	"fmt"
	"sync"
	"time"

	"github.com/dxyinme/Luka/chatMsg"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
)

const (
	// 用户信息管道大小
	channelSize = 100
)

// User , 用于创建一个websocket连通长连接。
type User struct {
	name 		string
	// 发送到客户端的channel
	writeCh 	*chan chatMsg.Msg
	// 读取客户端写入的channel
	readCh    	*chan []byte
	closeSign 	*chan byte
	ws        	*websocket.Conn
	isClosed  	bool
	mutex     	sync.Mutex
}

// 新建一个User
func NewUser(name string, Ws *websocket.Conn) *User {
	tmp1 := make(chan chatMsg.Msg, channelSize)
	tmp2 := make(chan []byte, channelSize)
	tmp3 := make(chan byte, 1)
	now := &User{
		name:      name,
		writeCh:   &tmp1,
		readCh:    &tmp2,
		closeSign: &tmp3,
		ws:        Ws,
		isClosed:  false,
	}
	go now.readLoop()
	go now.writeLoop()
	go now.readTransform()
	return now
}

// 将writeCh中的内容发送到客户端
func (u *User) writeLoop() {
	var (
		data chatMsg.Msg
		dataByte []byte
		err  error
	)
	for {
		select {
		case data = <-*u.writeCh:
		case <-*u.closeSign:
			{
				goto ERROR
			}
		}
		dataByte,err = data.Marshal()
		if err = u.ws.WriteMessage(websocket.TextMessage, dataByte); err != nil {
			glog.Errorln(err)
			goto ERROR
		}
	}
ERROR:
	u.Close()
}

func (u *User) Name() string {
	return u.name
}

// 将消息加入转发内容，传入UserPool的转发队列进行转发
func (u *User) readTransform() {
	var (
		msg []byte
		err error
	)
	for {
		msg, err = u.GetMessage()
		if err != nil {
			glog.Errorf("%s channel: %v\n", u.name, err)
			goto ERROR
		}
		userMsg := chatMsg.NewUserMsgByte(msg)
		err = keepUserPool.Upload(userMsg)
		if err != nil {
			glog.Infof("Upload msg error : %v ", err)
			goto ERROR
		}
	}
ERROR:
	u.Close()
}

// 接收长连接的消息,保存到readCh
func (u *User) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = u.ws.ReadMessage(); err != nil {
			goto ERROR
		}
		select {
		case *u.readCh <- data:
			{
				//glog.Infof("%s",string(data))
			}
		case <-*u.closeSign:
			{
				goto ERROR
			}
		}
	}
ERROR:
	u.Close()
}

//将信息写入writeCh
func (u *User) AddMessage(s chatMsg.Msg) error {
	select {
	case *(u.writeCh) <- s:
		{
			glog.Info("success: [" + s.GetFrom() + "] to [" + s.GetTarget() + "] time is " + time.Now().String())
		}
	case <-(*u.closeSign):
		return fmt.Errorf("write error : connection is closed")
	}
	return nil
}

//从readCh中读取信息
func (u *User) GetMessage() ([]byte, error) {
	var (
		data []byte = nil
		err  error  = nil
	)
	select {
	case data = <-(*u.readCh):
	case <-(*u.closeSign):
		{
			err = fmt.Errorf("read error : connection is closed")
		}
	}
	return data, err
}

// 该用户使用完毕， 关闭长连接
func (u *User) Close() error {
	var err error = nil
	u.mutex.Lock()
	if !u.isClosed {
		u.isClosed = true
		close(*u.closeSign)
		err = u.ws.Close()
	}
	u.mutex.Unlock()
	return err
}

// 开启User
func (u *User) Serve() error {
	select {
	case <-*u.closeSign:
	}
	return nil
}