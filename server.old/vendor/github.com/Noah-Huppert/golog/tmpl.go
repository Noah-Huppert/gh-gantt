package golog

import (
	"bytes"
	"text/template"
)

// MustMkTmpl creates a Go template. Panics if an error occurs.
func MustMkTmpl(tmpl string) *template.Template {
	return template.Must(template.New("log format").Parse(tmpl))
}

// MustExecTmpl executes a Go template and returns the result. Panics if an
// error occurs.
func MustExecTmpl(t *template.Template, data interface{}) string {
	var out bytes.Buffer

	err := t.Execute(&out, &data)
	if err != nil {
		panic(err)
	}

	return out.String()
}
