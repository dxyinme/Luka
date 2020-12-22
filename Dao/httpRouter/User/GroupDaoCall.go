package User

import (
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
	)
	body, err := util.ParseBody(r)
	reply := make(map[string]interface{})
	if util.SolveError(reply, err) {
		goto RET
	}
	uid = body["Uid"].(string)
	groupName = body["GroupName"].(string)
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
	)
	body, err := util.ParseBody(r)
	reply := make(map[string]interface{})
	if util.SolveError(reply, err) {
		goto RET
	}
	uid = body["Uid"].(string)
	groupName = body["GroupName"].(string)
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
	)
	body, err := util.ParseBody(r)
	reply := make(map[string]interface{})
	if util.SolveError(reply, err) {
		goto RET
	}
	uid = body["Uid"].(string)
	groupName = body["GroupName"].(string)
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
	)
	body, err := util.ParseBody(r)
	reply := make(map[string]interface{})
	if util.SolveError(reply, err) {
		goto RET
	}
	uid = body["Uid"].(string)
	groupName = body["GroupName"].(string)
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
		groupNamesList []string
	)
	body, err := util.ParseBody(r)
	reply := make(map[string]interface{})
	if util.SolveError(reply, err) {
		goto RET
	}
	uid = body["Uid"].(string)
	groupNamesList, err = g.userGroupDao.GetGroupNameList(uid)
	if util.SolveError(reply, err) {
		goto RET
	}
	reply["GroupNameList"] = groupNamesList
RET:
	_,_ = w.Write(util.ReParseBody(reply))
}

