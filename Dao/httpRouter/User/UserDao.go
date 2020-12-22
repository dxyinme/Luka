package User

import (
	"github.com/dxyinme/Luka/Dao/httpRouter"
	"github.com/gorilla/mux"
)

var (
	Router = mux.NewRouter()

	GroupDaoRouter = mux.NewRouter()
	gdr = &GroupDaoCall{}

	AuthDaoRouter = mux.NewRouter()
)

func Initial() {
	// groupDao init
	GroupDaoRouter.HandleFunc("/JoinGroupDao", gdr.JoinGroupDao)
	GroupDaoRouter.HandleFunc("/CreateGroupDao", gdr.CreateGroupDao)
	GroupDaoRouter.HandleFunc("/DeleteGroupDao", gdr.DeleteGroupDao)
	GroupDaoRouter.HandleFunc("/GetGroupNameList", gdr.GetGroupNameList)
	GroupDaoRouter.HandleFunc("/LeaveGroupDao", gdr.LeaveGroupDao)

	// authDao init
	// todo
	Router.HandleFunc("/test", httpRouter.TestRouter)
	Router.Handle("/group", GroupDaoRouter)
	Router.Handle("/auth", AuthDaoRouter)
}