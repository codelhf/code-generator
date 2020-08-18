package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
	"time"
)

type ProjectDao struct {
}

func (d *ProjectDao) List(pageNum, pageSize int, name, desc string) ([]Project, int64) {
	session := db.Engine.Table("t_project")
	if pageNum > 0 && pageSize > 0 {
		session.Limit(pageSize, db.GetOffset(pageNum, pageSize))
	}
	if util.IsNotBlank(name) {
		session.Where("name like ?", "%" + name + "%")
	}
	if util.IsNotBlank(desc) {
		session.Where("desc like ?", "%" + desc + "%")
	}
	dataList := make([]Project, 0)
	total, err := session.OrderBy("ute_time desc").FindAndCount(&dataList)
	util.CheckError(err)
	return dataList, total
}

func (d *ProjectDao) Select(id string) Project {
	session := db.Engine.Table("t_project")
	var data Project
	has, err := session.Where("id = ?", id).Get(&data)
	util.CheckError(err)
	if !has {
		return Project{}
	}
	return data
}

func (d *ProjectDao) Insert(project Project) bool {
	session := db.Engine.Table("t_project")
	project.Id = db.UUID()
	project.CteTime = util.DateToStr(time.Now())
	project.UteTime = util.DateToStr(time.Now())
	affected, err := session.Insert(&project)
	util.CheckError(err)
	return affected == 1
}

func (d *ProjectDao) Update(id string, project Project) bool {
	session := db.Engine.Table("t_project")
	project.UteTime = util.DateToStr(time.Now())
	affected, err := session.Where("id = ?", id).Update(&project)
	util.CheckError(err)
	return affected == 1
}

//删除项目
//1.删除数据库配置
//2.删除数据库表配置
//3.删除代码模板配置
//4.删除项目
func (d *ProjectDao) Delete(ids []string) bool {
	session := db.Engine.Table("t_project")
	var project Project
	affected, err := session.In("id", ids).Delete(&project)
	util.CheckError(err)
	return affected >= 1
}

//获取最后更新的项目
func (d *ProjectDao) LastId() string {
	session := db.Engine.Table("t_project")
	var data Project
	has, err := session.Desc("ute_time").Limit(1).Get(&data)
	util.CheckError(err)
	if !has {
		return ""
	}
	return data.Id
}
