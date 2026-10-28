package main

import (
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, filename string, data *templateData) {

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

	app.tp.Render(w, filename, app.defaultTemplateData(data, r))
}

func (app *application) defaultTemplateData(data *templateData, r *http.Request) *templateData {
	if data == nil {
		data = &templateData{}
	}

	data.Flash = app.session.PopString(r, "flash")

	return data
}
