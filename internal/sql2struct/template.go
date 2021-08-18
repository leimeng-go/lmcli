package sql2struct

const(
	structTpl=`type {{.TableName | ToCamelCase}} struct`
)