package field

// Comment set comment for field.
//
// String can be multiline.
func (f *uuidBuilder) Comment(s string) *uuidBuilder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *int8Builder) Comment(s string) *int8Builder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *int16Builder) Comment(s string) *int16Builder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *int32Builder) Comment(s string) *int32Builder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *int64Builder) Comment(s string) *int64Builder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *stringBuilder) Comment(s string) *stringBuilder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *timeBuilder) Comment(s string) *timeBuilder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *uint8Builder) Comment(s string) *uint8Builder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *uint16Builder) Comment(s string) *uint16Builder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *uint32Builder) Comment(s string) *uint32Builder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *uint64Builder) Comment(s string) *uint64Builder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *boolBuilder) Comment(s string) *boolBuilder {
	f.desc.Comment = Comment(s)

	return f
}

// Comment set comment for field.
//
// String can be multiline.
func (f *bytesBuilder) Comment(s string) *bytesBuilder {
	f.desc.Comment = Comment(s)

	return f
}
