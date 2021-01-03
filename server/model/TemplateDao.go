package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
)

type TemplateDao struct {
}

//对所有项目起作用
func (d *TemplateDao) List(pageNum, pageSize int, groupId, name, desc string) ([]Template, int64) {
	session := db.Engine.Table("t_template")
	if pageNum > 0 && pageSize > 0 {
		session.Limit(pageSize, db.GetOffset(pageNum, pageSize))
	}
	if util.IsNotBlank(groupId) {
		session.Where("group_id = ?", groupId)
	}
	if util.IsNotBlank(name) {
		session.Where("name like ?", "%"+name+"%")
	}
	if util.IsNotBlank(desc) {
		session.Where("desc like ?", "%"+desc+"%")
	}
	dataList := make([]Template, 0)
	total, err := session.OrderBy("name asc").FindAndCount(&dataList)
	util.CheckError(err)
	return dataList, total
}

func (d *TemplateDao) Export(ids []string) []Template {
	session := db.Engine.Table("t_template")
	dataList := make([]Template, 0)
	err := session.In("id", ids).Find(&dataList)
	util.CheckError(err)
	return dataList
}

func (d *TemplateDao) Import(templates []Template) bool {
	session := db.Engine.Table("t_template")
	for index, template := range templates {
		template.Id = db.UUID()
		templates[index] = template
	}
	affected, err := session.InsertMulti(&templates)
	util.CheckError(err)
	return affected >= 1
}

func (d *TemplateDao) Select(id string) Template {
	session := db.Engine.Table("t_template")
	var data Template
	has, err := session.Where("id = ?", id).Get(&data)
	util.CheckError(err)
	if !has {
		return Template{}
	}
	return data
}

func (d *TemplateDao) Check(template Template) bool {
	session := db.Engine.Table("t_template")
	if util.IsNotBlank(template.Id) {
		session.Where("id != ?", template.Id)
	}
	session.Where("group_id = ?", template.GroupId)
	session.Where("name = ?", template.Name)
	var data Template
	has, err := session.Get(&data)
	util.CheckError(err)
	return has
}

func (d *TemplateDao) Insert(template Template) bool {
	session := db.Engine.Table("t_template")
	template.Id = db.UUID()
	affected, err := session.Insert(&template)
	util.CheckError(err)
	return affected == 1
}

func (d *TemplateDao) Update(id string, template Template) bool {
	session := db.Engine.Table("t_template")
	affected, err := session.Where("id = ?", id).Update(&template)
	util.CheckError(err)
	return affected == 1
}

func (d *TemplateDao) Delete(ids []string) bool {
	session := db.Engine.Table("t_template")
	var template Template
	affected, err := session.In("id", ids).Delete(&template)
	util.CheckError(err)
	return affected >= 1
}
