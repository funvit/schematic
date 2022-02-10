package field

import "errors"

func (f *int64Builder) Positive() *int64Builder {
	f.ValidatorFunc(func(v int64) error {
		if v < 0 {
			return errors.New("must be positive")
		}
		return nil
	})

	return f
}

func (f *int32Builder) Positive() *int32Builder {
	f.ValidatorFunc(func(v int32) error {
		if v < 0 {
			return errors.New("must be positive")
		}
		return nil
	})

	return f
}

func (f *int8Builder) Positive() *int8Builder {
	f.ValidatorFunc(func(v int8) error {
		if v < 0 {
			return errors.New("must be positive")
		}
		return nil
	})

	return f
}

func (f *int16Builder) Positive() *int16Builder {
	f.ValidatorFunc(func(v int16) error {
		if v < 0 {
			return errors.New("must be positive")
		}
		return nil
	})

	return f
}
