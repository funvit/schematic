package field

func (f *boolBuilder) Annotate(s string) *boolBuilder {
	f.desc.Annotate = s

	return f
}

func (f *bytesBuilder) Annotate(s string) *bytesBuilder {
	f.desc.Annotate = s

	return f
}

func (f *int8Builder) Annotate(s string) *int8Builder {
	f.desc.Annotate = s

	return f
}

func (f *int16Builder) Annotate(s string) *int16Builder {
	f.desc.Annotate = s

	return f
}

func (f *int32Builder) Annotate(s string) *int32Builder {
	f.desc.Annotate = s

	return f
}

func (f *int64Builder) Annotate(s string) *int64Builder {
	f.desc.Annotate = s

	return f
}

func (f *stringBuilder) Annotate(s string) *stringBuilder {
	f.desc.Annotate = s

	return f
}

func (f *timeBuilder) Annotate(s string) *timeBuilder {
	f.desc.Annotate = s

	return f
}

func (f *uint8Builder) Annotate(s string) *uint8Builder {
	f.desc.Annotate = s

	return f
}

func (f *uint16Builder) Annotate(s string) *uint16Builder {
	f.desc.Annotate = s

	return f
}

func (f *uint32Builder) Annotate(s string) *uint32Builder {
	f.desc.Annotate = s

	return f
}

func (f *uint64Builder) Annotate(s string) *uint64Builder {
	f.desc.Annotate = s

	return f
}

func (f *uuidBuilder) Annotate(s string) *uuidBuilder {
	f.desc.Annotate = s

	return f
}
