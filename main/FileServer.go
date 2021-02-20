package main

import (
	"flag"
	"github.com/dxyinme/Luka/FileServer"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"net/http"
)

// FileServer for file transport
var (
	Router = mux.NewRouter()
	port = flag.String("port", ":10505", "port this fileServer listening")
)

func main() {
	flag.Parse()
	FileServer.Initial(Router.PathPrefix("/api").Subrouter())
	if err := http.ListenAndServe(*port, Router); err != nil {
		glog.Fatal(err)
	}
}