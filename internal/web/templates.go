package web

import (
	"embed"
	"html/template"
	"time"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

// Templates returns parsed html templates for Gin.
func Templates() *template.Template {
	funcs := template.FuncMap{
		"date": func(t time.Time) string { return t.Format("2006-01-02") },
	}
	tmpl := template.Must(template.New("").Funcs(funcs).ParseFS(templateFS, "templates/*.tmpl"))
	return tmpl
}

