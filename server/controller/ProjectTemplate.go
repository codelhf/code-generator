package controller

import (
	"code-generator-go/server/common"
	"code-generator-go/server/db"
	"code-generator-go/server/model"
	"net/http"
	"strings"
)

var projectTemplateDao model.ProjectTemplateDao

func ProjectTemplateList(w http.ResponseWriter, r *http.Request) {
	pageInfo := db.BuildPageInfo(r)
	projectId := r.FormValue("projectId")
	dataList, total := projectTemplateDao.List(pageInfo.PageNum, pageInfo.PageSize, projectId)
	pageInfo.List = dataList
	pageInfo.Total = total
	common.SuccessData(w, pageInfo)
}

func ProjectTemplateSelect(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data := projectTemplateDao.Select(id)
	common.SuccessData(w, data)
}

func ProjectTemplateCheck(w http.ResponseWriter, r *http.Request) {
	temp := model.ProjectTemplate{}
	common.Bind(r, &temp)
	has := projectTemplateDao.Check(temp)
	if has {
		common.FailMsg(w, "ProjectTemplate Exists")
		return
	}
	common.Success(w)
}

func ProjectTemplateInsert(w http.ResponseWriter, r *http.Request) {
	ProjectTemplateCheck(w, r)
	temp := model.ProjectTemplate{}
	common.Bind(r, &temp)
	row := projectTemplateDao.Insert(temp)
	if !row {
		common.FailMsg(w, "Save ProjectTemplate Failed")
		return
	}
	common.SuccessMsg(w, "Save ProjectTemplate Success")
}

func ProjectTemplateUpdate(w http.ResponseWriter, r *http.Request) {
	ProjectTemplateCheck(w, r)
	temp := model.ProjectTemplate{}
	common.Bind(r, &temp)
	row := projectTemplateDao.Update(temp.Id, temp)
	if !row {
		common.FailMsg(w, "Update ProjectTemplate Failed")
		return
	}
	common.SuccessMsg(w, "Update ProjectTemplate Success")
}

func ProjectTemplateDelete(w http.ResponseWriter, r *http.Request) {
	projectId := r.FormValue("projectId")
	ids := r.FormValue("ids")
	idList := strings.Split(ids, ",")
	row := projectTemplateDao.Delete(projectId, idList)
	if !row {
		common.FailMsg(w, "Delete ProjectTemplate Failed")
		return
	}
	common.SuccessMsg(w, "Delete ProjectTemplate Success")
}
