package sql2struct

type DataBaseModel interface{
	Connect()error
	GetFields(dbName,tableName string)(*TableFields,error)
}

type TableFields struct{
	TableName string
	List []*Field         
}
type Field struct{
	FieldName string //字段名称
	FieldType string //字段数据库类型
	Comment string //字段描述
}