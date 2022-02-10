package field

func (f *boolBuilder) Setter() *boolBuilder {
	f.desc.Setter = true
	return f
}

func (f *int8Builder) Setter() *int8Builder {
	f.desc.Setter = true
	return f
}

func (f *uint8Builder) Setter() *uint8Builder {
	f.desc.Setter = true
	return f
}

func (f *int16Builder) Setter() *int16Builder {
	f.desc.Setter = true
	return f
}

func (f *uint16Builder) Setter() *uint16Builder {
	f.desc.Setter = true
	return f
}

func (f *int32Builder) Setter() *int32Builder {
	f.desc.Setter = true
	return f
}

func (f *int64Builder) Setter() *int64Builder {
	f.desc.Setter = true
	return f
}

func (f *uint32Builder) Setter() *uint32Builder {
	f.desc.Setter = true
	return f
}

func (f *uint64Builder) Setter() *uint64Builder {
	f.desc.Setter = true
	return f
}

func (f *stringBuilder) Setter() *stringBuilder {
	f.desc.Setter = true
	return f
}

func (f *bytesBuilder) Setter() *bytesBuilder {
	f.desc.Setter = true
	return f
}

func (f *timeBuilder) Setter() *timeBuilder {
	f.desc.Setter = true
	return f
}

func (f *uuidBuilder) Setter() *uuidBuilder {
	f.desc.Setter = true
	return f
}
