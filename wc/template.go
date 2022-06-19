package wc

import (
	"fmt"
	"html/template"
	"io/fs"
	"path"
	"strings"
)

type TemplateFetcher struct {
	fs.FS
}

type TemplateType string

func (t TemplateType) Suffix() string {
	return fmt.Sprintf(".%s.gohtml", t)
}

var (
	TemplateTypePage      TemplateType = "page"
	TemplateTypeComponent TemplateType = "component"
)

func NewTemplateFetcher(fs fs.FS) *TemplateFetcher {
	return &TemplateFetcher{
		fs,
	}
}

func (f *TemplateFetcher) get(fsPath, suffix string) (*TemplateRef, error) {
	file, err := fs.ReadFile(f, fsPath)
	if err != nil {
		return nil, err
	}
	name := path.Base(fsPath)
	name = strings.TrimSuffix(name, suffix)
	t, err := template.New(name).Parse(string(file))
	if err != nil {
		return nil, err
	}
	ref := &TemplateRef{
		Name:     name,
		Path:     fsPath,
		Template: t,
	}
	return ref, nil
}

func (f *TemplateFetcher) getPageTemplate(fsPath string) (*TemplateRef, error) {
	t, err := f.get(fsPath, TemplateTypePage.Suffix())
	if err != nil {
		return nil, fmt.Errorf("getting %s: %w", TemplateTypePage, err)
	}
	return t, nil
}

func (f *TemplateFetcher) getComponent(fsPath string) (*TemplateRef, error) {
	t, err := f.get(fsPath, TemplateTypeComponent.Suffix())
	if err != nil {
		return nil, fmt.Errorf("getting %s: %w", TemplateTypeComponent, err)
	}
	return t, nil
}

func NewTemplateRegistry(prefix string, fetcher *TemplateFetcher) (*Registry, error) {
	registry := &Registry{
		prefix:  prefix,
		fetcher: fetcher,
		lookup:  make(map[string]*TemplateRef),
	}
	err := fs.WalkDir(fetcher, ".", func(fullPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if path.Ext(fullPath) == ".gohtml" {
			ref, err := fetcher.getComponent(fullPath)
			if err != nil {
				return fmt.Errorf("getting TemplateRef at %s: %w", fullPath, err)
			}
			registry.Add(ref)
		}
		// TODO: implement page type or condense into just the get method
		return nil
	})
	if err != nil {

	}
	return registry, nil
}

type Registry struct {
	prefix  string
	fetcher *TemplateFetcher
	lookup  map[string]*TemplateRef
}

// TODO: new can probably add them...
func (t *Registry) Add(ref *TemplateRef) {
	t.lookup[ref.Name] = ref
}

// TODO: TemplateRef probably needs to be able to make a
//       specific Component instance and a component should
//       be able to Render() since it has state

// TemplateRef tracks the tree of child TemplateRefs
type TemplateRef struct {
	Name     string
	Path     string
	Template *template.Template
}
