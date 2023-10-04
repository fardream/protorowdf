package protorowdf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"text/template"
)

//go:embed rustcopy.tmpl
var rustcopytemplatestring string

var rustcopytemplate *template.Template

type RustCopyConfig struct {
	GenInclude bool
	NoCargoOut bool
	FileName   string
}
type RustProtoFile struct {
	*ProtoFile

	*RustCopyConfig
}

func init() {
	rustcopytemplate = getOrPanic(template.New("rustcopy-template").Parse(rustcopytemplatestring))
}

func (r *RustProtoFile) WriteCopyDataCode(w io.Writer) error {
	return rustcopytemplate.ExecuteTemplate(w, "copy-for-file", r)
}

func (r *RustProtoFile) CopyDataCode() (string, error) {
	var s bytes.Buffer

	if err := r.WriteCopyDataCode(&s); err != nil {
		return "", err
	}

	return s.String(), nil
}

func (r *RustProtoFile) IncludeLine() string {
	if !r.GenInclude {
		return ""
	}

	filename := r.FileName
	if r.FileName == "" {
		filename = r.PackageName + ".rs"
	}

	if r.NoCargoOut {
		return fmt.Sprintf("include!(\"%s\");\n\n", filename)
	}

	return fmt.Sprintf("include!(concat!(env!(\"OUT_DIR\"), \"/%s\"));\n\n", filename)
}
