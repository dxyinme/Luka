package Group

import (
	"errors"
	"github.com/dxyinme/Luka/Dao/GroupDao"
	"github.com/dxyinme/Luka/Dao/httpRouter/util"
	"net/http"
)

type DaoCall struct {
	dao GroupDao.MainInfoInGroupDao
}


func (d *DaoCall) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var(
		uid string
		groupName string
		err error
	)
	err = r.ParseForm()
	reply := make(map[string]interface{})
	if util.SolveError(reply, err) {
		goto RET
	}

	uid = r.PostForm.Get("Uid")
	groupName = r.PostForm.Get("GroupName")
	if uid == "" || groupName == "" {
		util.SolveError(reply, errors.New("post form empty"))
		goto RET
	}

	err = d.dao.CreateGroup(groupName, uid)
	if util.SolveError(reply, err) {
		goto RET
	}
RET:
	_,_ = w.Write(util.ReParseBody(reply))
}

func (d *DaoCall) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	var(
		uid string
		groupName string
		err error
	)
	err = r.ParseForm()
	reply := make(map[string]interface{})
	if util.SolveError(reply, err) {
		goto RET
	}

	uid = r.PostForm.Get("Uid")
	groupName = r.PostForm.Get("GroupName")
	if uid == "" || groupName == "" {
		util.SolveError(reply, errors.New("post form empty"))
		goto RET
	}
	err = d.dao.DeleteGroup(groupName, uid)
	if util.SolveError(reply, err) {
		goto RET
	}
RET:
	_,_ = w.Write(util.ReParseBody(reply))
}

func (d *DaoCall) GetAllGroup(w http.ResponseWriter, r *http.Request) {
	var(
		groupNameList []string
		uidList []string
		err error
	)
	err = r.ParseForm()
	reply := make(map[string]interface{})
	if util.SolveError(reply, err) {
		goto RET
	}
	groupNameList, uidList, err = d.dao.GetAllGroup()
	if util.SolveError(reply, err) {
		goto RET
	}
	reply["GroupNameList"] = groupNameList
	reply["UidList"] = uidList

RET:
	_,_ = w.Write(util.ReParseBody(reply))
}