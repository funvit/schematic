package schemagen

import (
	"github.com/funvit/schematic/field"
)

type Schema struct {
	FieldsSet
}

func (_ *Schema) isSchema() {

}

type FieldsSet interface {
	Fields() []field.Field
}

type SchemaConfiger interface {
	Model() SchemaConfig
}

type SchemaConfig struct {
	JSON bool
}
