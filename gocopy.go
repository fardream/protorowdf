package protorowdf

import (
	"bytes"
	_ "embed"
	"io"
	"strings"
	"text/template"

	"mvdan.cc/gofumpt/format"
)

//go:embed gocopy.tmpl
var gocopytemplatestring string

var gocopytemplate *template.Template

type GoProtoFile struct {
	*ProtoFile

	ManualGoPackage string
}

func init() {
	gocopytemplate = getOrPanic(template.New("gocopy-template").Parse(gocopytemplatestring))
}

func (g *GoProtoFile) GoPackageName() string {
	if g.ManualGoPackage != "" {
		return g.ManualGoPackage
	}

	v := strings.Split(g.PackageName, ".")

	return v[len(v)-1]
}

func (g *GoProtoFile) WriteCopyDataCode(w io.Writer) error {
	return gocopytemplate.ExecuteTemplate(w, "copy-for-file", g)
}

func (g *GoProtoFile) CopyDataCode() (string, error) {
	var s bytes.Buffer

	if err := g.WriteCopyDataCode(&s); err != nil {
		return "", err
	}

	formatted, err := format.Source(s.Bytes(), format.Options{LangVersion: "v1.21", ExtraRules: true})
	if err != nil {
		return "", nil
	}

	return string(formatted), nil
}
