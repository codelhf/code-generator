package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
)

type TemplateGroupDao struct {
}

func (d *TemplateGroupDao) List() ([]TemplateGroup, int64) {
	session := db.Engine.Table("t_template_group")
	dataList := make([]TemplateGroup, 0)
	total, err := session.OrderBy("name asc").FindAndCount(&dataList)
	util.CheckError(err)
	return dataList, total
}

func (d *TemplateGroupDao) Select(id string) TemplateGroup {
	session := db.Engine.Table("t_template_group")
	var data TemplateGroup
	has, err := session.Where("id = ?", id).Get(&data)
	util.CheckError(err)
	if !has {
		return TemplateGroup{}
	}
	return data
}

func (d *TemplateGroupDao) Insert(templateGroup TemplateGroup) bool {
	session := db.Engine.Table("t_template_group")
	templateGroup.Id = db.UUID()
	affected, err := session.Insert(&templateGroup)
	util.CheckError(err)
	return affected == 1
}

func (d *TemplateGroupDao) Update(id string, templateGroup TemplateGroup) bool {
	session := db.Engine.Table("t_template_group")
	affected, err := session.Where("id = ?", id).Update(&templateGroup)
	util.CheckError(err)
	return affected == 1
}

func (d *TemplateGroupDao) Delete(ids []string) bool {
	session := db.Engine.Table("t_template_group")
	var templateGroup TemplateGroup
	affected, err := session.In("id", ids).Delete(&templateGroup)
	util.CheckError(err)
	return affected >= 1
}
