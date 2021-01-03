package controller

import (
	"code-generator-go/server/common"
	"code-generator-go/server/db"
	"code-generator-go/server/model"
	"code-generator-go/server/util"
	"net/http"
	"strings"
)

var templateFieldDao model.TemplateFieldDao

func TemplateFieldList(w http.ResponseWriter, r *http.Request) {
	pageInfo := db.BuildPageInfo(r)
	name := r.FormValue("name")
	desc := r.FormValue("desc")
	Type := util.ParseInt(r.FormValue("type"))
	dataList, total := templateFieldDao.List(pageInfo.PageNum, pageInfo.PageSize, name, desc, Type)
	pageInfo.List = dataList
	pageInfo.Total = total
	common.SuccessData(w, pageInfo)
}

func TemplateFieldSelect(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data := templateFieldDao.Select(id)
	common.SuccessData(w, data)
}

func TemplateFieldCheck(w http.ResponseWriter, r *http.Request) {
	temp := model.TemplateField{}
	common.Bind(r, &temp)
	if !templateFieldCheck(w, temp) {
		return
	}
	common.Success(w)
}

func templateFieldCheck(w http.ResponseWriter, temp model.TemplateField) bool {
	has := templateFieldDao.Check(temp)
	if has {
		common.FailMsg(w, "TemplateField Exists")
		return false
	}
	return true
}

func TemplateFieldInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.TemplateField{}
	common.Bind(r, &temp)
	if !templateFieldCheck(w, temp) {
		return
	}
	row := templateFieldDao.Insert(temp)
	if !row {
		common.FailMsg(w, "Save TemplateField Failed")
		return
	}
	common.SuccessMsg(w, "Save TemplateField Success")
}

func TemplateFieldUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.TemplateField{}
	common.Bind(r, &temp)
	if temp.Type == 1 {
		common.FailMsg(w, "Can not Update Default TemplateField")
		return
	}
	if !templateFieldCheck(w, temp) {
		return
	}
	row := templateFieldDao.Update(temp.Id, temp)
	if !row {
		common.FailMsg(w, "Update TemplateField Failed")
		return
	}
	common.SuccessMsg(w, "Update TemplateField Success")
}

func TemplateFieldDelete(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	idList := strings.Split(ids, ",")
	row := templateFieldDao.Delete(idList)
	if !row {
		common.FailMsg(w, "Delete TemplateField Failed")
		return
	}
	common.SuccessMsg(w, "Delete TemplateField Success")
}
