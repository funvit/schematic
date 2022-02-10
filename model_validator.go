package schemagen

import (
	"reflect"
	"strings"
	"time"

	"github.com/funvit/schematic/field"
	"github.com/funvit/schematic/field/types"
	"github.com/funvit/schematic/pkg/scope_error"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func ValidateModel(
	m interface{},
	fields map[string]*field.Descriptor,
	fieldErrors map[string]*scope_error.ScopeError,
) error {

	var mVal reflect.Value

	switch reflect.TypeOf(m).Kind() {
	case reflect.Struct:
		mVal = reflect.ValueOf(m)
	case reflect.Ptr:
		mVal = reflect.ValueOf(m).Elem()
	default:
		return errors.New("struct of pointer expected")
	}

	for _, f := range fields {
		var fn string

		if f.Private {
			fn = SnakeToCamel(f.Name)
		} else {
			fn = strings.Title(SnakeToCamel(f.Name))
		}

		mf := mVal.FieldByName(fn)

		for _, validator := range f.Validators {
			var err error

			switch f.Type {
			case types.Bool:
				// not so useful :-)
				err = validator.Apply.(func(bool2 bool) error)(mf.Bool())

			case types.Int8:
				err = validator.Apply.(func(int8) error)(int8(mf.Int()))
			case types.Uint8:
				err = validator.Apply.(func(uint8) error)(uint8(mf.Uint()))

			case types.Int16:
				err = validator.Apply.(func(int16) error)(int16(mf.Int()))
			case types.Uint16:
				err = validator.Apply.(func(uint16) error)(uint16(mf.Uint()))

			case types.Int32:
				err = validator.Apply.(func(int32) error)(int32(mf.Int()))
			case types.Uint32:
				err = validator.Apply.(func(uint32) error)(uint32(mf.Uint()))

			case types.Int64:
				err = validator.Apply.(func(int64) error)(mf.Int())
			case types.Uint64:
				err = validator.Apply.(func(uint64) error)(mf.Uint())

			case types.Bytes:
				err = validator.Apply.(func([]byte) error)(mf.Bytes())
			case types.String:
				err = validator.Apply.(func(string) error)(mf.String())

			case types.Time:
				err = validator.Apply.(func(time.Time) error)(mf.Interface().(time.Time))

			case types.UUID:
				err = validator.Apply.(func(uuid.UUID) error)(mf.Interface().(uuid.UUID))

			default:
				return errors.Errorf("validate model: unknown field type: %s", f.Type)
			}

			if err != nil {
				fe, ok := fieldErrors[f.Name]
				if ok {
					return ErrModelValidationError.WithCause(
						ErrModelFieldValidationError.WithCause(
							fe.WithCause(err)))
				}
				return ErrModelValidationError.WithCause(
					ErrModelFieldValidationError.WithCause(
						errors.WithMessage(err, fn)))
			}
		}
	}

	return nil
}
