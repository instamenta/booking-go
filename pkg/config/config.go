package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache       bool
	TemplateCached map[string]*template.Template
	InfoLog        *log.Logger
	InProduction   bool
	Session        *scs.SessionManager
}
