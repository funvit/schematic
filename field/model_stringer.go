package field

func (f *boolBuilder) ModelStringer() *boolBuilder {
	f.desc.ModelStringer = true
	return f
}

func (f *int8Builder) ModelStringer() *int8Builder {
	f.desc.ModelStringer = true
	return f
}

func (f *uint8Builder) ModelStringer() *uint8Builder {
	f.desc.ModelStringer = true
	return f
}

func (f *int16Builder) ModelStringer() *int16Builder {
	f.desc.ModelStringer = true
	return f
}

func (f *uint16Builder) ModelStringer() *uint16Builder {
	f.desc.ModelStringer = true
	return f
}

func (f *int32Builder) ModelStringer() *int32Builder {
	f.desc.ModelStringer = true
	return f
}

func (f *int64Builder) ModelStringer() *int64Builder {
	f.desc.ModelStringer = true
	return f
}

func (f *uint32Builder) ModelStringer() *uint32Builder {
	f.desc.ModelStringer = true
	return f
}

func (f *uint64Builder) ModelStringer() *uint64Builder {
	f.desc.ModelStringer = true
	return f
}

func (f *stringBuilder) ModelStringer() *stringBuilder {
	f.desc.ModelStringer = true
	return f
}

func (f *timeBuilder) ModelStringer() *timeBuilder {
	f.desc.ModelStringer = true
	return f
}

func (f *uuidBuilder) ModelStringer() *uuidBuilder {
	f.desc.ModelStringer = true
	return f
}
