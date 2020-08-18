package codegen

import (
	"code-generator-go/server/model"
	"code-generator-go/server/util"
	"strings"
)

type SingleInvoker struct {
	TableName    string
	ClassName    string
	GeneratedKey string
	IsView       bool
	ColumnInfo   []ColumnInfo
	CustomField  []model.TemplateField
	config       model.Configuration
	dbUtil       DbUtil
}

func (s *SingleInvoker) execute() string {
	if !s.checkBeforeExecute() {
		return ""
	}
	s.InitTableInfo()
	//模板数据
	data2 := make(map[string]interface{})
	//模板之间的依赖需要提前获取所有的template的配置以template名字作为key
	//放在上面防止下面key被覆盖
	projectTemplateList := s.config.ProjectTemplateList
	for i := 0; i < len(projectTemplateList); i++ {
		template := projectTemplateList[i]
		data2[template.Name] = template
	}
	//自定义模板字段
	//放在上面防止下面key被覆盖
	customFieldList := s.CustomField
	for i := 0; i < len(customFieldList); i++ {
		customField := customFieldList[i]
		data2[customField.Name] = customField.Value
	}
	//代码注释配置
	global := s.config.Project
	data2["author"] = global.Author
	data2["company"] = global.Organization
	data2["createTime"] = global.DateFormat
	//代码生成策略
	data2["httpPrefix"] = global.HttpPrefix
	data2["generateRemark"] = global.GenerateRemark == 1
	//代码文件编码
	fileEncode := global.FileEncode

	//实体表信息
	data2["isView"] = s.IsView
	data2["tableName"] = s.TableName
	data2["allColumn"] = s.ColumnInfo
	data2["pKeyColumn"] = s.GetPrimaryKeyColumnInfo()
	data2["ClassName"] = s.ClassName
	data2["className"] = util.FirstToLowerCase(s.ClassName)

	resultString := ""
	// 循环模板生成代码
	for i := 0; i < len(projectTemplateList); i++ {
		template := projectTemplateList[i]
		if template.IsGenerate != 1 {
			continue
		}
		//设置当前模板配置, 通过Template获取包名等
		data2["Template"] = template
		var fileName string
		if template.Type == 2 {
			//后缀加文件格式
			fileName = template.FileSuffix + "." + template.FileType
		} else {
			//表名加后缀加文件格式
			fileName = s.ClassName + template.FileSuffix + "." + template.FileType
		}
		//文件名包含 / 最后一个 / 后面为生成文件名
		if strings.Contains(fileName, "/") {
			index := strings.LastIndex(fileName, "/")
			template.PackageName = template.PackageName + "/" + util.FirstToLowerCase(util.Substring(fileName, 0, index))
			fileName = util.Substring(fileName, index, len(fileName))
		}
		//生成代码字符串
		bytes, err := GenerateToCode(template.Name, template.Template, data2)
		if err != nil {
			resultString += err.Error()
			return resultString
		}
		//将代码字符串写入文件并返回结果
		resultString += CreateFile(template.Directory, template.PackageName, fileName, fileEncode, string(bytes), template.IsOverride == 1)
		resultString += "\n"
	}
	return resultString
}

func (s *SingleInvoker) checkBeforeExecute() bool {
	if &s.config == nil {
		return false
	}
	if &s.dbUtil == nil {
		return false
	}
	if util.IsBlank(s.TableName) {
		return false
	}
	if util.IsBlank(s.ClassName) {
		s.ClassName = TableName2ClassName(s.TableName)
	}
	return true
}

func (s *SingleInvoker) InitTableInfo() {
	fullColumn := s.dbUtil.GetTableInfo(s.TableName)
	for i := 0; i < len(fullColumn); i++ {
		columnInfo := fullColumn[i]
		// 关键字处理
		if &s.config.DataSource.DelimitKeyword != nil && columnInfo.IsKeyword {
			columnInfo.ColumnName = s.config.DataSource.DelimitKeyword + columnInfo.ColumnName + s.config.DataSource.DelimitKeyword
		}
		// 自定义主键
		if &s.GeneratedKey != nil && strings.EqualFold(s.GeneratedKey, columnInfo.ColumnName) {
			columnInfo.IsPrimaryKey = true
		}
		fullColumn[i] = columnInfo
	}
	s.ColumnInfo = fullColumn
}

func (s *SingleInvoker) GetPrimaryKeyColumnInfo() ColumnInfo {
	for i := 0; i < len(s.ColumnInfo); i++ {
		columnInfo := s.ColumnInfo[i]
		if columnInfo.IsPrimaryKey {
			return columnInfo
		}
	}
	//没有主键也没有自定义主键
	return NewColumnInfo("id", "varchar", true, false, "", s.dbUtil.TypeList)
}
