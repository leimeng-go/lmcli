package sql2struct

type DataBaseModel interface {
	Connect() error
	GetFields(dbName, tableName string) (*TableFields, error)
}

//DBInfo db相关信息
//todo 使用uri方式会不会更好呢？
type DBInfo struct {
	UserName  string
	Password  string
	Host      string
	Port      int
	Charset   string
	DBType    string
	DBName    string
	TableName string
}
