package sql2struct

import (
	"fmt"
	"lmcli/internal/word"
	"sort"
	"strings"
)

type TableFields struct{
	TableName string
	List []*Field
}
type Field struct{
	FieldName string  `json:"field_name" `//字段名称
	FieldType string `json:"field_type"` //字段数据库类型
	Comment string //字段描述
	Tags []string //tag集合
	TagStr string
}
//GetTags 生成结构体tag部分
//`json:"" bson:""`
func (this Field)GetTags()string {
	buf:=strings.Builder{}
	buf.WriteString("`")
	//根据ASCII表排序字符串列表
	sort.Strings(this.Tags)
	for i,v:=range this.Tags{
		if len(this.Tags) == i+1 {
			buf.WriteString(fmt.Sprintf(`%s:%s`,v,this.FormatFieldByTag(v)))
		}else {
			buf.WriteString(fmt.Sprintf(`%s:%s `,v,this.FormatFieldByTag(v)))
		}
	}
	buf.WriteString("`")
	return buf.String()
}

func (this Field)FormatFieldByTag(tag string)(formatStr string) {
	switch tag {
	case "json":
		formatStr=word.UnderscoreToLowerCamelCaseV1(this.FieldName)
	case "bson":
		formatStr=strings.ToLower(this.FieldName)
	default:
		formatStr="为定义tag"
	}
	return
}