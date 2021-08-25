package sql2struct

import (
	"fmt"
	"lmcli/internal/word"
	"os"
	"text/template"
)

const (
	structTpl = `
	type {{.TableName|ToCamelCase}} struct {
	   {{range .Columns}} {{$length:=len .Comment}} {{if gt $length 0}} //{{.Comment}} {{else}} // {{.Name}} {{end}}
	      {{ $typeLen:= len .Type}} {{if gt $typeLen 0}} {{.Name|ToCamelCase}} {{.Type | TypeToGO}} {{.Tag}} {{else}} {{.Name}} {{end}}
	   {{end}}	 
	}
	
	func (model {{.TableName | ToCamelCase}}) TableName() string{
		return "{{.TableName}}"
	}
	`
)
var (
	DBTypeToStructType=map[string]string{
		"int":"int32",
		"tinyint":"int8",
		"smallint":"int",
		"mediumint":"int64",
		"bigint":"int64",
		"bit":"int",
		"bool":"bool",
		"enum":"string",
		"set":"string",
		"varchar":"string",
	}
)

func TypeToGO(typeMysql string)string{
	return DBTypeToStructType[typeMysql]
}
type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}
func NewStructTemplate()*StructTemplate{
	return &StructTemplate{
		structTpl: structTpl,
	}
}
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name: column.ColumnName,
			Type: column.DataType,
			Tag:  fmt.Sprintf("`json:"+"%s"+"`",column.ColumnName),
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}
func (t *StructTemplate)Generate(tableName string,tplColumns []*StructColumn)error{
	tpl:=template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase":word.UnderscoreToUpperCamelCase,
		"TypeToGO":TypeToGO,
	}).Parse(t.structTpl))

	tplDB:=StructTemplateDB{
		TableName: tableName,
		Columns: tplColumns,
	}

	err:=tpl.Execute(os.Stdout,tplDB)
	if err!=nil{
		return err
	}
	return nil
}