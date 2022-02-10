package types

// A Type represents a field type.
type Type string

// List of field types.
const (
	Invalid = "invalid"
	Bool    = "bool"
	Time    = "time.Time"
	UUID    = "uuid.UUID"
	Bytes   = "[]byte"
	String  = "string"
	Int8    = "int8"
	Int16   = "int16"
	Int32   = "int32"
	Int64   = "int64"
	Uint8   = "uint8"
	Uint16  = "uint16"
	Uint32  = "uint32"
	Uint64  = "uint64"
)

func init() {
	_ = Invalid
}

func (t Type) String() string {
	return string(t)
}
