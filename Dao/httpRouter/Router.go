package httpRouter

import (
	"github.com/dxyinme/Luka/Dao/httpRouter/User"
	"github.com/dxyinme/Luka/Dao/httpRouter/util"
	"github.com/gorilla/mux"
)

var (
	Router = mux.NewRouter()
)


func Initial() {
	userRouter := Router.PathPrefix("/User").Subrouter()
	userRouter.HandleFunc("/test", util.TestRouter)
	User.Initial(userRouter)
}