package model

import (
	"{{.module}}/pkg/basemodel"
)

const TableName{{.struct}} = "{{.tableName}}"

type {{.struct}} struct {
	basemodel.Model
  {{range .fields -}}
	{{.Name}} {{.Type}} `{{range $tagName, $tagValues := .Tag}}{{$tagName}}:"{{$tagValues}}"{{end}}`
	{{end}}
	// Associations code here, example:
	//   Langs []Lang `gorm:"many2many:users_langs" json:"langs,omitempty"`
	//   Roles []Role `gorm:"foreignKey:role_id" json:"roles,omitempty"`
}

func (*{{.struct}}) TableName() string {
	return TableName{{.struct}}
}
