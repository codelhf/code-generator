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

func AllDBType(w http.ResponseWriter, r *http.Request) {
	dataList := typeColumnDao.AllDBType()
	common.SuccessData(w, dataList)
}

func TypeColumnList(w http.ResponseWriter, r *http.Request) {
	pageInfo := db.BuildPageInfo(r)
	languageId := r.FormValue("languageId")
	dbType := util.ParseInt(r.FormValue("dbId"))
	columnType := r.FormValue("columnType")
	fieldType := r.FormValue("fieldType")
	dataList, total := typeColumnDao.List(pageInfo.PageNum, pageInfo.PageSize, languageId, columnType, fieldType, dbType)
	dbTypeList := typeColumnDao.AllDBType()
	languageList, _ := typeLanguageDao.List()
	dataDTOList := make([]model.TypeColumnDTO, 0)
	for _, typeColumn := range dataList {
		dataDTO := model.TypeColumnDTO{}
		dataDTO.Id = typeColumn.Id
		dataDTO.LanguageId = typeColumn.LanguageId
		dataDTO.DbType = typeColumn.DbType
		dataDTO.ColumnType = typeColumn.ColumnType
		dataDTO.FieldType = typeColumn.FieldType
		for _, dbType := range dbTypeList {
			if dataDTO.DbType == dbType.Id {
				dataDTO.DbName = dbType.Name
			}
		}
		for _, language := range languageList {
			if dataDTO.LanguageId == language.Id {
				dataDTO.LanguageName = language.Name
			}
		}
		dataDTOList = append(dataDTOList, dataDTO)
	}
	pageInfo.List = dataDTOList
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
