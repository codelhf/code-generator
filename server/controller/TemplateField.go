package controller

import (
	"code-generator-go/server/common"
	"code-generator-go/server/db"
	"code-generator-go/server/model"
	"net/http"
	"strings"
)

var templateFieldDao model.TemplateFieldDao

func TemplateFieldList(w http.ResponseWriter, r *http.Request) {
	pageInfo := db.BuildPageInfo(r)
	name := r.FormValue("name")
	desc := r.FormValue("desc")
	dataList, total := templateFieldDao.List(pageInfo.PageNum, pageInfo.PageSize, name, desc)
	pageInfo.List = dataList
	pageInfo.Total = total
	common.SuccessData(w, pageInfo)
}

func TemplateFieldSelect(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data := templateFieldDao.Select(id)
	common.SuccessData(w, data)
}

func TemplateFieldInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.TemplateField{}
	common.Bind(r, &temp)
	row := templateFieldDao.Insert(temp)
	if !row {
		common.FailMsg(w, "save templateField failed")
		return
	}
	common.SuccessMsg(w, "save templateField success")
}

func TemplateFieldUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.TemplateField{}
	common.Bind(r, &temp)
	if temp.Type == 1 {
		common.FailMsg(w,"can not update default field")
		return
	}
	row := templateFieldDao.Update(temp.Id, temp)
	if !row {
		common.FailMsg(w, "update templateField failed")
		return
	}
	common.SuccessMsg(w, "update templateField success")
}

func TemplateFieldDelete(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	idList := strings.Split(ids, ",")
	row := templateFieldDao.Delete(idList)
	if !row {
		common.FailMsg(w, "delete templateField failed")
		return
	}
	common.SuccessMsg(w, "delete templateField success")
}
