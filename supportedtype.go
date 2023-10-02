package protorowdf

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goccy/go-yaml"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ProtoTypeString returns the proto type as a string
func (t SupportedType) ProtoTypeString() (string, error) {
	switch t {
	case SupportedType_Bool:
		return "bool", nil
	case SupportedType_Bytes:
		return "bytes", nil
	case SupportedType_Double:
		return "double", nil
	case SupportedType_Fixed32:
		return "fixed32", nil
	case SupportedType_Fixed64:
		return "fixed64", nil
	case SupportedType_Float:
		return "float", nil
	case SupportedType_Int32:
		return "int32", nil
	case SupportedType_Int64:
		return "int64", nil
	case SupportedType_Sfixed32:
		return "sfixed32", nil
	case SupportedType_Sfixed64:
		return "sfixed64", nil
	case SupportedType_Sint32:
		return "sint32", nil
	case SupportedType_Sint64:
		return "sint64", nil
	case SupportedType_String:
		return "string", nil
	case SupportedType_Uint32:
		return "uint32", nil
	case SupportedType_Uint64:
		return "uint64", nil
	default:
		return "", fmt.Errorf("unknown type %d", t)
	}
}

func (t SupportedType) MustProtoTypeString() string {
	return getOrPanic(t.ProtoTypeString())
}

var _ yaml.BytesMarshaler = SupportedType_Unknown

func (t SupportedType) domarshal(marshaler func(v any) ([]byte, error)) ([]byte, error) {
	s, err := t.ProtoTypeString()
	if err != nil {
		return nil, err
	}

	return marshaler(s)
}

func (t SupportedType) MarshalYAML() ([]byte, error) {
	return t.domarshal(yaml.Marshal)
}

var (
	_ json.Unmarshaler      = (*SupportedType)(nil)
	_ yaml.BytesUnmarshaler = (*SupportedType)(nil)
)

var titleCaser = cases.Title(language.AmericanEnglish)

func (t *SupportedType) dounmarshal(b []byte, unmarshaler func([]byte, any) error) error {
	var i int32
	err := unmarshaler(b, &i)
	if err == nil {
		_, valid := SupportedType_name[i]
		if !valid {
			return fmt.Errorf("%d/%s is not valid SupportedType", i, b)
		}

		*t = SupportedType(i)
		return nil
	}

	var s string
	err = unmarshaler(b, &s)
	if err != nil {
		return fmt.Errorf("failed to parse the input %s as string: %w", b, err)
	}
	s = titleCaser.String(strings.ToLower(s))

	v, valid := SupportedType_value[s]
	if !valid {
		return fmt.Errorf("failed to parse %s as a supported type", b)
	}

	*t = SupportedType(v)

	return nil
}

func (t *SupportedType) UnmarshalJSON(b []byte) error {
	return t.dounmarshal(b, json.Unmarshal)
}

func (t *SupportedType) UnmarshalYAML(b []byte) error {
	return t.dounmarshal(b, yaml.Unmarshal)
}
