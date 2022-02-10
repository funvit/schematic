package field

import (
	"time"

	"github.com/google/uuid"
)

func (f *boolBuilder) Default(v bool) *boolBuilder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *bytesBuilder) Default(v []byte) *bytesBuilder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *int16Builder) Default(v int16) *int16Builder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *int8Builder) Default(v int8) *int8Builder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *int32Builder) Default(v int32) *int32Builder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *int64Builder) Default(v int64) *int64Builder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *stringBuilder) Default(v string) *stringBuilder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *timeBuilder) Default(v time.Time) *timeBuilder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *uint8Builder) Default(v uint8) *uint8Builder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *uint16Builder) Default(v uint16) *uint16Builder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *uint32Builder) Default(v uint32) *uint32Builder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *uint64Builder) Default(v uint64) *uint64Builder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}

func (f *uuidBuilder) Default(v uuid.UUID) *uuidBuilder {
	f.desc.Default = true
	f.desc.DefaultValue = v

	return f
}
