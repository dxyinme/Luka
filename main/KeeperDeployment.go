package main

import (
	"flag"
	"github.com/dxyinme/Luka/Keeper"
	"github.com/dxyinme/Luka/util/config"
	"github.com/golang/glog"
	"net/http"
)



func InitialKeeper() {
	flag.StringVar(&config.KeeperName, "keeper", "test", "this keeper's name.")
	flag.StringVar(&config.KeeperUrl, "keeperUrl", "127.0.0.1:10137", "this keeper's url.")
	flag.StringVar(&config.MasterUrl, "masterUrl", "127.0.0.1:6965", "this master's url.")
}

// 一个 Keeper 有且只能有一个 Connector
func main() {
	InitialKeeper()
	flag.Parse()
	defer glog.Flush()
	newKeeper := Keeper.NewConnector(
		config.KeeperName,
		config.KeeperUrl,
		// 跨域
		func(r *http.Request) bool {
			return true
		})
	config.HasMaster = newKeeper.Register(config.MasterUrl)

	if config.HasMaster {
		glog.Infof("keeper %s is register success", config.KeeperName)
	} else {
		glog.Infof("have not a master , single keeper %s is working", config.KeeperName)
	}

	http.HandleFunc("/ConnectIt", newKeeper.ConnectIt)
	if err := http.ListenAndServe(":10137", nil); err != nil {
		glog.Fatal("ListenAndServe:", err)
	}
}
