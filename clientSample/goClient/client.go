package main

import (
	"flag"
	"github.com/dxyinme/Luka/chatMsg"
	"github.com/dxyinme/Luka/util"
	"github.com/gorilla/websocket"
	"log"
)
var(
	addr string
	name string
)

func main(){
	flag.StringVar(&addr, "KeeperUrl", "localhost:10137","keeper url")
	flag.StringVar(&name, "name", "test_","user name")
	url := "ws://" + addr + "/ConnectIt?name=" + name
	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial(url,nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:",err)
			return
		}
		var now = chatMsg.TmpMsg{}
		err = util.IJson.Unmarshal(message,&now)
		log.Println(err)
		log.Println(now)
	}
}
