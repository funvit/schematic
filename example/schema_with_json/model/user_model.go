// Code generated by schematic, DO NOT EDIT.
//
// Schema file: github.com/funvit/schematic/example/schema_with_json/schema/user.go

package model

import (
	"github.com/google/uuid"
	"time"

	schemagen "github.com/funvit/schematic"
	"github.com/pkg/errors"
)

type User struct {
	Id        int64
	Name      string
	IsDeleted bool
	rev       uuid.UUID
	Binary    []byte
	CreatedAt time.Time
	Int32     int32
}

// GetRev gets field rev value.
//
// Generated due to schema field marked with Getter().
//
// Note: getter can help to satisfy some interface(s).
func (m *User) GetRev() uuid.UUID {
	return m.rev
}

// UserWithDefaults constructs model User
// with default values defined in schema.
//
// Example:
//
//    // construct a new model
//    myModel := UserWithDefaults()
//
//    // set public fields directly
//    myModel.SomePublicField = 4221
//
//    // or via setter (if generated)
//    myModel.SetName("Bob")
//
func UserWithDefaults() *User {
	return &User{
		rev:   uuid.MustParse("db6feed5-3225-434b-861e-42bdee8ddce3"),
		Int32: 1,
	}
}

// MarshalJSON .
//
// Generated due to schema config JSON set to true.
//
// Note: schema names are used (Annotate json tags not used).
func (m *User) MarshalJSON() ([]byte, error) {

	values := make(map[string]interface{}, 7)
	values["id"] = m.Id
	values["name"] = m.Name
	values["is_deleted"] = m.IsDeleted
	values["rev"] = m.rev
	values["binary"] = m.Binary
	values["created_at"] = m.CreatedAt
	values["int32"] = m.Int32

	return schemagen.MarshalJSON(values, userModelFields)
}

// UnmarshalJSON .
//
// Generated due to schema config JSON set to true.
//
// Note: schema names are used (Annotate json tags not used).
func (m *User) UnmarshalJSON(b []byte) error {

	vs, err := schemagen.UnmarshalJSON(b, userModelFields)
	if err != nil {
		return err
	}

	{
		v, ok := vs["id"].(int64)
		if !ok {
			return errors.Errorf("json key %s is missing", "id")
		}

		m.Id = int64(v)
	}

	{
		v, ok := vs["name"].(string)
		if !ok {
			return errors.Errorf(
				"json key %s have unexpected type %T",
				"name",
				vs["name"],
			)
		}

		m.Name = v
	}

	{
		v, ok := vs["is_deleted"].(bool)
		if !ok {
			return errors.Errorf(
				"json key %s have unexpected type %T",
				"is_deleted",
				vs["is_deleted"],
			)
		}

		m.IsDeleted = v
	}

	{
		v, ok := vs["rev"].(uuid.UUID)
		if !ok {
			return errors.Errorf(
				"json key %s have unexpected type %T",
				"rev",
				vs["rev"],
			)
		}

		m.rev = v
	}

	{
		v, ok := vs["binary"].([]byte)
		if !ok {
			return errors.Errorf(
				"json key %s have unexpected type %T",
				"binary",
				vs["binary"],
			)
		}

		if len(v) > 0 {
			m.Binary = v
		}
	}

	{
		v, ok := vs["created_at"].(time.Time)
		if !ok {
			return errors.Errorf(
				"json key %s have unexpected type %T",
				"created_at",
				vs["created_at"],
			)
		}

		m.CreatedAt = v
	}

	{
		v, ok := vs["int32"].(int64)
		if !ok {
			return errors.Errorf("json key %s is missing", "int32")
		}

		m.Int32 = int32(v)
	}

	return nil
}
