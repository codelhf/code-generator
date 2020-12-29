package model

type Template struct {
	Id       string `json:"id"`
	GroupId  string `json:"groupId"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Type     int    `json:"type"`
	FileType string `json:"fileType"`
	Template string `json:"template"`
}
