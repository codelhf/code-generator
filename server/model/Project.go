package model

type Project struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Author         string `json:"author"`
	Organization   string `json:"organization"`
	DateFormat     string `json:"dateFormat"`
	FileEncode     string `json:"fileEncode"`
	HttpPrefix     string `json:"httpPrefix"`
	GenerateRemark int    `json:"generateRemark"`
	CteTime        string `json:"cteTime"`
	UteTime        string `json:"uteTime"`
}
