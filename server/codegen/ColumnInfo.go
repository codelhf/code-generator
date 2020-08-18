package codegen

import (
	"code-generator-go/server/model"
)

type ColumnInfo struct {
	ColumnName   string //列名
	FieldName    string //字段名
	Type         string //列类型值
	ColumnType   string //列类型
	FieldType    string //字段类型
	IsPrimaryKey bool   //是否主键
	IsKeyword    bool   //是否关键字
	Comment      string //列注释
}

func NewColumnInfo(ColumnName, Type string, IsPrimaryKey, IsKeyword bool, Comment string, typeList []model.TypeColumn) ColumnInfo {
	return ColumnInfo{
		ColumnName:   ColumnName,
		FieldName:    ColumnName2PropertyName(ColumnName),
		Type:         Type,
		ColumnType:   SqlTypeToColumnType(Type, typeList),
		FieldType:    SqlTypeToFieldType(Type, typeList),
		IsPrimaryKey: IsPrimaryKey,
		IsKeyword:    IsKeyword,
		Comment:      Comment,
	}
}
