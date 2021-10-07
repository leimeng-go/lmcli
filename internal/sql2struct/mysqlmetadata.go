package sql2struct


var(
	mysqlToGoType=map[string]string{
		"int":       "int32",
		"tinyint":   "int8",
		"smallint":  "int",
		"mediumint": "int64",
		"bigint":    "int64",
		"bit":       "int",
		"bool":      "bool",
		"enum":      "string",
		"set":  "string",
		"varchar":"string",
	}
)

func getMysqlMapType(typeMysql string)string{
	return mysqlToGoType[typeMysql]
}