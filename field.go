package protorowdf

import "io"

func (f *Field) PrettyComments() []string {
	return prettyComments(f)
}

func (f *Field) WriteFieldDefinition(w io.Writer) error {
	return protoTemplate.ExecuteTemplate(w, "Field", f)
}

func (f *Field) CleanPluralName() string {
	return getPluralName(f)
}

func (f *Field) GoType() string {
	return getOrPanic(f.DataType.GoType())
}

func (f *Field) RustType() string {
	return getOrPanic(f.DataType.RustType())
}

func (f *Field) GoName() string {
	return GoProtoName(f.Name)
}

func (f *Field) GoPluralName() string {
	return GoProtoName(f.CleanPluralName())
}

func (f *Field) RustNeedClone() string {
	switch f.DataType {
	case SupportedType_Bytes, SupportedType_String:
		return ".clone()"
	default:
		return ""
	}
}
