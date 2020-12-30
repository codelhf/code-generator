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
	has := typeLanguageDao.Check(temp.Name)
	if has {
		common.SuccessMsg(w, "TypeLanguage Exists")
		return
	}
	common.Success(w)
}

func TypeLanguageInsert(w http.ResponseWriter, r *http.Request) {
	temp := model.TypeLanguage{}
	common.Bind(r, &temp)
	row := typeLanguageDao.Insert(temp)
	if row == -1 {
		common.FailMsg(w, "TypeLanguage Exists")
		return
	}
	if row == 0 {
		common.FailMsg(w, "Save TypeLanguage Failed")
		return
	}
	common.SuccessMsg(w, "Save TypeLanguage Success")
}

func TypeLanguageUpdate(w http.ResponseWriter, r *http.Request) {
	temp := model.TypeLanguage{}
	common.Bind(r, &temp)
	row := typeLanguageDao.Update(temp.Id, temp)
	if row == -1 {
		common.FailMsg(w, "TypeLanguage Exists")
		return
	}
	if row == 0 {
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
