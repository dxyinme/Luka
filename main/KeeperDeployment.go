package main

import (
	"flag"
	"net/http"

	"github.com/dxyinme/Luka/Keeper"
	"github.com/dxyinme/Luka/util"
	"github.com/golang/glog"
)



func InitialKeeper() {
	flag.StringVar(&util.KeeperName, "keeper", "test", "this keeper's name.")
	flag.StringVar(&util.KeeperUrl, "keeperUrl", "127.0.0.1:10137", "this keeper's url.")
	flag.StringVar(&util.MasterUrl, "masterUrl", "127.0.0.1:6965", "this master's url.")
}

// 一个 Keeper 有且只能有一个 Connector
func main() {
	InitialKeeper()
	flag.Parse()
	defer glog.Flush()
	newKeeper := Keeper.NewConnector(
		util.KeeperName,
		util.KeeperUrl,
		// 跨域
		func(r *http.Request) bool {
			return true
		})
	isRegister := newKeeper.Register(util.MasterUrl)

	if isRegister {
		glog.Infof("keeper %s is register success", util.KeeperName)
	} else {
		glog.Infof("have not a master , single keeper %s is working", util.KeeperName)
	}

	http.HandleFunc("/ConnectIt", newKeeper.ConnectIt)
	if err := http.ListenAndServe(":10137", nil); err != nil {
		glog.Fatal("ListenAndServe:", err)
	}
}
