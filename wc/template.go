package wc

import "html/template"

type TemplateFetcher interface {
	GetPage(name string) template.Template
	GetTemplate(name string) template.Template
}

type TemplateManager struct {
	fetcher TemplateFetcher
}

func NewTemplateManager(fetcher TemplateFetcher) TemplateManager {
	return TemplateManager{fetcher}
}
