package protorowdf

import (
	_ "embed"
	"text/template"
)

//go:embed proto.tmpl
var protoTemplateString string

var protoTemplate *template.Template

func init() {
	protoTemplate = getOrPanic(template.New("proto-template").Parse(protoTemplateString))
}
