package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
)

type ProjectDbDao struct {
}

func (d *ProjectDbDao) List(pageNum, pageSize int, typ int, host string) ([]ProjectDb, int64) {
	session := db.Engine.Table("t_project_db")
	if pageNum > 0 && pageSize > 0 {
		session.Limit(pageSize, db.GetOffset(pageNum, pageSize))
	}
	if typ > 0 {
		session.Where("type = ?", typ)
	}
	if util.IsNotBlank(host) {
		session.Where("host like ?", "%" + host + "%")
	}
	dataList := make([]ProjectDb, 0)
	total, err := session.FindAndCount(&dataList)
	util.CheckError(err)
	return dataList, total
}

func (d *ProjectDbDao) Select(projectId string, id string) ProjectDb {
	session := db.Engine.Table("t_project_db")
	if util.IsNotBlank(projectId) {
		session.Where("project_id = ?", projectId)
	}
	if util.IsNotBlank(id) {
		session.Where("id = ?", id)
	}
	var data ProjectDb
	has, err := session.Get(&data)
	util.CheckError(err)
	if !has {
		return ProjectDb{}
	}
	return data
}

func (d *ProjectDbDao) Insert(projectId string, projectDb ProjectDb) bool {
	session := db.Engine.Table("t_project_db")
	projectDb.Id = db.UUID()
	projectDb.ProjectId = projectId
	affected, err := session.Insert(&projectDb)
	util.CheckError(err)
	return affected == 1
}

func (d *ProjectDbDao) Update(projectId string, id string, projectDb ProjectDb) bool {
	session := db.Engine.Table("t_project_db")
	affected, err := session.Where("project_id = ? and id = ?", projectId, id).Update(&projectDb)
	util.CheckError(err)
	return affected == 1
}

func (d *ProjectDbDao) Delete(projectId string, ids []string) bool {
	session := db.Engine.Table("t_project_db")
	if util.IsNotBlank(projectId) {
		session.Where("project_id = ?", projectId)
	}
	if len(ids) > 0 {
		session.In("id", ids)
	}
	var projectDb ProjectDb
	affected, err := session.Delete(&projectDb)
	util.CheckError(err)
	return affected >= 1
}
