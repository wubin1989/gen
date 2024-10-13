package template

// Model used as a variable because it cannot load template file after packed, params still can pass file
const Model = NotEditMark + `
package {{.StructInfo.Package}}

import (
	"{{.ConfigPackage}}"
	"fmt"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"

	"encoding/json"
	"time"

	"github.com/wubin1989/datatypes"
	"github.com/wubin1989/gorm"
	"github.com/wubin1989/gorm/schema"
	{{range .ImportPkgPaths}}{{.}} ` + "\n" + `{{end}}
)

// {{.ModelStructName}} {{.StructComment}}
type {{.ModelStructName}} struct {
    {{range .Fields}}
    {{if .MultilineComment -}}
	/*
{{.ColumnComment}}
    */
	{{end -}}
    {{.Name}} {{.Type | convert}} ` + "`{{.Tags}}` " +
	"{{if not .MultilineComment}}{{if .ColumnComment}}// {{.ColumnComment}}{{end}}{{end}}" +
	`{{end}}
}

// TableName {{.ModelStructName}}'s table name
func (*{{.ModelStructName}}) TableName() string {
	{{if .TableName -}}var TableName{{.ModelStructName}} string{{- end}}
	
	if stringutils.IsNotEmpty(config.G_Config.Db.Name) {
		TableName{{.ModelStructName}} = fmt.Sprintf("%s.{{.TableName}}", config.G_Config.Db.Name)
	} else {
		TableName{{.ModelStructName}} = "{{.TableName}}"
	}

	return TableName{{.ModelStructName}}
}
`

// ModelMethod model struct DIY method
const ModelMethod = `

{{if .Doc -}}// {{.DocComment -}}{{end}}
func ({{.GetBaseStructTmpl}}){{.MethodName}}({{.GetParamInTmpl}})({{.GetResultParamInTmpl}}){{.Body}}
`
