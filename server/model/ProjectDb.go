package model

type ProjectDb struct {
	Id             string `json:"id"`
	ProjectId      string `json:"projectId"`
	Type           int    `json:"type"`
	Url            string `json:"url"`
	Database       string `json:"database"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	DelimitKeyword string `json:"delimitKeyword"`
	LanguageType   int    `json:"languageType"`
}

const (
	MysqlType      = 1
	MysqlName      = "mysql"
	OracleType     = 2
	OracleName     = "oci8"
	PostgreSQLType = 3
	PostgreSQLName = "postgres"
)
