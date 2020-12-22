package codegen

import (
	"code-generator-go/server/model"
	"code-generator-go/server/util"
	"time"
)

type CodeGenerator struct {
	config model.Configuration
}

var typeColumnDao model.TypeColumnDao
var projectTableDao model.ProjectTableDao
var templateFieldDao model.TemplateFieldDao

func NewGenerator(projectId string) CodeGenerator {
	codeGenerator := CodeGenerator{}
	config := model.NewConfig(projectId)
	config.Project.DateFormat = util.DateToStrFormat(time.Now(), config.Project.DateFormat)
	codeGenerator.config = config
	return codeGenerator
}

func (c *CodeGenerator) Generate(tables []model.ProjectTable) string {
	DbUtil := DbUtil{}
	DbUtil.TypeList = typeColumnDao.ListAll(c.config.DataSource.Type, c.config.DataSource.LanguageType)
	customFields := templateFieldDao.CustomField()
	resultString := ""
	// 目前支持单表生成
	if DbUtil.InitConnection(c.config.DataSource) {
		for i := 0; i < len(tables); i++ {
			table := tables[i]
			projectTable := projectTableDao.Select(table.ProjectId, table.TableName)
			if util.IsBlank(projectTable.Id) {
				projectTableDao.Insert(table)
			} else {
				projectTable.DomainName = table.DomainName
				projectTable.DomainDesc = table.DomainDesc
				projectTableDao.Update(projectTable.Id, projectTable)
			}
			tableName := table.TableName
			domainName := table.DomainName
			domainDesc := table.DomainDesc
			isView := table.TableType == 2
			resultString += c.single(tableName, domainName, domainDesc, isView, customFields, c.config, DbUtil)
		}
		DbUtil.Close()
	}

	return resultString
}

func (c *CodeGenerator) single(tableName, domainName, domainDesc string, isView bool,
	customFields []model.TemplateField, config model.Configuration, dbUtil DbUtil) string {
	invoker := SingleInvoker{}
	invoker.TableName = tableName
	invoker.ClassName = domainName
	invoker.ClassDesc = domainDesc
	invoker.IsView = isView
	invoker.CustomField = customFields
	invoker.config = config
	invoker.dbUtil = dbUtil
	return invoker.execute()
}
