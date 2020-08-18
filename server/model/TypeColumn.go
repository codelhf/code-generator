package model

type TypeColumn struct {
	Id           string `json:"id"`
	DbType       int    `json:"dbType"`
	ColumnType   string `json:"columnType"`
	LanguageType int    `json:"languageType"`
	FieldType    string `json:"fieldType"`
	CteTime      string `json:"cteTime"`
	UteTime      string `json:"uteTime"`
}
