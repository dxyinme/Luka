package main

import (
	"flag"
	"github.com/dxyinme/Luka/lived"
	"github.com/golang/glog"
	"net/http"
)

var (
	addr = flag.String("addr", ":21350", "")
)

func main() {
	flag.Parse()
	lived.Init()
	controller := lived.NewController()
	http.HandleFunc("/newRoom", controller.NewRoom)
	http.HandleFunc("/bootLive", controller.BootLive)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		glog.Fatal(err)
	}
}