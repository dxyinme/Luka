package Keeper

import (
	"golang.org/x/net/websocket"
	"log"
	"sync"
	"time"
)

const (
	channelSize = 100
)
type User struct {
	name string
	ch *chan string
}
// 新建一个User
func NewUser(name string) *User{
	tmp := make(chan string, channelSize)
	return &User{
		name: name,
		ch:  &tmp,
	}
}

func (u *User) PostRecv(ws *websocket.Conn) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func () {
		for {
			var msg string
			var ok bool
			select {
			case msg, ok = <-(*u.ch):
			}
			if !ok {
				break
			}
			log.Println("send : " + msg)
			for i := 1 ; i <= 5; i ++ {
				errRecv := websocket.Message.Send(ws, msg)
				if errRecv != nil {
					log.Printf("Can't send because %v" ,errRecv)
				} else {
					break
				}
				time.Sleep(time.Second)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

//增加信息
func (u *User) AddMessage(s string) {
	*(u.ch) <- s
}