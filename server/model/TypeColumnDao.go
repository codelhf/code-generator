package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
)

type TypeColumnDao struct {
}

func (d *TypeColumnDao) AllDBType() []TypeDB {
	dataList := make([]TypeDB, 0)
	dataList = append(dataList, TypeDB{MySQLType, "MySQL"})
	dataList = append(dataList, TypeDB{OracleType, "Oracle"})
	dataList = append(dataList, TypeDB{PostgreSQLType, "PostgreSQL"})
	return dataList
}

//codegen使用
func (d *TypeColumnDao) ListAll(dbType int, languageId string) []TypeColumn {
	session := db.Engine.Table("t_type_column")
	if dbType > 0 {
		session.Where("db_type = ?", dbType)
	}
	if util.IsNotBlank(languageId) {
		session.Where("language_id = ?", languageId)
	}
	dataList := make([]TypeColumn, 0)
	err := session.Find(&dataList)
	util.CheckError(err)
	return dataList
}

func (d *TypeColumnDao) List(pageNum, pageSize int, languageId, columnType, fieldType string, dbType int) ([]TypeColumn, int64) {
	session := db.Engine.Table("t_type_column")
	if pageNum > 0 && pageSize > 0 {
		session.Limit(pageSize, db.GetOffset(pageNum, pageSize))
	}
	if util.IsNotBlank(languageId) {
		session.Where("language_id = ?", languageId)
	}
	if dbType > 0 {
		session.Where("db_type = ?", dbType)
	}
	if util.IsNotBlank(columnType) {
		session.Where("column_type like ?", "%"+columnType+"%")
	}
	if util.IsNotBlank(fieldType) {
		session.Where("field_type like ?", "%"+fieldType+"%")
	}
	dataList := make([]TypeColumn, 0)
	total, err := session.OrderBy("column_type, language_id asc").FindAndCount(&dataList)
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
	affected, err := session.Insert(&typeColumn)
	util.CheckError(err)
	return affected == 1
}

func (d *TypeColumnDao) Update(id string, typeColumn TypeColumn) bool {
	session := db.Engine.Table("t_type_column")
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
