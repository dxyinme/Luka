package main

import (
	"flag"
	"github.com/dxyinme/Luka/Dao/httpRouter"
	"github.com/golang/glog"
	"net/http"
)

// this is the micro server for DB-Operation
// after Glamorgann-kv https://github.com/Glamorgann/Glamorgann finished
// this will be used to realize the business between kv and user.

var (
	DBSAddr = flag.String("DBServer", ":12777", "DBServer port")
)

func main() {
	flag.Parse()
	httpRouter.Initial()
	if err := http.ListenAndServe(*DBSAddr, httpRouter.Router); err != nil {
		glog.Fatal(err)
	}
}
