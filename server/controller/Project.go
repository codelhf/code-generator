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

func ProjectInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.Project{}
	common.Bind(r, &temp)
	row := projectDao.Insert(temp)
	if !row {
		common.FailMsg(w, "save project failed")
		return
	}
	common.SuccessMsg(w, "save project success")
}

func ProjectUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.Project{}
	common.Bind(r, &temp)
	row := projectDao.Update(temp.Id, temp)
	if !row {
		common.FailMsg(w, "update project failed")
		return
	}
	common.SuccessMsg(w, "update project success")
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
		common.FailMsg(w, "delete project failed")
		return
	}
	common.SuccessMsg(w, "delete project success")
}

func ProjectLastId(w http.ResponseWriter, r *http.Request) {
	projectId := projectDao.LastId()
	common.SuccessData(w, projectId)
}