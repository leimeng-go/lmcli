package sql2struct

import (
	"html/template"
	"lmcli/internal/word"
	"os"
)

const (
	structTpl = `
   type {{ .TableName|ToUpperCamelCase}} struct{
     {{- range .List }} 
     {{.FieldName | ToUpperCamelCase}} {{if .FieldType}} {{.FieldType}} {{else}} {{end}} {{if .TagStr}} {{.TagStr}} {{else}} {{end}} {{if .Comment}} //{{.Comment}} {{else}} {{end}}
     {{- end}}
   }
  `
)

type StructTemplate struct {
	structTpl string
}

func (t *StructTemplate) Generate(data *TableFields) error {
	t.structTpl = structTpl
	tpl := template.Must(template.New("sql2struct").Funcs(map[string]interface{}{
		"ToUpperCamelCase": word.CamelCaseToUnderscore,
	}).Parse(t.structTpl))
	err := tpl.Execute(os.Stdout, &data)
	if err != nil {
		return err
	}
	return nil
}
