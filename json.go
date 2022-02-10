package schemagen

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/funvit/schematic/field"
	"github.com/funvit/schematic/field/types"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fastjson"
)

func MarshalJSON(values map[string]interface{}, fields map[string]*field.Descriptor) ([]byte, error) {

	msi := make(map[string]interface{}, len(fields))

	for _, f := range fields {

		switch f.Type {
		case types.Bool:
			v := values[f.Name]
			msi[f.Name] = v

		case types.Time:
			v := values[f.Name].(time.Time).Unix()
			msi[f.Name] = v

		case types.UUID:
			v := values[f.Name].(uuid.UUID).String()
			msi[f.Name] = v

		case types.Bytes:
			msi[f.Name] = base64.StdEncoding.EncodeToString(values[f.Name].([]byte))

		case types.Int8, types.Int16, types.Int32, types.Int64,
			types.Uint8, types.Uint16, types.Uint32, types.Uint64,
			types.String:
			msi[f.Name] = values[f.Name]

		default:
			return nil, errors.Errorf("unknown field %s type: %s", f.Name, f.Type)
		}
	}

	return json.Marshal(msi)
}

func UnmarshalJSON(b []byte, fields map[string]*field.Descriptor) (values map[string]interface{}, err error) {

	var p fastjson.Parser

	v, err := p.ParseBytes(b)
	if err != nil {
		return nil, errors.WithMessage(err, "parse bytes")
	}

	defer func() {
		if r := recover(); r != nil {
			err = errors.Wrap(errors.New(fmt.Sprint(r)), "panic")
		}
	}()

	values = make(map[string]interface{}, len(fields))

	for _, f := range fields {
		switch f.Type {
		case types.Bool:
			values[f.Name] = v.GetBool(f.Name)

		case types.Int8, types.Int16, types.Int32, types.Int64:
			values[f.Name] = v.GetInt64(f.Name)

		case types.Uint8, types.Uint16, types.Uint32, types.Uint64:
			values[f.Name] = v.GetUint64(f.Name)

		case types.Bytes:
			enc := v.GetStringBytes(f.Name)
			b := make([]byte, base64.StdEncoding.DecodedLen(len(enc)))

			n, err := base64.StdEncoding.Decode(b, enc)
			if err != nil {
				return nil, errors.WithMessagef(err, "parsing bytes field %s", f.Name)
			}
			values[f.Name] = b[:n]

		case types.String:
			values[f.Name] = string(v.GetStringBytes(f.Name))

		case types.Time:
			values[f.Name] = time.Unix(v.GetInt64(f.Name), 0)

		case types.UUID:
			v, err := uuid.ParseBytes(v.GetStringBytes(f.Name))
			if err != nil {
				return nil, errors.WithMessagef(err, "parsing uuid.UUID field %s", f.Name)
			}

			values[f.Name] = v

		default:
			return nil, errors.Errorf("unknown model field type: %s", f.Type)
		}
	}

	return values, nil
}
