package model

type TypeColumn struct {
	Id           string `json:"id"`
	LanguageId   string `json:"languageId"`
	DbType       int    `json:"dbType"`
	ColumnType   string `json:"columnType"`
	FieldType    string `json:"fieldType"`
}

type TypeColumnDTO struct {
	Id           string `json:"id"`
	LanguageId   string `json:"languageId"`
	LanguageName string `json:"languageName"`
	DbType       int    `json:"dbType"`
	DbName       string `json:"dbName"`
	ColumnType   string `json:"columnType"`
	FieldType    string `json:"fieldType"`
}

type TypeDB struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

const (
	MySQLType      = 1
	MysqlName      = "mysql"
	OracleType     = 2
	OracleName     = "oci8"
	PostgreSQLType = 3
	PostgreSQLName = "postgres"
)
