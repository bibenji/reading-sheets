package main

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
	data     interface{}
}

func (t *templateHandler) SetData(data []interface{}) {
	t.data = data
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))

	})

	// t.templ.Execute(w, r)
	t.templ.Execute(w, t.data)
}