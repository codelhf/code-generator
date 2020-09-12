package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
)

type TemplateFieldDao struct {
}

//对所有项目起作用
func (d *TemplateFieldDao) CustomField() []TemplateField {
	session := db.Engine.Table("t_template_field")
	dataList := make([]TemplateField, 0)
	err := session.Where("type = ?", 2).Find(&dataList)
	util.CheckError(err)
	return dataList
}

func (d *TemplateFieldDao) List(pageNum, pageSize int, name, desc string) ([]TemplateField, int64) {
	session := db.Engine.Table("t_template_field")
	if pageNum > 0 && pageSize > 0 {
		session.Limit(pageSize, db.GetOffset(pageNum, pageSize))
	}
	if util.IsNotBlank(name) {
		session.Where("name like ?", "%"+name+"%")
	}
	if util.IsNotBlank(desc) {
		session.Where("desc like ?", "%"+desc+"%")
	}
	dataList := make([]TemplateField, 0)
	total, err := session.OrderBy("type asc").FindAndCount(&dataList)
	util.CheckError(err)
	return dataList, total
}

func (d *TemplateFieldDao) Select(id string) TemplateField {
	session := db.Engine.Table("t_template_field")
	var data TemplateField
	has, err := session.Where("id = ?", id).Get(&data)
	util.CheckError(err)
	if !has {
		return TemplateField{}
	}
	return data
}

func (d *TemplateFieldDao) Insert(templateField TemplateField) bool {
	session := db.Engine.Table("t_template_field")
	templateField.Id = db.UUID()
	templateField.Type = 2
	affected, err := session.Insert(&templateField)
	util.CheckError(err)
	return affected == 1
}

func (d *TemplateFieldDao) Update(id string, templateField TemplateField) bool {
	session := db.Engine.Table("t_template_field")
	templateField.Type = 2
	affected, err := session.Where("id = ?", id).Update(&templateField)
	util.CheckError(err)
	return affected == 1
}

func (d *TemplateFieldDao) Delete(ids []string) bool {
	session := db.Engine.Table("t_template_field")
	var templateField TemplateField
	affected, err := session.In("id", ids).Delete(&templateField)
	util.CheckError(err)
	return affected >= 1
}
