package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
)

type TypeLanguageDao struct {
}

func (d *TypeLanguageDao) List() ([]TypeLanguage, int64) {
	session := db.Engine.Table("t_type_language")
	dataList := make([]TypeLanguage, 0)
	total, err := session.OrderBy("id asc").FindAndCount(&dataList)
	util.CheckError(err)
	return dataList, total
}

func (d *TypeLanguageDao) Select(id string) TypeLanguage {
	session := db.Engine.Table("t_type_language")
	var data TypeLanguage
	has, err := session.Where("id = ?", id).Get(&data)
	util.CheckError(err)
	if !has {
		return TypeLanguage{}
	}
	return data
}

func (d *TypeLanguageDao) Check(typeLanguage TypeLanguage) bool {
	session := db.Engine.Table("t_type_language")
	if util.IsNotBlank(typeLanguage.Id) {
		session.Where("id != ?", typeLanguage.Id)
	}
	session.Where("name = ?", typeLanguage.Name)
	var data TypeLanguage
	has, err := session.Get(&data)
	util.CheckError(err)
	return has
}

func (d *TypeLanguageDao) Insert(typeLanguage TypeLanguage) bool {
	session := db.Engine.Table("t_type_language")
	typeLanguage.Id = db.NextId()
	affected, err := session.Insert(&typeLanguage)
	util.CheckError(err)
	return affected == 1
}

func (d *TypeLanguageDao) Update(id string, typeLanguage TypeLanguage) bool {
	session := db.Engine.Table("t_type_language")
	affected, err := session.Where("id = ?", id).Update(&typeLanguage)
	util.CheckError(err)
	return affected == 1
}

func (d *TypeLanguageDao) Delete(ids []string) bool {
	session := db.Engine.Table("t_type_column")
	var typeColumn TypeColumn
	affected, err := session.In("language_id", ids).Delete(&typeColumn)
	session = db.Engine.Table("t_type_language")
	var typeLanguage TypeLanguage
	affected, err = session.In("id", ids).Delete(&typeLanguage)
	util.CheckError(err)
	return affected >= 1
}
