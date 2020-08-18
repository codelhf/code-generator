package model

type ProjectTable struct {
	Id          string `json:"id"`
	ProjectId   string `json:"projectId"`
	TableName   string `json:"tableName"`
	TableType   int    `json:"tableType"`
	DomainName  string `json:"domainName"`
	GenerateKey string `json:"generateKey"`
}
