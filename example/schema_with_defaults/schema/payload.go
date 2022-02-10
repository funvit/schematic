package schema

import (
	"time"

	schemagen "github.com/funvit/schematic"
	"github.com/funvit/schematic/field"
	"github.com/google/uuid"
)

//go:generate schematic model . --target ../model Payload
type Payload struct {
	schemagen.Schema
}

func (Payload) Fields() []field.Field {
	return []field.Field{
		field.Int8("int8").
			Getter().
			Private().
			Default(8).
			ModelStringer(),

		field.Int16("int16").
			Getter().
			Private().
			Default(16).
			ModelStringer(),

		field.Int32("int32").
			Getter().
			Private().
			Default(32).
			ModelStringer(),

		field.Int64("int64").
			Getter().
			Private().
			Default(64).
			ModelStringer(),

		field.Uint8("uint8").
			Getter().
			Private().
			Default(18).
			ModelStringer(),

		field.Uint16("uint16").
			Getter().
			Private().
			Default(116).
			ModelStringer(),

		field.Uint32("uint32").
			Getter().
			Private().
			Default(132).
			ModelStringer(),

		field.Uint64("uint64").
			Getter().
			Private().
			Default(164).
			ModelStringer(),

		field.Bool("bool").
			Getter().
			Private().
			Default(true).
			ModelStringer(),

		field.String("string").
			Private().
			Default("default string value").
			Getter().
			ModelStringer(),

		field.Bytes("bytes").
			Private().
			Default([]byte("some bytes")).
			Getter(),

		field.UUID("uuid").
			Private().
			Getter().
			Default(uuid.MustParse("d540fd1b-8d05-4727-b3ff-56274ad606ee")),

		field.Time("time").Private().Getter().Default(time.Unix(100, 200)),
	}
}
