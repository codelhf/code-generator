package model

type ProjectTemplate struct {
	Id          string `json:"id"`
	ProjectId   string `json:"projectId"`
	GroupId     string `json:"groupId"`
	TemplateId  string `json:"templateId"`
	Directory   string `json:"directory"`
	PackageName string `json:"packageName"`
	FileSuffix  string `json:"fileSuffix"`
	IsGenerate  int    `json:"isGenerate"`
	IsOverride  int    `json:"isOverride"`
}

type ProjectTemplateDTO struct {
	Id          string `json:"id"`
	ProjectId   string `json:"projectId"`
	GroupId     string `json:"groupId"`
	TemplateId  string `json:"templateId"`
	Directory   string `json:"directory"`
	PackageName string `json:"packageName"`
	FileSuffix  string `json:"fileSuffix"`
	IsGenerate  int    `json:"isGenerate"`
	IsOverride  int    `json:"isOverride"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	FileType    string `json:"fileType"`
	Template    string `json:"template"`
}
