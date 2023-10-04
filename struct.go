package protorowdf

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func (s *Struct) MessageDefinition(w io.Writer) error {
	if err := s.ValidTagNumbers(); err != nil {
		return err
	}
	return protoTemplate.Execute(w, s)
}

type DuplicateTagNumberError struct {
	MessageName string
	Duplicates  map[int32][]string
}

var _ error = (*DuplicateTagNumberError)(nil)

func (d DuplicateTagNumberError) Error() string {
	if d.Duplicates == nil {
		return ""
	}

	ses := make([]string, 0, len(d.Duplicates))
	for k, v := range d.Duplicates {
		ses = append(ses, fmt.Sprintf("fields %s share the same field number %d", strings.Join(v, " "), k))
	}

	return strings.Join(ses, "\n")
}

func IsDuplicateTagNumberError(err error) *DuplicateTagNumberError {
	r := &DuplicateTagNumberError{}
	if errors.As(err, r) {
		return r
	}

	return nil
}

func (s *Struct) PrettyComments() []string {
	return prettyComments(s)
}

func (m *Struct) ValidTagNumbers() error {
	numbertofields := make(map[int32][]string)

	for _, f := range m.Fields {
		numbertofields[f.TagNum] = append(numbertofields[f.TagNum], f.Name)
	}

	toremove := make([]int32, 0, len(numbertofields))
	for k, v := range numbertofields {
		if len(v) <= 1 {
			toremove = append(toremove, k)
		}
	}

	for _, k := range toremove {
		delete(numbertofields, k)
	}

	if len(numbertofields) > 0 {
		return &DuplicateTagNumberError{
			MessageName: m.Name,
			Duplicates:  numbertofields,
		}
	}

	return nil
}

func (m *Struct) CleanPluralName() string {
	return getPluralName(m)
}

func (m *Struct) FirstField() *Field {
	if len(m.Fields) == 0 {
		panic(fmt.Errorf("%s has no fields", m.Name))
	}

	return m.Fields[0]
}
