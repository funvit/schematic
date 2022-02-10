package field

func (f *boolBuilder) Getter() *boolBuilder {
	f.desc.Getter = true
	return f
}

func (f *int8Builder) Getter() *int8Builder {
	f.desc.Getter = true
	return f
}

func (f *uint8Builder) Getter() *uint8Builder {
	f.desc.Getter = true
	return f
}

func (f *int16Builder) Getter() *int16Builder {
	f.desc.Getter = true
	return f
}

func (f *uint16Builder) Getter() *uint16Builder {
	f.desc.Getter = true
	return f
}

func (f *int32Builder) Getter() *int32Builder {
	f.desc.Getter = true
	return f
}

func (f *int64Builder) Getter() *int64Builder {
	f.desc.Getter = true
	return f
}

func (f *uint32Builder) Getter() *uint32Builder {
	f.desc.Getter = true
	return f
}

func (f *uint64Builder) Getter() *uint64Builder {
	f.desc.Getter = true
	return f
}

func (f *stringBuilder) Getter() *stringBuilder {
	f.desc.Getter = true
	return f
}

func (f *bytesBuilder) Getter() *bytesBuilder {
	f.desc.Getter = true
	return f
}

func (f *timeBuilder) Getter() *timeBuilder {
	f.desc.Getter = true
	return f
}

func (f *uuidBuilder) Getter() *uuidBuilder {
	f.desc.Getter = true
	return f
}
