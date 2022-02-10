package field

import "github.com/funvit/schematic/field/types"

type bytesBuilder struct {
	desc *Descriptor
}

func Bytes(name string) *bytesBuilder {
	return &bytesBuilder{desc: &Descriptor{
		Name: name,
		Type: types.Bytes,
	}}
}

type boolBuilder struct {
	desc *Descriptor
}

func Bool(name string) *boolBuilder {
	return &boolBuilder{&Descriptor{
		Name: name,
		Type: types.Bool,
	}}
}

type int8Builder struct {
	desc *Descriptor
}

func Int8(name string) *int8Builder {
	return &int8Builder{desc: &Descriptor{
		Name: name,
		Type: types.Int8,
	}}
}

type int16Builder struct {
	desc *Descriptor
}

func Int16(name string) *int16Builder {
	return &int16Builder{desc: &Descriptor{
		Name: name,
		Type: types.Int16,
	}}
}

type int32Builder struct {
	desc *Descriptor
}

func Int32(name string) *int32Builder {
	return &int32Builder{desc: &Descriptor{
		Name: name,
		Type: types.Int32,
	}}
}

type int64Builder struct {
	desc *Descriptor
}

func Int64(name string) *int64Builder {
	return &int64Builder{desc: &Descriptor{
		Name: name,
		Type: types.Int64,
	}}
}

type stringBuilder struct {
	desc *Descriptor
}

func String(name string) *stringBuilder {
	return &stringBuilder{desc: &Descriptor{
		Name: name,
		Type: types.String,
	}}
}

type timeBuilder struct {
	desc *Descriptor
}

func Time(name string) *timeBuilder {
	return &timeBuilder{desc: &Descriptor{
		Name: name,
		Type: types.Time,
	}}
}

type uint8Builder struct {
	desc *Descriptor
}

func Uint8(name string) *uint8Builder {
	return &uint8Builder{desc: &Descriptor{
		Name: name,
		Type: types.Uint8,
	}}
}

type uint16Builder struct {
	desc *Descriptor
}

func Uint16(name string) *uint16Builder {
	return &uint16Builder{desc: &Descriptor{
		Name: name,
		Type: types.Uint16,
	}}
}

type uint32Builder struct {
	desc *Descriptor
}

func Uint32(name string) *uint32Builder {
	return &uint32Builder{desc: &Descriptor{
		Name: name,
		Type: types.Uint32,
	}}
}

type uint64Builder struct {
	desc *Descriptor
}

func Uint64(name string) *uint64Builder {
	return &uint64Builder{desc: &Descriptor{
		Name: name,
		Type: types.Uint64,
	}}
}

type uuidBuilder struct {
	desc *Descriptor
}

func UUID(name string) *uuidBuilder {
	return &uuidBuilder{desc: &Descriptor{
		Name: name,
		Type: types.UUID,
	}}
}
