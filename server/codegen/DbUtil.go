package codegen

import (
	"code-generator-go/server/model"
	"code-generator-go/server/util"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	//_ "github.com/mattn/go-oci8"
	"strings"
)

type DbUtil struct {
	Db       *sql.DB
	DbType   int
	DbName   string
	TypeList []model.TypeColumn
}

//初始化数据库连接
func (d *DbUtil) InitConnection(dataSource model.ProjectDb) bool {
	var err error
	d.DbType = dataSource.Type
	d.DbName = dataSource.Database
	driverName := GetDriverName(dataSource.Type)
	dataSourceName := getDataSourceName(dataSource)
	d.Db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return false
	}
	return true
}

//关闭数据库连接
func (d *DbUtil) Close() bool {
	_ = d.Db.Close()
	return true
}

//获取所有的表
func (d *DbUtil) GetAllTable() []model.ProjectTable {
	var sqlStr string
	var rows *sql.Rows
	if d.DbType == model.MySQLType {
		sqlStr = "SELECT table_name, table_type FROM information_schema.`TABLES` WHERE table_schema = ? ORDER BY table_type, table_name"
		prepare, err := d.Db.Prepare(sqlStr)
		checkErr(err)
		rows, err = prepare.Query(d.DbName)
		checkErr(err)
	} else if d.DbType == model.OracleType {
		sqlStr = "SELECT table_name, table_type FROM user_tab_comments WHERE table_name NOT LIKE '%$0' ORDER BY table_type, table_name"
		prepare, err := d.Db.Prepare(sqlStr)
		checkErr(err)
		rows, err = prepare.Query()
		checkErr(err)
	} else if d.DbType == model.PostgreSQLType {
		sqlStr = "SELECT * FROM (SELECT tablename AS table_name, 'TABLE' AS table_type FROM pg_tables WHERE schemaname='public' UNION SELECT viewname AS table_name, 'VIEW' AS table_type FROM pg_views WHERE schemaname='public') a ORDER BY table_type, table_name"
		prepare, err := d.Db.Prepare(sqlStr)
		checkErr(err)
		rows, err = prepare.Query()
		checkErr(err)
	}
	defer rows.Close()
	var resultList []model.ProjectTable
	for rows.Next() {
		table := model.ProjectTable{}
		tableType := ""
		err := rows.Scan(&table.TableName, &tableType)
		if "VIEW" == tableType {
			table.TableType = 2
		} else {
			table.TableType = 1
		}
		checkErr(err)
		resultList = append(resultList, table)
	}
	return resultList
}

//获取表结构数据
func (d *DbUtil) GetTableInfo(tableName string) (columnList []ColumnInfo) {
	var sqlStr string
	var rows *sql.Rows
	if d.DbType == model.MySQLType {
		sqlStr = "SELECT c.column_name, c.data_type, c.column_key, c.column_comment FROM information_schema.COLUMNS c WHERE c.table_schema = ? AND c.table_name = ?"
		prepare, err := d.Db.Prepare(sqlStr)
		checkErr(err)
		rows, err = prepare.Query(d.DbName, tableName)
		checkErr(err)
	} else if d.DbType == model.OracleType {
		sqlStr = `SELECT c.column_name, c.data_type, c.column_key, c.comments
FROM 
(
	SELECT * FROM	(
		SELECT t.column_name, t.data_type, c.comments
		FROM user_tab_columns t, user_col_comments c
		WHERE t.column_name = c.column_name AND t.table_name = :name AND c.table_name = :name
	) a left join (
		SELECT a.constraint_name,  a.column_name as column_key
		FROM user_cons_columns a, user_constraints b 
		WHERE a.constraint_name = b.constraint_name AND b.constraint_type = 'P' AND a.table_name = :name
	) b ON (a.column_name = b.column_key)
) c `
		prepare, err := d.Db.Prepare(sqlStr)
		checkErr(err)
		rows, err = prepare.Query(tableName, tableName, tableName)
		checkErr(err)
	} else if d.DbType == model.PostgreSQLType {
		sqlStr = `SELECT a.attname AS column_name, t.typname AS data_type, p.conname AS column_key, d.description AS column_comment
FROM 
	pg_attribute a
	LEFT JOIN pg_class c ON a.attrelid = c.oid
	LEFT JOIN pg_type t ON a.atttypid = t.oid
	LEFT JOIN pg_description d ON a.attrelid = d.objoid AND a.attnum = d.objsubid
	LEFT JOIN pg_constraint p ON a.attnum = p.conkey[1] AND p.contype = 'p' AND p.conrelid = c.oid
WHERE c.relname = $1 AND a.attnum > 0`
		prepare, err := d.Db.Prepare(sqlStr)
		checkErr(err)
		rows, err = prepare.Query(tableName)
		checkErr(err)
	}
	defer rows.Close()
	var resultList []ColumnInfo
	for rows.Next() {
		ColumnName := ""
		Type := ""
		PrimaryKey := sql.NullString{String: "", Valid: false}
		Comment := sql.NullString{String: "", Valid: false}
		err := rows.Scan(&ColumnName, &Type, &PrimaryKey, &Comment)
		checkErr(err)
		var IsPrimaryKey bool
		if PrimaryKey.Valid && util.IsNotBlank(PrimaryKey.String) {
			IsPrimaryKey = true
		} else {
			IsPrimaryKey = false
		}
		columnInfo := NewColumnInfo(ColumnName, Type, IsPrimaryKey, true, Comment.String, d.TypeList)
		resultList = append(resultList, columnInfo)
	}
	return resultList
}

//获取数据库驱动名
func GetDriverName(Type int) string {
	if Type == model.MySQLType {
		return model.MysqlName
	} else if Type == model.OracleType {
		return model.OracleName
	} else if Type == model.PostgreSQLType {
		return model.PostgreSQLName
	} else {
		return ""
	}
}

//获取数据库的连接URL
func getDataSourceName(dataSource model.ProjectDb) string {
	if dataSource.Type == model.MySQLType {
		return dataSource.Url
	} else if dataSource.Type == model.OracleType {
		return dataSource.Url
	} else if dataSource.Type == model.PostgreSQLType {
		return dataSource.Url
	} else {
		return ""
	}
}

/**
 * 以驼峰命名法生成类名，用于未指定类名时自动生成类名，如sys_user自动生成类名SysUser
 */
func TableName2ClassName(tableName string) string {
	if util.IsBlank(tableName) {
		return util.EMPTY
	}
	// 表名中不包含 "_"
	if !strings.Contains(tableName, "_") {
		return util.FirstToUpperCase(strings.ToLower(tableName))
	}
	strLists := strings.Split(strings.ToLower(tableName), "_")
	str := ""
	for i := 0; i < len(strLists); i++ {
		str += util.FirstToUpperCase(strLists[i])
	}
	return str
}

/**
 * 数据库列名转换为实体的属性名
 */
func ColumnName2PropertyName(columnName string) string {
	if util.IsBlank(columnName) {
		return util.EMPTY
	}
	// 列名中不包含 "_"
	if !strings.Contains(columnName, "_") {
		return strings.ToLower(columnName)
	}
	strLists := strings.Split(strings.ToLower(columnName), "_")
	str := ""
	for i := 0; i < len(strLists); i++ {
		str += util.FirstToUpperCase(strLists[i])
	}
	return util.FirstToLowerCase(str)
}

/**
 * 数据库列类型转换为实体的列类型
 */
func SqlTypeToColumnType(sqlType string, typeList []model.TypeColumn) string {
	for i := 0; i < len(typeList); i++ {
		typeItem := typeList[i]
		if strings.EqualFold(strings.ToLower(sqlType), strings.ToLower(typeItem.ColumnType)) {
			if strings.EqualFold("varchar", strings.ToLower(sqlType)) {
				return "VARCHAR"
			}
			if strings.EqualFold("int", strings.ToLower(sqlType)) {
				return "INTEGER"
			}
			if strings.EqualFold("datetime", strings.ToLower(sqlType)) {
				return "TIMESTAMP"
			}
			if strings.EqualFold("decimal", strings.ToLower(sqlType)) {
				return "DECIMAL"
			}
			if strings.EqualFold("text", strings.ToLower(sqlType)) {
				return "VARCHAR"
			}
			return typeItem.ColumnType
		}
	}
	return ""
}

/**
 * 数据库列类型转换为实体的属性类型
 */
func SqlTypeToFieldType(sqlType string, typeList []model.TypeColumn) string {
	for i := 0; i < len(typeList); i++ {
		typeItem := typeList[i]
		if strings.EqualFold(strings.ToLower(sqlType), strings.ToLower(typeItem.ColumnType)) {
			return typeItem.FieldType
		}
	}
	return ""
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
