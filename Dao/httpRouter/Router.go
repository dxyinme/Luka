package httpRouter

import (
	httpGroup "github.com/dxyinme/Luka/Dao/httpRouter/Group"
	httpUtil "github.com/dxyinme/Luka/Dao/httpRouter/util"
	"github.com/gorilla/mux"
)


var (
	Router = mux.NewRouter()
)


func Initial() {
	//userRouter := Router.PathPrefix("/User").Subrouter()
	//userRouter.HandleFunc("/test", util.TestRouter)
	//User.Initial(userRouter)
	groupRouter := Router.PathPrefix("/group").Subrouter()
	groupRouter.HandleFunc("/test", httpUtil.TestRouter)
	httpGroup.Initial(groupRouter)
}
