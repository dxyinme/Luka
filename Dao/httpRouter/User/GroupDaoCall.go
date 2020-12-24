package User

import (
	"errors"
	"github.com/dxyinme/Luka/Dao/UserDao"
	"github.com/dxyinme/Luka/Dao/httpRouter/util"
	"net/http"
)

type GroupDaoCall struct {
	userGroupDao UserDao.UserGroupDao
}

func (g *GroupDaoCall) JoinGroupDao(w http.ResponseWriter, r *http.Request) {
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


	err = g.userGroupDao.JoinGroupDao(uid, groupName)
	if util.SolveError(reply, err) {
		goto RET
	}
RET:
	_,_ = w.Write(util.ReParseBody(reply))
}

func (g *GroupDaoCall) LeaveGroupDao(w http.ResponseWriter, r *http.Request) {
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

	err = g.userGroupDao.LeaveGroupDao(uid, groupName)
	if util.SolveError(reply, err) {
		goto RET
	}
RET:
	_,_ = w.Write(util.ReParseBody(reply))
}

func (g *GroupDaoCall) CreateGroupDao(w http.ResponseWriter, r *http.Request) {
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

	err = g.userGroupDao.CreateGroupDao(uid, groupName)
	if util.SolveError(reply, err) {
		goto RET
	}
RET:
	_,_ = w.Write(util.ReParseBody(reply))
}

func (g *GroupDaoCall) DeleteGroupDao(w http.ResponseWriter, r *http.Request) {
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

	err = g.userGroupDao.DeleteGroupDao(uid, groupName)
	if util.SolveError(reply, err) {
		goto RET
	}
RET:
	_,_ = w.Write(util.ReParseBody(reply))
}

func (g *GroupDaoCall) GetGroupNameList(w http.ResponseWriter, r *http.Request) {
	var(
		uid string
		err error
		groupNamesList []string
	)

	err = r.ParseForm()
	reply := make(map[string]interface{})
	if util.SolveError(reply, err) {
		goto RET
	}

	uid = r.PostForm.Get("Uid")
	if uid == "" {
		util.SolveError(reply, errors.New("post form empty"))
		goto RET
	}

	groupNamesList, err = g.userGroupDao.GetGroupNameList(uid)
	if util.SolveError(reply, err) {
		goto RET
	}
	reply["GroupNameList"] = groupNamesList
RET:
	_,_ = w.Write(util.ReParseBody(reply))
}