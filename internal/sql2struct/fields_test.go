package sql2struct

import "testing"

func TestField_GetTags(t *testing.T) {
	f := &Field{
		FieldName: "test",
		FieldType: "string",
		Comment:   "这是一个测试的字段",
		Tags:      []string{"json", "bson"},
	}
	t.Log(f.GetTags())
}
