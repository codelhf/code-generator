package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
)

type ProjectTableDao struct {
}

// 无用接口接口
func (d *ProjectTableDao) List(pageNum, pageSize int, projectId, tableName string) ([]ProjectTable, int64) {
	session := db.Engine.Table("t_project_table")
	if pageNum > 0 && pageSize > 0 {
		session.Limit(pageSize, db.GetOffset(pageNum, pageSize))
	}
	if util.IsNotBlank(projectId) {
		session.Where("project_id = ?", projectId)
	}
	if util.IsNotBlank(tableName) {
		session.Where("table_name = ?", tableName)
	}
	dataList := make([]ProjectTable, 0)
	total, err := session.FindAndCount(&dataList)
	util.CheckError(err)
	return dataList, total
}

// 数据库表为单选操作，查询表配置接口
func (d *ProjectTableDao) Select(projectId, tableName string) ProjectTable {
	session := db.Engine.Table("t_project_table")
	var data ProjectTable
	has, err := session.Where("project_id = ? and table_name = ?", projectId, tableName).Get(&data)
	util.CheckError(err)
	if !has {
		return ProjectTable{}
	}
	return data
}

// 生成代码之前保存表配置
func (d *ProjectTableDao) Insert(table ProjectTable) bool {
	session := db.Engine.Table("t_project_table")
	table.Id = db.NextId()
	affected, err := session.Insert(&table)
	util.CheckError(err)
	return affected == 1
}

// 生成代码之前更新表配置
func (d *ProjectTableDao) Update(id string, table ProjectTable) bool {
	session := db.Engine.Table("t_project_table")
	affected, err := session.Where("id = ?", id).Update(&table)
	util.CheckError(err)
	return affected == 1
}

// 项目删除后同步删除表配置(数据库表有删除、表名修改时不常用会出现多余表配置)
// 数据库表名修改相当于新增表，表配置新增即可
func (d *ProjectTableDao) Delete(projectId string, ids []string) bool {
	session := db.Engine.Table("t_project_table")
	var table ProjectTable
	affected, err := session.Where("project_id = ?", projectId).In("id", ids).Delete(&table)
	util.CheckError(err)
	return affected >= 1
}
