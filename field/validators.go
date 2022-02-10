package field

import (
	"time"

	"github.com/google/uuid"
)

func (f *uuidBuilder) ValidatorFunc(fn func(v uuid.UUID) error) *uuidBuilder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *uuidBuilder) ValidatorFuncWithError(fn func(v uuid.UUID) error, err string) *uuidBuilder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}

func (f *int8Builder) ValidatorFunc(fn func(v int8) error) *int8Builder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *int8Builder) ValidatorFuncWithError(fn func(v int8) error, err string) *int8Builder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}

func (f *int16Builder) ValidatorFunc(fn func(v int16) error) *int16Builder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *int16Builder) ValidatorFuncWithError(fn func(v int16) error, err string) *int16Builder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}

func (f *int32Builder) ValidatorFunc(fn func(v int32) error) *int32Builder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *int32Builder) ValidatorFuncWithError(fn func(v int32) error, err string) *int32Builder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}

func (f *int64Builder) ValidatorFunc(fn func(v int64) error) *int64Builder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *int64Builder) ValidatorFuncWithError(fn func(v int64) error, err string) *int64Builder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}
func (f *stringBuilder) ValidatorFunc(fn func(v string) error) *stringBuilder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *stringBuilder) ValidatorFuncWithError(fn func(v string) error, err string) *stringBuilder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}

func (f *bytesBuilder) ValidatorFunc(fn func(v []byte) error) *bytesBuilder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *bytesBuilder) ValidatorFuncWithError(fn func(v []byte) error, err string) *bytesBuilder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}

func (f *timeBuilder) ValidatorFunc(fn func(v time.Time) error) *timeBuilder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *timeBuilder) ValidatorFuncWithError(fn func(v time.Time) error, err string) *timeBuilder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}

func (f *uint8Builder) ValidatorFunc(fn func(v uint8) error) *uint8Builder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *uint8Builder) ValidatorFuncWithError(fn func(v uint8) error, err string) *uint8Builder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}

func (f *uint16Builder) ValidatorFunc(fn func(v uint16) error) *uint16Builder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *uint16Builder) ValidatorFuncWithError(fn func(v uint16) error, err string) *uint16Builder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}

func (f *uint32Builder) ValidatorFunc(fn func(v uint32) error) *uint32Builder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *uint32Builder) ValidatorFuncWithError(fn func(v uint32) error, err string) *uint32Builder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}

func (f *uint64Builder) ValidatorFunc(fn func(v uint64) error) *uint64Builder {

	return f.ValidatorFuncWithError(fn, "")
}

// ValidatorFuncWithError sets validator func with expected error string.
func (f *uint64Builder) ValidatorFuncWithError(fn func(v uint64) error, err string) *uint64Builder {

	f.desc.Validators = append(f.desc.Validators, validator{
		Apply:         fn,
		ExpectedError: err,
	})

	return f
}
