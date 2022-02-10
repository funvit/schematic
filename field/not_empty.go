package field

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (f *uuidBuilder) NotEmpty() *uuidBuilder {
	f.ValidatorFunc(func(v uuid.UUID) error {
		if v == uuid.Nil {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *bytesBuilder) NotEmpty() *bytesBuilder {
	f.ValidatorFunc(func(v []byte) error {
		if len(v) == 0 {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *int8Builder) NotEmpty() *int8Builder {
	f.ValidatorFunc(func(v int8) error {
		if v == 0 {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *int16Builder) NotEmpty() *int16Builder {
	f.ValidatorFunc(func(v int16) error {
		if v == 0 {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *int32Builder) NotEmpty() *int32Builder {
	f.ValidatorFunc(func(v int32) error {
		if v == 0 {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *int64Builder) NotEmpty() *int64Builder {
	f.ValidatorFunc(func(v int64) error {
		if v == 0 {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *stringBuilder) NotEmpty() *stringBuilder {
	f.ValidatorFunc(func(v string) error {
		if strings.TrimSpace(v) == "" {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *timeBuilder) NotEmpty() *timeBuilder {
	f.ValidatorFunc(func(v time.Time) error {
		if v.IsZero() {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *uint8Builder) NotEmpty() *uint8Builder {
	f.ValidatorFunc(func(v uint8) error {
		if v == 0 {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *uint16Builder) NotEmpty() *uint16Builder {
	f.ValidatorFunc(func(v uint16) error {
		if v == 0 {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *uint32Builder) NotEmpty() *uint32Builder {
	f.ValidatorFunc(func(v uint32) error {
		if v == 0 {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}

func (f *uint64Builder) NotEmpty() *uint64Builder {
	f.ValidatorFunc(func(v uint64) error {
		if v == 0 {
			return errors.New("must be not empty")
		}
		return nil
	})

	return f
}
