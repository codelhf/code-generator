package controller

import (
	"code-generator-go/server/common"
	"code-generator-go/server/db"
	"code-generator-go/server/model"
	"code-generator-go/server/util"
	"net/http"
	"strings"
)

var templateDao model.TemplateDao

func TemplateList(w http.ResponseWriter, r *http.Request) {
	pageInfo := db.BuildPageInfo(r)
	groupId := r.FormValue("groupId")
	name := r.FormValue("name")
	desc := r.FormValue("desc")
	dataList, total := templateDao.List(pageInfo.PageNum, pageInfo.PageSize, groupId, name, desc)
	pageInfo.List = dataList
	pageInfo.Total = total
	common.SuccessData(w, pageInfo)
}

func TemplateExport(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("fileName")
	sheetName := r.FormValue("sheetName")
	heads := r.FormValue("heads")
	ids := r.FormValue("ids")
	idList := strings.Split(ids, ",")
	dataList := templateDao.Export(idList)
	util.WriteExcel(w, fileName, sheetName, heads, dataList)
}

func TemplateImport(w http.ResponseWriter, r *http.Request) {
	sheetName := r.FormValue("sheetName")
	heads := r.FormValue("heads")
	dataList := make([]model.Template, 0)
	util.ReadExcel(r, "uploadFile", sheetName, heads, &dataList)
	row := templateDao.Import(dataList)
	if !row {
		common.FailMsg(w, "Import Templates Failed")
		return
	}
	common.SuccessMsg(w, "Import Templates Success")
}

func TemplateSelect(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data := templateDao.Select(id)
	common.SuccessData(w, data)
}

func TemplateNameCheck(w http.ResponseWriter, r *http.Request) {
	temp := model.Template{}
	common.Bind(r, &temp)
	if !templateCheck(w, temp) {
		return
	}
	common.Success(w)
}

func templateCheck(w http.ResponseWriter, temp model.Template) bool {
	row := templateDao.Check(temp)
	if row {
		common.FailMsg(w, "Template is Exists")
		return false
	}
	return true
}

func TemplateInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.Template{}
	common.Bind(r, &temp)
	if !templateCheck(w, temp) {
		return
	}
	row := templateDao.Insert(temp)
	if !row {
		common.FailMsg(w, "Save Template Failed")
		return
	}
	common.SuccessMsg(w, "Save Template Success")
}

func TemplateUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.Template{}
	common.Bind(r, &temp)
	if !templateCheck(w, temp) {
		return
	}
	row := templateDao.Update(temp.Id, temp)
	if !row {
		common.FailMsg(w, "Update Template Failed")
		return
	}
	common.SuccessMsg(w, "Update Template Success")
}

func TemplateDelete(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	idList := strings.Split(ids, ",")
	row := templateDao.Delete(idList)
	if !row {
		common.FailMsg(w, "Delete Template Failed")
		return
	}
	common.SuccessMsg(w, "Delete Template Success")
}
