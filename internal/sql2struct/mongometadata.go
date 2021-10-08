package sql2struct

var (
	mongoToGoType = map[string]string{
		"string": "string",
		"double": "float64",
		"bool":   "bool",
		"Date":   "time.time",
		"null":   "null",
		"int":    "int32",
		"object": "interface{}",
		"long":   "int64",
		"array":  "[]interface{}",
		"date":   "time.Time",
	}
)

func getMongoMapType(typeMongo string) string {
	return mongoToGoType[typeMongo]
}
