package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type MysqlToStruct struct{
	DBModel
	
}
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}



type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

func NewMysqlDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}
func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=true&loc=Local", m.DBInfo.UserName, m.DBInfo.Password, m.DBInfo.Host, m.DBInfo.Charset)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetFields(dbName, tableName string) (*TableFields, error) {
	query := `SELECT COLUMN_NAME,DATA_TYPE,COLUMN_KEY,IS_NULLABLE,COLUMN_TYPE,COLUMN_COMMENT FROM COLUMNS WHERE TABLE_SCHEMA=? AND TABLE_NAME=?`
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn

	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	list:=make([]*Field,0)
	for _,v:=range columns {
		element:=&Field{
			FieldName: v.ColumnName,
			FieldType: getMysqlMapType(v.DataType),
			Comment:   v.ColumnComment,
			Tags: []string{"json"},
		}
		element.TagStr=element.GetTags()
		list=append(list, element )
	}
	return &TableFields{
		tableName,
		list,
	}, nil
}
