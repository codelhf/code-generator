package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
	"time"
)

type TypeColumnDao struct {
}

func (d *TypeColumnDao) ListAll(dbType, languageType int) []TypeColumn {
	session := db.Engine.Table("t_type_column")
	if dbType > 0 {
		session.Where("db_type = ?", dbType)
	}
	if languageType > 0 {
		session.Where("language_type = ?", languageType)
	}
	dataList := make([]TypeColumn, 0)
	err := session.Find(&dataList)
	util.CheckError(err)
	return dataList
}

func (d *TypeColumnDao) List(pageNum, pageSize int, dbType, languageType int, columnType, fieldType string) ([]TypeColumn, int64) {
	session := db.Engine.Table("t_type_column")
	if pageNum > 0 && pageSize > 0 {
		session.Limit(pageSize, db.GetOffset(pageNum, pageSize))
	}
	if dbType > 0 {
		session.Where("db_type = ?", dbType)
	}
	if util.IsNotBlank(columnType) {
		session.Where("column_type like ?", "%" + columnType + "%")
	}
	if languageType > 0 {
		session.Where("language_type = ?", languageType)
	}
	if util.IsNotBlank(fieldType) {
		session.Where("field_type like ?", "%" + fieldType + "%")
	}
	dataList := make([]TypeColumn, 0)
	total, err := session.OrderBy("column_type, language_type asc").FindAndCount(&dataList)
	util.CheckError(err)
	return dataList, total
}

func (d *TypeColumnDao) Select(id string) TypeColumn {
	session := db.Engine.Table("t_type_column")
	var data TypeColumn
	has, err := session.Where("id = ?", id).Get(&data)
	util.CheckError(err)
	if !has {
		return TypeColumn{}
	}
	return data
}

func (d *TypeColumnDao) Insert(typeColumn TypeColumn) bool {
	session := db.Engine.Table("t_type_column")
	typeColumn.Id = db.UUID()
	typeColumn.CteTime = util.DateToStr(time.Now())
	typeColumn.UteTime = util.DateToStr(time.Now())
	affected, err := session.Insert(&typeColumn)
	util.CheckError(err)
	return affected == 1
}

func (d *TypeColumnDao) Update(id string, typeColumn TypeColumn) bool {
	session := db.Engine.Table("t_type_column")
	typeColumn.UteTime = util.DateToStr(time.Now())
	affected, err := session.Where("id = ?", id).Update(&typeColumn)
	util.CheckError(err)
	return affected == 1
}

func (d *TypeColumnDao) Delete(ids []string) bool {
	session := db.Engine.Table("t_type_column")
	var typeColumn TypeColumn
	affected, err := session.In("id", ids).Delete(&typeColumn)
	util.CheckError(err)
	return affected >= 1
}
