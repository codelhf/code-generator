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
	LanguageId     string `json:"languageId"`
}
