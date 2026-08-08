package main

import (
	"net/http"
)

func (app *application) render(w http.ResponseWriter, filename string, data interface{}) {

	//fullPath := filepath.Join(app.templateDir, filename)
	//tmpl, err := template.ParseFiles(fullPath)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//err = tmpl.Execute(w, data)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	if app.tp == nil {
		http.Error(w, "template rendering engine is not set", http.StatusInternalServerError)
		return
	}
	app.tp.Render(w, filename, data)
}
