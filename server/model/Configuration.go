package model

import "code-generator-go/server/db"

type Configuration struct {
	Project             Project
	DataSource          ProjectDb
	ProjectTemplateList []ProjectTemplateDTO
}

func NewConfig(projectId string) Configuration {
	config := Configuration{}
	session := db.Engine.Prepare()
	var project Project
	session.Table("t_project").Where("id = ?", projectId).Get(&project)
	config.Project = project
	//查询数据库
	var projectDb ProjectDb
	session.Table("t_project_db").Where("project_id = ?", projectId).Get(&projectDb)
	config.DataSource = projectDb
	//模板配置
	var projectTemplateList []ProjectTemplateDTO
	sqlStr := "select pt.id, pt.project_id, pt.template_id, pt.directory, pt.package_name, pt.file_suffix, "
	sqlStr += "pt.is_generate, pt.is_override, t.name, t.type, t.file_type, t.template "
	sqlStr += "from t_project_template pt left join t_template t "
	sqlStr += "where (pt.project_id = ? and pt.template_id = t.id)"
	session.SQL(sqlStr, projectId).Find(&projectTemplateList)
	config.ProjectTemplateList = projectTemplateList

	return config
}
