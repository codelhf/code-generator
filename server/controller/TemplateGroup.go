package controller

import (
	"code-generator-go/server/common"
	"code-generator-go/server/db"
	"code-generator-go/server/model"
	"net/http"
	"strings"
)

var templateGroupDao model.TemplateGroupDao

func TemplateGroupList(w http.ResponseWriter, r *http.Request) {
	pageInfo := db.BuildPageInfo(r)
	dataList, total := templateGroupDao.List()
	pageInfo.List = dataList
	pageInfo.Total = total
	common.SuccessData(w, dataList)
}

func TemplateGroupSelect(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data := templateGroupDao.Select(id)
	common.SuccessData(w, data)
}

func TemplateGroupCheck(w http.ResponseWriter, r *http.Request) {
	temp := model.TemplateGroup{}
	common.Bind(r, &temp)
	if !templateGroupCheck(w, temp) {
		return
	}
	common.Success(w)
}

func templateGroupCheck(w http.ResponseWriter, temp model.TemplateGroup) bool {
	has := templateGroupDao.Check(temp)
	if has {
		common.FailMsg(w, "TemplateGroup Exists")
		return false
	}
	return true
}

func TemplateGroupInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.TemplateGroup{}
	common.Bind(r, &temp)
	if !templateGroupCheck(w, temp) {
		return
	}
	row := templateGroupDao.Insert(temp)
	if !row {
		common.FailMsg(w, "Save TemplateGroup Failed")
		return
	}
	common.SuccessMsg(w, "Save TemplateGroup Success")
}

func TemplateGroupUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.TemplateGroup{}
	common.Bind(r, &temp)
	if !templateGroupCheck(w, temp) {
		return
	}
	row := templateGroupDao.Update(temp.Id, temp)
	if !row {
		common.FailMsg(w, "Update TemplateGroup Failed")
		return
	}
	common.SuccessMsg(w, "Update TemplateGroup Success")
}

func TemplateGroupDelete(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	idList := strings.Split(ids, ",")
	row := templateGroupDao.Delete(idList)
	if !row {
		common.FailMsg(w, "Delete TemplateGroup Failed")
		return
	}
	common.SuccessMsg(w, "Delete TemplateGroup Success")
}
