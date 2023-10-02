package protorowdf_test

import (
	_ "embed"
	"testing"

	"github.com/fardream/protorowdf"
)

//go:embed testdata/input.yml
var parseyamlInput []byte

func TestParseYAML(t *testing.T) {
	_, err := protorowdf.ParseYAML(parseyamlInput)
	if err != nil {
		t.Fatal(err)
	}
}
