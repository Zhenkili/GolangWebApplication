package config

import (
	"html/template"
	"log"
)

type Appconfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
}
