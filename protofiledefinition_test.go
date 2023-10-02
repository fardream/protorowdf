package protorowdf_test

import (
	_ "embed"
	"testing"

	"github.com/fardream/protorowdf"
)

//go:embed testdata/input.yml
var protofileinput []byte

//go:embed testdata/output.proto
var protofileexpected string

func TestProtoFile_ProtoFileDefinition(t *testing.T) {
	v, err := protorowdf.ParseYAML(parseyamlInput)
	if err != nil {
		t.Fatal(err)
	}

	def, err := v.ProtoFileDefinition()
	if err != nil {
		t.Fatal(err)
	}

	want := protofileexpected

	if def != want {
		l := min(len(def), len(want))
		for i := 0; i < l; i++ {
			if def[i] != want[i] {
				t.Logf("%d %s %s", i, string(def[i:]), string(want[i:]))

				break
			}
		}

		t.Errorf("want: \n%s\ngot: \n%s\n", want, def)
	}
}
