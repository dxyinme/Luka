package httpRouter

import (
	httpGroup "github.com/dxyinme/Luka/Dao/httpRouter/Group"
	httpUser "github.com/dxyinme/Luka/Dao/httpRouter/User"
	httpUtil "github.com/dxyinme/Luka/Dao/httpRouter/util"
	"github.com/gorilla/mux"
)


var (
	Router = mux.NewRouter()
)


func Initial() {
	userRouter := Router.PathPrefix("/User").Subrouter()
	userRouter.HandleFunc("/test", httpUtil.TestRouter)
	httpUser.Initial(userRouter)
	groupRouter := Router.PathPrefix("/Group").Subrouter()
	groupRouter.HandleFunc("/test", httpUtil.TestRouter)
	httpGroup.Initial(groupRouter)
}
