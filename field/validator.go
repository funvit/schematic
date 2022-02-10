package field

type validator struct {
	// func(v interface{}) error
	Apply         interface{}
	ExpectedError string
}
