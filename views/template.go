package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	HtmlTpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("error parsing FS template: %w", err)
	}
	return Template{
		HtmlTpl: tpl,
	}, nil
}

func Parse(filePath string) (Template, error) {
	tpl, err := template.ParseFiles(filePath)
	if err != nil {
		return Template{}, fmt.Errorf("error parsing template: %w", err)
	}
	return Template{
		HtmlTpl: tpl,
	}, nil
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=urf-8")
	err := t.HtmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Error execute template: %v", err)
		http.Error(w, "There was an error execute the template", http.StatusInternalServerError)
		return
	}
}
