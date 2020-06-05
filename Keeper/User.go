package Keeper

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

const (
	// 用户信息管道大小
	channelSize = 100
)

// User , 用于创建一个websocket连通长连接。
type User struct {
	name string
	writeCh *chan []byte
	readCh *chan []byte
	closeSign *chan byte
	ws *websocket.Conn
	isClosed bool
	mutex sync.Mutex
}

// 新建一个User
func NewUser(name string, Ws *websocket.Conn) *User{
	tmp1 := make(chan []byte, channelSize)
	tmp2 := make(chan []byte, channelSize)
	tmp3 := make(chan byte, 1)
	now := &User{
		name: 		name,
		writeCh:  	&tmp1,
		readCh:  	&tmp2,
		closeSign:  &tmp3,
		ws: 		Ws,
		isClosed: 	false,
	}
	go now.readLoop()
	go now.writeLoop()
	return now
}

// 将writeCh中的内容发送到客户端
func (u *User) writeLoop() {
	var (
		data []byte
		err error
	)
	for {
		select {
		case data = <- *u.writeCh:
		case <- *u.closeSign:{
			goto ERROR
		}
		}
		if err = u.ws.WriteMessage(websocket.TextMessage,data);err != nil {
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
		err error
	)
	for {
		if _,data,err = u.ws.ReadMessage(); err != nil {
			goto ERROR
		}
		select {
		case *u.readCh <- data:{
			log.Println(u.name + " : " + string(data))
		}
		case  <- *u.closeSign:{
			goto ERROR
		}
		}
	}
ERROR:
	u.Close()
}

//将信息写入writeCh
func (u *User) AddMessage(s []byte) error {
	select {
	case *(u.writeCh) <- s:
	case  <- (*u.closeSign):
		return fmt.Errorf("write error : connection is closed")
	}
	return nil
}

//从readCh中读取信息
func (u *User) GetMessage() ([]byte,error) {
	var (
		data []byte = nil
		err error = nil
	)
	select {
	case  data = <- (*u.readCh):
	case  <- (*u.closeSign):{
		err = fmt.Errorf("read error : connection is closed")
	}
	}
	return data , err
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
	case <- *u.closeSign:
	}
	return nil
}