package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	// Can use it globally
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProd        bool
	Session       *scs.SessionManager
}
