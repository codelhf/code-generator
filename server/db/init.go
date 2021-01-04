package db

import (
	"code-generator-go/server/util"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var Engine *xorm.Engine

func init() {
	dir, err := os.Getwd()
	Engine, err = xorm.NewEngine("sqlite3", dir+"/db/data.db")
	util.CheckError(err)
	//是否显示SQL语句
	Engine.ShowSQL(true)
	//数据库打开的最大连接数
	Engine.SetMaxOpenConns(2)
	//自动Model
	//res, err := Engine.ImportFile(dir+"/db/update.sql")
	//util.CheckError(err)
	//fmt.Println(res)
}

func NextId() string {
	worker, _ := util.CreateWorker(0, 0)
	return worker.NextIdStr()
}
