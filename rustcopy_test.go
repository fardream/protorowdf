package protorowdf_test

import (
	_ "embed"
	"testing"

	"github.com/fardream/protorowdf"
)

//go:embed testdata/input.yml
var rustcopyinput []byte

//go:embed testdata/rusty/src/lib.rs
var rustcopyexpected string

func TestRustProfileFile_CopyDataCode(t *testing.T) {
	r := &protorowdf.RustProtoFile{
		RustCopyConfig: &protorowdf.RustCopyConfig{
			GenInclude: true,
		},
	}

	var err error

	r.ProtoFile, err = protorowdf.ParseYAML(rustcopyinput)

	if err != nil {
		t.Fatal(err)
	}

	def, err := r.CopyDataCode()
	if err != nil {
		t.Fatal(err)
	}

	want := rustcopyexpected

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
