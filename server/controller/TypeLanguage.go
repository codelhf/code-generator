package controller

import (
	"code-generator-go/server/common"
	"code-generator-go/server/db"
	"code-generator-go/server/model"
	"net/http"
	"strings"
)

var typeLanguageDao model.TypeLanguageDao

func TypeLanguageList(w http.ResponseWriter, r *http.Request) {
	pageInfo := db.BuildPageInfo(r)
	dataList, total := typeLanguageDao.List()
	pageInfo.List = dataList
	pageInfo.Total = total
	common.SuccessData(w, dataList)
}

func TypeLanguageSelect(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data := typeLanguageDao.Select(id)
	common.SuccessData(w, data)
}

func TypeLanguageCheck(w http.ResponseWriter, r *http.Request) {
	temp := model.TypeLanguage{}
	common.Bind(r, &temp)
	if !typeLanguageCheck(w, temp) {
		return
	}
	common.Success(w)
}

func typeLanguageCheck(w http.ResponseWriter, temp model.TypeLanguage) bool {
	has := typeLanguageDao.Check(temp)
	if has {
		common.FailMsg(w, "TypeLanguage Exists")
		return false
	}
	return true
}

func TypeLanguageInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.TypeLanguage{}
	common.Bind(r, &temp)
	if !typeLanguageCheck(w, temp) {
		return
	}
	row := typeLanguageDao.Insert(temp)
	if !row {
		common.FailMsg(w, "Save TypeLanguage Failed")
		return
	}
	common.SuccessMsg(w, "Save TypeLanguage Success")
}

func TypeLanguageUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.TypeLanguage{}
	common.Bind(r, &temp)
	if !typeLanguageCheck(w, temp) {
		return
	}
	row := typeLanguageDao.Update(temp.Id, temp)
	if !row {
		common.FailMsg(w, "Update TypeLanguage Failed")
		return
	}
	common.SuccessMsg(w, "Update TypeLanguage Success")
}

func TypeLanguageDelete(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")
	idList := strings.Split(ids, ",")
	row := typeLanguageDao.Delete(idList)
	if !row {
		common.FailMsg(w, "Delete TypeLanguage Failed")
		return
	}
	common.SuccessMsg(w, "Delete TypeLanguage Success")
}
