package Group

import (
	"github.com/dxyinme/Luka/Dao/GroupDao"
	"github.com/dxyinme/Luka/Dao/httpRouter/util"
	"github.com/gorilla/mux"
)

var (
	gdr = &DaoCall{}
)

func Initial(router *mux.Router) {
	gdr.dao = GroupDao.NewMIIGDImpl()

	groupInfoRouter := router.PathPrefix("/GroupInfo/").Subrouter()

	groupInfoRouter.HandleFunc("/test", util.TestRouter)
	groupInfoRouter.HandleFunc("/CreateGroup", gdr.CreateGroup)
	groupInfoRouter.HandleFunc("/DeleteGroup", gdr.DeleteGroup)
	groupInfoRouter.HandleFunc("/GetAllGroup", gdr.GetAllGroup)

}