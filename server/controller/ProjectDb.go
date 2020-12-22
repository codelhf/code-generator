package controller

import (
	"code-generator-go/server/common"
	"code-generator-go/server/model"
	"net/http"
)

func ProjectDbSelect(w http.ResponseWriter, r *http.Request) {
	projectId := r.FormValue("projectId")
	id := r.FormValue("id")
	data := projectDbDao.Select(projectId, id)
	common.SuccessData(w, data)
}

func ProjectDbInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.ProjectDb{}
	common.Bind(r, &temp)
	row := projectDbDao.Insert(temp.ProjectId, temp)
	if !row {
		common.FailMsg(w, "Save ProjectDb Failed")
		return
	}
	common.SuccessMsg(w, "Save ProjectDb Success")
}

func ProjectDbUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.ProjectDb{}
	common.Bind(r, &temp)
	row := projectDbDao.Update(temp.ProjectId, temp.Id, temp)
	if !row {
		common.FailMsg(w, "Update ProjectDb Failed")
		return
	}
	common.SuccessMsg(w, "Update ProjectDb Success")
}
