package schema

import (
	schemagen "github.com/funvit/schematic"
	"github.com/funvit/schematic/field"
	"github.com/google/uuid"
)

//go:generate schematic model . --target ../model User
type User struct {
	schemagen.Schema
}

func (User) Fields() []field.Field {
	return []field.Field{
		field.Int64("id"),

		field.String("name"),

		field.Bool("is_deleted"),

		field.UUID("rev").Private().
			Default(uuid.MustParse("db6feed5-3225-434b-861e-42bdee8ddce3")).
			Getter(),

		field.Bytes("binary"),

		field.Time("created_at"),

		field.Int32("int32").Default(1),
	}
}

func (User) Model() schemagen.SchemaConfig {
	return schemagen.SchemaConfig{
		JSON: true,
	}
}
