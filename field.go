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
