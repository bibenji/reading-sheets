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
	data     map[string]interface{}
}

func (t *templateHandler) SetData(data map[string]interface{}) {
	t.data = data
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(
			filepath.Join("templates", "layout.html"),
			filepath.Join("templates", t.filename)))
	})

	rData := map[string]interface{}{
		"Host": r.Host,
	}

	for k, v := range t.data {
		if _, ok := t.data[k]; ok {
			rData[k] = v
		}
	}

	// you can do something like that too
	// if authCookie, err := r.Cookie("auth"); err == nil {
	// 	data["UserData"] = objx.MustFromBase64(authCookie.Value)
	// }

	t.templ.Execute(w, rData)
}
