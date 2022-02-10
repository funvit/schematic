package field

func (f *boolBuilder) Private() *boolBuilder {
	f.desc.Private = true
	return f
}

func (f *int8Builder) Private() *int8Builder {
	f.desc.Private = true
	return f
}

func (f *uint8Builder) Private() *uint8Builder {
	f.desc.Private = true
	return f
}

func (f *int16Builder) Private() *int16Builder {
	f.desc.Private = true
	return f
}

func (f *uint16Builder) Private() *uint16Builder {
	f.desc.Private = true
	return f
}

func (f *int32Builder) Private() *int32Builder {
	f.desc.Private = true
	return f
}

func (f *int64Builder) Private() *int64Builder {
	f.desc.Private = true
	return f
}

func (f *uint32Builder) Private() *uint32Builder {
	f.desc.Private = true
	return f
}

func (f *uint64Builder) Private() *uint64Builder {
	f.desc.Private = true
	return f
}

func (f *stringBuilder) Private() *stringBuilder {
	f.desc.Private = true
	return f
}

func (f *bytesBuilder) Private() *bytesBuilder {
	f.desc.Private = true
	return f
}

func (f *timeBuilder) Private() *timeBuilder {
	f.desc.Private = true
	return f
}

func (f *uuidBuilder) Private() *uuidBuilder {
	f.desc.Private = true
	return f
}
