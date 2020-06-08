package Keeper

import (
	"Luka/chatMsg"
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
	// 发送到客户端的channel
	writeCh *chan []byte
	// 读取客户端写入的channel
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
	go now.readTransform()
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
			log.Println(err)
			goto ERROR
		}
		fmt.Println("ws to:" + string(data))
	}
ERROR:
	u.Close()
}

// 将消息加入转发内容，传入UserPool的转发队列进行转发
func (u *User) readTransform() {
	var (
		msg []byte
		err error
	)
	for {
		msg,err = u.GetMessage()
		if err != nil {
			log.Printf("%s channel: %v\n", u.name, err)
			goto ERROR
		}
		textMsg := chatMsg.NewTextMsgUnmarshal(msg)
		// log.Println("keepUserPool:",keepUserPool)
		if textMsg == nil {
			log.Printf("[textMsg json error]")
			continue
		}
		select {
		case *keepUserPool.TextMsgCh <- *textMsg:{

		}
		case <- *keepUserPool.closeSign :{
			goto ERROR
		}
		}

	}
ERROR:
	UserPoolClose()
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
			//t := chatMsg.NewTextMsgUnmarshal(data)
			//log.Println(u.name , " : " , t)
			//if err != nil {
			//	log.Println("failed !")
			//}
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
	fmt.Println(string(s))
	select {
	case *(u.writeCh) <- s:{
		fmt.Println("success: " + string(s))
	}
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