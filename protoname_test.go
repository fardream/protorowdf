package protorowdf_test

import (
	"testing"

	"github.com/fardream/protorowdf"
)

func TestGoProtoName(t *testing.T) {
	testcases := [][2]string{
		{"_birth_year", "XBirthYear"},
		{"birth_year_2", "BirthYear_2"},
		{"birth_Year_2", "Birth_Year_2"},
	}

	for _, c := range testcases {
		goname := protorowdf.GoProtoName(c[0])
		if c[1] != goname {
			t.Errorf("want: %s, got: %s for %s", c[1], goname, c[0])
		}
	}
}
