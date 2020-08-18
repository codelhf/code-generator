package db

import (
	"code-generator-go/server/util"
	"net/http"
	"strconv"
)

type PageInfo struct {
	PageNum  int         `json:"pageNum" form:"pageNum"`   //页码
	PageSize int         `json:"pageSize" form:"pageSize"` //每页条数
	Total    int64       `json:"total" form:"total"`       //总条数
	List     interface{} `json:"list" form:"total"`        //列表
}

func BuildPageInfo(r *http.Request) PageInfo {
	query := r.URL.Query()
	pageNum := query.Get("pageNum")
	pageSize := query.Get("pageSize")
	var num, size int
	var err error
	if util.IsNotBlank(pageNum) {
		num, err = strconv.Atoi(pageNum)
		util.CheckError(err)
	}
	if util.IsNotBlank(pageSize) {
		size, err = strconv.Atoi(pageSize)
		util.CheckError(err)
	}
	pageInfo := PageInfo{}
	pageInfo.PageNum = num
	pageInfo.PageSize = size
	return pageInfo
}

func GetOffset(pageNum, pageSize int) int {
	return (pageNum - 1) * pageSize
}
