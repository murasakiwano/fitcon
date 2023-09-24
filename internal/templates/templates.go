package templates

import (
	"embed"
	"html/template"
	"io"

	"github.com/murasakiwano/fitcon_templates/internal/fitconner"
)

//go:embed templates/*
var fitconTemplates embed.FS

func ParseFitConnerTable(w io.Writer, player fitconner.FitConner) error {
	templ, err := template.ParseFS(fitconTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "fitcon.gohtml", player); err != nil {
		return err
	}

	return nil
}
