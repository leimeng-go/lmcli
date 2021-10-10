package sql2struct

import (
	"os"
	"sort"
	"strings"
	"testing"
	"text/template"
)

func TestStructTemplate_Generate(t *testing.T) {
	tpl, err := template.New("test").Funcs(map[string]interface{}{
		"ToUpperCamelCase": strings.Title,
	}).Parse(structTpl)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data := TableFields{
		TableName: "mytest",
		List: FieldList{
			{
				FieldName: "name",
				FieldType: "string",
				Comment:   "用户的姓名",
				Tags:      []string{"bson"},
			},
			{
				FieldName: "age_state",
				FieldType: "int",
				Comment:   "用户的年龄",
				Tags:      []string{"bson"},
			},
		},
	}
	list:=data.List
	sort.Sort(list)
	t.Log(list)
	for i, _ := range data.List {
		data.List[i].TagStr = data.List[i].GetTags()
	}
	tpl.Execute(os.Stdout, &data)
}
func TestStrCompare(t *testing.T){
	a:= "name"
	b:= "age"
	t.Log(a[0]>b[0])
}
