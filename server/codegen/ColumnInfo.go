package codegen

import (
	"code-generator-go/server/model"
)

type ColumnInfo struct {
	ColumnName   string //列名
	ColumnType   string //列类型
	FieldName    string //字段名
	FieldType    string //字段类型
	JdbcType     string //jdbc类型
	IsPrimaryKey bool   //是否主键
	IsKeyword    bool   //是否关键字
	Comment      string //列注释
}

func NewColumnInfo(ColumnName, ColumnType string, IsPrimaryKey, IsKeyword bool, Comment string, typeList []model.TypeColumn) ColumnInfo {
	return ColumnInfo{
		ColumnName:   ColumnName,
		ColumnType:   ColumnType,
		FieldName:    ColumnName2PropertyName(ColumnName),
		FieldType:    ColumnTypeToFieldType(ColumnType, typeList),
		JdbcType:     ColumnTypeToJdbcType(ColumnType, typeList),
		IsPrimaryKey: IsPrimaryKey,
		IsKeyword:    IsKeyword,
		Comment:      Comment,
	}
}
