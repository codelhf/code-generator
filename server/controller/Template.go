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
	name := r.FormValue("name")
	desc := r.FormValue("desc")
	dataList, total := templateDao.List(pageInfo.PageNum, pageInfo.PageSize, name, desc)
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
		common.FailMsg(w, "import templates failed")
		return
	}
	common.SuccessMsg(w,"import templates success")
}

func TemplateSelect(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data := templateDao.Select(id)
	common.SuccessData(w, data)
}

func TemplateNameExists(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	row := templateDao.NameExists(id, name)
	if row {
		common.FailMsg(w, "template is exists")
		return
	}
	common.SuccessMsg(w, "")
}

func TemplateInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.Template{}
	common.Bind(r, &temp)
	row := templateDao.Insert(temp)
	if !row {
		common.FailMsg(w, "save template failed")
		return
	}
	common.SuccessMsg(w, "save template success")
}

func TemplateUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.Template{}
	common.Bind(r, &temp)
	row := templateDao.Update(temp.Id, temp)
	if !row {
		common.FailMsg(w, "update template failed")
		return
	}
	common.SuccessMsg(w, "update template success")
}

func TemplateDelete(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	idList := strings.Split(ids, ",")
	row := templateDao.Delete(idList)
	if !row {
		common.FailMsg(w, "delete template failed")
		return
	}
	common.SuccessMsg(w, "delete template success")
}
