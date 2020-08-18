package controller

import (
	"code-generator-go/server/codegen"
	"code-generator-go/server/common"
	"code-generator-go/server/model"
	"code-generator-go/server/util"
	"net/http"
	"sort"
)

var projectTableDao model.ProjectTableDao

func ProjectTableAll(w http.ResponseWriter, r *http.Request) {
	projectId := r.FormValue("projectId")
	dataSource := projectDbDao.Select(projectId, "")
	if util.IsBlank(dataSource.Url) {
		common.FailMsg(w, "please add database")
		return
	}
	data := make([]model.ProjectTable, 0)
	dbUtil := codegen.DbUtil{}
	if dbUtil.InitConnection(dataSource) {
		data = dbUtil.GetAllTable()
		dbUtil.Close()
		for i := 0; i < len(data); i++ {
			data[i].ProjectId = projectId
		}
		sort.Slice(data, func(i, j int) bool {
			return data[i].TableType < data[j].TableType
		})
	}
	common.SuccessData(w, data)
}

func ProjectTableSelect(w http.ResponseWriter, r *http.Request) {
	projectId := r.FormValue("projectId")
	tableName := r.FormValue("tableName")
	data := projectTableDao.Select(projectId, tableName)
	common.SuccessData(w, data)
}

func ProjectTableGenerate(w http.ResponseWriter, r *http.Request) {
	projectId := r.FormValue("projectId")
	tableList := make([]model.ProjectTable, 0)
	common.Bind(r, &tableList)
	generator := codegen.NewGenerator(projectId)
	result := generator.Generate(tableList)
	common.SuccessData(w, result)
}
