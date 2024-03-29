package controller

import (
	"code-generator-go/server/common"
	"code-generator-go/server/db"
	"code-generator-go/server/model"
	"net/http"
	"strings"
)

var projectDao model.ProjectDao
var projectDbDao model.ProjectDbDao
var projectTable model.ProjectTableDao
var projectTemplate model.ProjectTemplateDao

func ProjectList(w http.ResponseWriter, r *http.Request) {
	pageInfo := db.BuildPageInfo(r)
	name := r.FormValue("name")
	desc := r.FormValue("desc")
	dataList, total := projectDao.List(pageInfo.PageNum, pageInfo.PageSize, name, desc)
	pageInfo.List = dataList
	pageInfo.Total = total
	common.SuccessData(w, pageInfo)
}

func ProjectSelect(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data := projectDao.Select(id)
	common.SuccessData(w, data)
}

func ProjectCheck(w http.ResponseWriter, r *http.Request) {
	temp := model.Project{}
	common.Bind(r, &temp)
	if !projectCheck(w, temp) {
		return
	}
	common.Success(w)
}

func projectCheck(w http.ResponseWriter, temp model.Project) bool {
	has := projectDao.Check(temp)
	if has {
		common.FailMsg(w, "Project Exists")
		return false
	}
	return true
}

func ProjectInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.Project{}
	common.Bind(r, &temp)
	if !projectCheck(w, temp) {
		return
	}
	row := projectDao.Insert(temp)
	if !row {
		common.FailMsg(w, "Save Project Failed")
		return
	}
	userLogDao.Update(model.LAST_PROJECT, temp.Id)
	common.SuccessMsg(w, "Save Project Success")
}

func ProjectUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.Project{}
	common.Bind(r, &temp)
	if !projectCheck(w, temp) {
		return
	}
	row := projectDao.Update(temp.Id, temp)
	if !row {
		common.FailMsg(w, "Update Project Failed")
		return
	}
	userLogDao.Update(model.LAST_PROJECT, temp.Id)
	common.SuccessMsg(w, "Update Project Success")
}

func ProjectDelete(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	idList := strings.Split(ids, ",")
	for i := 0; i < len(idList); i++ {
		projectId := idList[i]
		projectDbDao.Delete(projectId, nil)
		projectTable.Delete(projectId, nil)
		projectTemplate.Delete(projectId, nil)
	}
	row := projectDao.Delete(idList)
	if !row {
		common.FailMsg(w, "Delete Project Failed")
		return
	}
	common.SuccessMsg(w, "Delete Project Success")
}

func ProjectLastId(w http.ResponseWriter, r *http.Request) {
	projectId := userLogDao.Select(model.LAST_PROJECT)
	common.SuccessData(w, projectId)
}
