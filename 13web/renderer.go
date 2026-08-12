package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type TemplateRenderer struct {
	cache       map[string]*template.Template
	mutex       sync.RWMutex
	dev         bool
	templateDir string
}

func NewTemplateRenderer(templateDir string, isDev bool) *TemplateRenderer {
	return &TemplateRenderer{
		templateDir: templateDir,
		cache:       make(map[string]*template.Template),
		dev:         isDev,
	}
}

func (t *TemplateRenderer) Render(w http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := t.getTemplate(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplateRenderer) getTemplate(templateName string) (*template.Template, error) {
	if !t.dev { // if we are not in the dev mode
		// we got the lock, then we get template from the cache
		t.mutex.RLock()
		if tmpl, ok := t.cache[templateName]; ok {
			// if we find the template in the cache, we unlock and return it
			t.mutex.RUnlock()
			return tmpl, nil
		}
		t.mutex.RUnlock()
	}

	// if it is not in cache, parse it to cache and return the template
	tmpl, err := t.parseTemplate(templateName)
	if err != nil {
		return nil, err
	}

	if !t.dev {
		t.mutex.Lock()
		t.cache[templateName] = tmpl
		t.mutex.Unlock()
	}

	return tmpl, nil
}

func (t *TemplateRenderer) parseTemplate(templateName string) (*template.Template, error) {
	templatePath := filepath.Join(t.templateDir, templateName)

	files := []string{templatePath}

	layoutPath := filepath.Join(t.templateDir, "layouts/*.html")
	layout, err := filepath.Glob(layoutPath)
	if err == nil {
		files = append(files, layout...)
	}

	partialPath := filepath.Join(t.templateDir, "partials/*.html")
	partials, err := filepath.Glob(partialPath)
	if err == nil {
		files = append(files, partials...)
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
