package schema

import (
	schemagen "github.com/funvit/schematic"
	"github.com/funvit/schematic/field"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

//go:generate schematic model . --target ../model Account
type Account struct {
	schemagen.Schema
}

func (Account) Fields() []field.Field {
	return []field.Field{
		field.Int64("id").
			Getter().
			Private().
			Default(0).
			Comment("Value auto-generated in repository.").
			ModelStringer(),

		field.String("login").
			NotEmpty().
			ModelStringer().
			Getter().
			Setter(),

		field.String("name").
			Getter().
			Setter().
			ModelStringer().
			ValidatorFuncWithError(
				func(v string) error {
					if len(v) <= 3 {
						return errors.New("must have len greater than 3")
					}
					return nil
				},
				"must have len greater than 3",
			),

		field.String("password").
			Sensitive().
			Comment("Value expected to be hashed.").
			ModelStringer(),

		field.Bool("is_deleted").
			Comment("Used for soft-delete.").
			Getter().
			Setter(),

		field.UUID("rev").
			NotEmpty().
			ModelStringer(),

		field.UUID("some_external_uuid").ValidatorFunc(func(v uuid.UUID) error {
			// dummy
			return nil
		}),

		field.Uint8("color"),

		field.Int8("my_byte"),

		field.Bytes("binary"),
		field.Time("created_at").
			NotEmpty().
			ModelStringer(),
	}
}
