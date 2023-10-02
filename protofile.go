package protorowdf

import (
	"errors"
	"io"
	"strings"
)

func (f *ProtoFile) ValidateTagNumbers() error {
	var errs []error
	for _, s := range f.Structs {
		if err := s.ValidTagNumbers(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (f *ProtoFile) WriteProtoFileDefinition(w io.Writer) error {
	f.UpdateFieldIndent()
	if err := f.ValidateTagNumbers(); err != nil {
		return err
	}
	return protoTemplate.ExecuteTemplate(w, "ProtoFile", f)
}

func (f *ProtoFile) ProtoFileDefinition() (string, error) {
	var b strings.Builder

	if err := f.WriteProtoFileDefinition(&b); err != nil {
		return "", err
	}

	return b.String(), nil
}

func (f *ProtoFile) PrettyComments() []string {
	return prettyComments(f)
}

func (f *ProtoFile) UpdateFieldIndent() {
	indent := f.FieldIndent
	if indent == "" {
		indent = "  "
	}

	for _, s := range f.Structs {
		for _, field := range s.Fields {
			if field.Indent == "" {
				field.Indent = indent
			}
		}
	}
}
