package schema

import (
	schemagen "github.com/funvit/schematic"
	"github.com/funvit/schematic/field"
)

//go:generate schematic model . --target ../model Account
type Account struct {
	schemagen.Schema
}

func (Account) Fields() []field.Field {
	// Note: no validators!
	return []field.Field{
		field.Int64("id").Comment("Value auto-generated in repository.").Annotate(`json:"id"`),

		field.String("login").Annotate(`json:"login"`),

		field.String("name"),

		field.String("password").Comment("Value expected to be hashed."),

		field.Bool("is_deleted").Comment("Used for soft-delete."),

		field.UUID("rev"),

		field.UUID("some_external_uuid"),

		field.Uint8("color"),

		field.Int8("my_byte"),

		field.Bytes("binary"),
		field.Time("created_at"),
	}
}
