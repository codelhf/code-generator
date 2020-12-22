package controller

import (
	"code-generator-go/server/common"
	"code-generator-go/server/db"
	"code-generator-go/server/model"
	"code-generator-go/server/util"
	"net/http"
	"strings"
)

var typeColumnDao model.TypeColumnDao

func TypeColumnList(w http.ResponseWriter, r *http.Request) {
	pageInfo := db.BuildPageInfo(r)
	dbType := util.ParseInt(r.FormValue("dbType"))
	columnType := r.FormValue("columnType")
	languageType := util.ParseInt(r.FormValue("languageType"))
	fieldType := r.FormValue("fieldType")
	dataList, total := typeColumnDao.List(pageInfo.PageNum, pageInfo.PageSize, dbType, languageType, columnType, fieldType)
	pageInfo.List = dataList
	pageInfo.Total = total
	common.SuccessData(w, pageInfo)
}

func TypeColumnSelect(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data := typeColumnDao.Select(id)
	common.SuccessData(w, data)
}

func TypeColumnInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.TypeColumn{}
	common.Bind(r, &temp)
	row := typeColumnDao.Insert(temp)
	if !row {
		common.FailMsg(w, "Save TypeColumn Failed")
		return
	}
	common.SuccessMsg(w, "Save TypeColumn Success")
}

func TypeColumnUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.TypeColumn{}
	common.Bind(r, &temp)
	row := typeColumnDao.Update(temp.Id, temp)
	if !row {
		common.FailMsg(w, "Update TypeColumn Failed")
		return
	}
	common.SuccessMsg(w, "Update TypeColumn Success")
}

func TypeColumnDelete(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	idList := strings.Split(ids, ",")
	row := typeColumnDao.Delete(idList)
	if !row {
		common.FailMsg(w, "Delete TypeColumn Failed")
		return
	}
	common.SuccessMsg(w, "Delete TypeColumn Success")
}
