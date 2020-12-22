package main

import (
	"github.com/dxyinme/Luka/Dao/httpRouter"
	"github.com/golang/glog"
	"net/http"
)

// this is the micro server for DB-Operation
// after Glamorgann-kv https://github.com/Glamorgann/Glamorgann finished
// this will be used to realize the business between kv and user.
func main() {
	httpRouter.Initial()
	if err := http.ListenAndServe(":12777", httpRouter.Router); err != nil {
		glog.Fatal(err)
	}

}
