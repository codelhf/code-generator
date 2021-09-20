package model

import (
	"code-generator-go/server/db"
	"code-generator-go/server/util"
)

const (
	LAST_PROJECT      = "last_project"
	CODE_DIRECTORY    = "code_directory"
	CODE_PACKAGE_NAME = "code_package_name"
)

type UserLogDao struct {
}

func (d *UserLogDao) Select(name string) string {
	session := db.Engine.Table("t_user_log")
	var data UserLog
	has, err := session.Where("name = ?", name).Get(&data)
	util.CheckError(err)
	if !has {
		return util.EMPTY
	}
	return data.Value
}

func (d *UserLogDao) Update(name string, value string) bool {
	session := db.Engine.Table("t_user_log")
	data := UserLog{Value: value}
	affected, err := session.Where("name = ?", name).Update(&data)
	util.CheckError(err)
	return affected == 1
}
