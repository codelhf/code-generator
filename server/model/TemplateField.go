package model

type TemplateField struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Type  int    `json:"type"`
	Value string `json:"value"`
}
