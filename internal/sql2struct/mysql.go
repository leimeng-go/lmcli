package sql2struct

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo DBInfo
}

type DBInfo struct{
	DBType string
	Host string 
    UserName string 
	Password string
	Charset string
}

type TableColumn struct{
	ColumnName string
	DataType string
	IsNullable string
	ColumnKey string
	ColumnType string
	ColumnComment string 
}

func (m *DBModel)Connect()error{
	var err error
	dsn:=fmt.Sprintf("%s:%s@tcp(%s)/infomation_schema?charset=%s&parseTime=true&loc=Local",m.DBInfo.UserName,m.DBInfo.Password,m.DBInfo.Host,m.DBInfo.Charset)
	m.DBEngin,err=sql.Open(m.DBInfo.DBType,dsn)
	if err!=nil{
		return err
	}
	return nil
}