package sql2struct

import (
	"html/template"
	"os"
	"strings"
	"testing"
)

func TestStructTemplate_Generate(t *testing.T) {
	tpl,err:=template.New("test").Funcs(map[string]interface{}{
		"ToUpperCamelCase":strings.Title,
	}).Parse(structTpl)
	if err!=nil{
		t.Error(err.Error())
		return
	}
	data:= TableFields{
		TableName: "mytest",
		List: []*Field{
			{
				FieldName: "name",
				FieldType: "string",
				Comment:   "用户的姓名",
				Tags: []string{"bson","json","validate"},
			},
			{
				FieldName: "age",
				FieldType: "int",
				Comment:   "用户的年龄",
				Tags:      []string{"bson","json"},
			},
		},
	}
	for i,_:=range data.List {
		data.List[i].TagStr=data.List[i].GetTags()
	}
	tpl.Execute(os.Stdout,&data)
}
