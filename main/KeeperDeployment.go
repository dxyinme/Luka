package main

import (
	"flag"
	"net/http"

	"github.com/dxyinme/Luka/Keeper"
	"github.com/golang/glog"
)

var (
	keeperName string
	keeperUrl  string
	masterUrl  string
)

func InitialKeeper() {
	flag.StringVar(&keeperName, "keeper", "test", "this keeper's name.")
	flag.StringVar(&keeperUrl, "keeperUrl", "127.0.0.1:10137", "this keeper's url.")
	flag.StringVar(&masterUrl, "masterUrl", "127.0.0.1:6965", "this master's url.")
}

// 一个 Keeper 有且只能有一个 Connector
func main() {
	InitialKeeper()
	flag.Parse()
	defer glog.Flush()
	newKeeper := Keeper.NewConnector(
		keeperName,
		keeperUrl,
		// 跨域
		func(r *http.Request) bool {
			return true
		})
	isRegister := newKeeper.Register(masterUrl)

	if isRegister {
		glog.Infof("keeper %s is register success", keeperName)
	} else {
		glog.Infof("have not a master , single keeper %s is working", keeperName)
	}

	http.HandleFunc("/ConnectIt", newKeeper.ConnectIt)
	if err := http.ListenAndServe(":10137", nil); err != nil {
		glog.Fatal("ListenAndServe:", err)
	}
}
