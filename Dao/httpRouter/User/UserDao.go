package User

import (
	"github.com/dxyinme/Luka/Dao/UserDao"
	"github.com/dxyinme/Luka/Dao/httpRouter/util"
	"github.com/gorilla/mux"
)

var (
	gdr = &GroupDaoCall{}
)

func Initial(router *mux.Router) {
	gdr.userGroupDao = UserDao.NewUGDImpl()

	groupRouter := router.PathPrefix("/Group/").Subrouter()

	groupRouter.HandleFunc("/test", util.TestRouter)

	groupRouter.HandleFunc("/JoinGroupDao", gdr.JoinGroupDao)
	groupRouter.HandleFunc("/CreateGroupDao", gdr.CreateGroupDao)
	groupRouter.HandleFunc("/DeleteGroupDao", gdr.DeleteGroupDao)
	groupRouter.HandleFunc("/GetGroupNameList", gdr.GetGroupNameList)
	groupRouter.HandleFunc("/LeaveGroupDao", gdr.LeaveGroupDao)

}