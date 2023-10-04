package protorowdf_test

import (
	_ "embed"
	"testing"

	"github.com/fardream/protorowdf"
)

//go:embed testdata/input.yml
var gocopyinput []byte

//go:embed testdata/copy.go
var gocopyexpected string

func TestGoProtoFile_CopyDataCode(t *testing.T) {
	g := &protorowdf.GoProtoFile{
		ManualGoPackage: "_package",
	}

	var err error
	g.ProtoFile, err = protorowdf.ParseYAML(gocopyinput)
	if err != nil {
		t.Fatal(err)
	}

	def, err := g.CopyDataCode()
	if err != nil {
		t.Fatal(err)
	}

	want := gocopyexpected

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
