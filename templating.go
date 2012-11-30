package webshell

import (
	"bytes"
	"html/template"
	"path/filepath"
)

func CompileTemplate(filename string) (tpl *template.Template, err error) {
	template_name := filepath.Base(filename)
	tpl = template.New(template_name)
	if err != nil {
		return
	}
	tpl, err = tpl.ParseFiles(filename)
	return
}

func ServeTemplate(tpl *template.Template, in interface{}) (out []byte, err error) {
	buffer := new(bytes.Buffer)
	err = tpl.Execute(buffer, in)
	if err == nil {
		out = buffer.Bytes()
	}
	return
}

// ServeTemplate serves the template specified in filename, executed with the
// data specified in 'in', and returns a byte slice and error.
func ServeTemplateFile(filename string, in interface{}) (out []byte, err error) {
	tpl, err := CompileTemplate(filename)
	if err != nil {
		return
	}
	out, err = ServeTemplate(tpl, in)
	return
}

// Check a template file for errors.
func CheckTemplate(filename string) (err error) {
	template_name := filepath.Base(filename)
	t := template.New(template_name)
	if err != nil {
		return
	}
	t, err = t.ParseFiles(filename)
	return
}
