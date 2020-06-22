package main

import (
	"flag"
	"net/http"

	"github.com/dxyinme/Luka/Keeper"
	"github.com/golang/glog"
)

var(
	keeperName string
)

func Initial() {
	flag.StringVar(&keeperName, "keeper","test","this keeper's name.")
}

// 一个 Keeper 有且只能有一个 Connector
func main() {
	Initial()
	flag.Parse()
	defer glog.Flush()
	newKeeper := Keeper.NewConnector(
		keeperName,
		// 跨域
		func(r *http.Request) bool {
			return true
		})
	http.HandleFunc("/ConnectIt", newKeeper.ConnectIt)
	if err := http.ListenAndServe(":10137", nil); err != nil {
		glog.Fatal("ListenAndServe:", err)
	}
}
