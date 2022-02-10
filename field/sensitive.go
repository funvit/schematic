package field

func (f *stringBuilder) Sensitive() *stringBuilder {
	f.desc.Sensitive = true

	return f
}
