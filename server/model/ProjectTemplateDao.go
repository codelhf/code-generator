package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
)

type ProjectTemplateDao struct {
}

//在项目配置表单为小列表功能
func (d *ProjectTemplateDao) List(pageNum, pageSize int, projectId string) ([]ProjectTemplateDTO, int64) {
	session := db.Engine.Table("t_project_template")
	if pageNum > 0 && pageSize > 0 {
		session.Limit(pageSize, db.GetOffset(pageNum, pageSize))
	}
	if util.IsNotBlank(projectId) {
		session.Where("project_id = ?", projectId)
	}
	dataList := make([]ProjectTemplate, 0)
	total, err := session.FindAndCount(&dataList)
	util.CheckError(err)
	var ids []string
	for i := 0; i < len(dataList); i++ {
		ids = append(ids, dataList[i].TemplateId)
	}
	// 查询
	session.Table("t_template")
	templateList := make([]Template, 0)
	err = session.In("id", ids).OrderBy("name").Find(&templateList)
	util.CheckError(err)
	dtoList := toDTO(dataList, templateList)
	return dtoList, total
}

func toDTO(dataList []ProjectTemplate, templateList []Template) []ProjectTemplateDTO {
	dtoList := make([]ProjectTemplateDTO, 0)
	for i := 0; i < len(templateList); i++ {
		template := templateList[i]
		for j := 0; j < len(dataList); j++ {
			data := dataList[j]
			if data.TemplateId == template.Id {
				dto := ProjectTemplateDTO{}
				dto.Id = data.Id
				dto.ProjectId = data.ProjectId
				dto.GroupId = data.GroupId
				dto.TemplateId = data.TemplateId
				dto.Directory = data.Directory
				dto.PackageName = data.PackageName
				dto.FileSuffix = data.FileSuffix
				dto.IsGenerate = data.IsGenerate
				dto.IsOverride = data.IsOverride

				dto.Name = template.Name
				dto.Type = template.Type
				dto.FileType = template.FileType

				dtoList = append(dtoList, dto)
			}
		}
	}

	return dtoList
}

func (d *ProjectTemplateDao) Select(id string) ProjectTemplate {
	session := db.Engine.Table("t_project_template")
	var data ProjectTemplate
	has, err := session.Where("id = ?", id).Get(&data)
	util.CheckError(err)
	if !has {
		return ProjectTemplate{}
	}
	return data
}

func (d *ProjectTemplateDao) Check(pTemplate ProjectTemplate) bool {
	session := db.Engine.Table("t_project_template")
	if util.IsNotBlank(pTemplate.Id) {
		session.Where("id != ?", pTemplate.Id)
	}
	session.Where("project_id = ?", pTemplate.ProjectId)
	session.Where("template_id = ?", pTemplate.TemplateId)
	var data ProjectTemplate
	has, err := session.Get(&data)
	util.CheckError(err)
	return has
}

func (d *ProjectTemplateDao) Insert(pTemplate ProjectTemplate) bool {
	session := db.Engine.Table("t_project_template")
	pTemplate.Id = db.NextId()
	affected, err := session.Insert(&pTemplate)
	util.CheckError(err)
	return affected == 1
}

func (d *ProjectTemplateDao) Update(id string, pTemplate ProjectTemplate) bool {
	session := db.Engine.Table("t_project_template")
	affected, err := session.Where("id = ?", id).Update(&pTemplate)
	util.CheckError(err)
	return affected == 1
}

func (d *ProjectTemplateDao) Delete(projectId string, ids []string) bool {
	session := db.Engine.Table("t_project_template")
	if util.IsNotBlank(projectId) {
		session.Where("project_id = ?", projectId)
	}
	var pTemplate ProjectTemplate
	affected, err := session.In("id", ids).Delete(&pTemplate)
	util.CheckError(err)
	return affected >= 1
}
