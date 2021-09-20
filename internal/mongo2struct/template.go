package mongo2struct

import (
	"fmt"
	"lmcli/internal/word"
	"os"
	"text/template"
)

const (
	structTpl = `
	 type {{.CollectionName|ToCamelCase}} struct{
		 {{range .List}} {{$length:=len .Description}} {{if gt $length 0}} // {{Discription}} 
		 {{end}}
	 }
	`
)
var (
	MongoDBTypeToStructType=map[string]string{
		"string": "string",
		"double": "float64",
		"bool":   "bool",
		"Date":   "time.time",
		"null":   "null",
		"int":    "int32",
		"object": "interface{}",
		"long":   "int64",
		"array":  "[]interface{}",
	}
)


func TypeToGo(typeMongo string)string{
	return MongoDBTypeToStructType[typeMongo]
}

type StructTemplate struct{
	structTpl string
}

type StructColumn struct{
	Name string 
	Type string `json:"type" bson:"type"`
	Tag string //json with bson tag
	Comment string
}

type StructTemplateDB struct{
	CollectionName string
	Columns []*StructColumn
}

func NewStructTemplate()*StructTemplate{
	return &StructTemplate{
		structTpl: structTpl,
	}
}

func (t *StructTemplate)AssemblyColumns(collection *CollectionStruct)[]*StructColumn{
	tplColumns:=make([]*StructColumn,0,len(collection.List))
	for _,column:=range collection.List{
		tplColumns=append(tplColumns, &StructColumn{
			Name: column.ColumnKey,
			Type: column.ColumnType,
			Tag: fmt.Sprintf(`"json:"+"%s"+" "+"bson:"+"%s"`,word.UnderscoreToUpperCamelCase(column.ColumnKey),column.ColumnKey),
			Comment: column.Description,
		})
	}
	return tplColumns
}

func (t *StructTemplate)Generate(collectionName string,tplColumns []*StructColumn)error{
	tpl:=template.Must(template.New("mongostruct").Funcs(template.FuncMap{
		"ToCamelCase":word.UnderscoreToUpperCamelCase,
		"TypeToGo": TypeToGo,
	}).Parse(t.structTpl))

	tplDB:=StructTemplateDB{
        CollectionName:collectionName,
		Columns: tplColumns,
	}
    
	err:=tpl.Execute(os.Stdout,tplDB)
	if err!=nil{
		return err
	}
	return nil
}