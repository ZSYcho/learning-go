package main

import (
	"fmt"
	"net/http"
)

var htmlContent = `
<!DOCTYPE html>
<html>
<head><title>%s</title></head>
<body>
	%s
</body>
</html>
`

// we have three handlers

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s %s", r.Method, r.URL)
	if r.Method == http.MethodGet {
		// we want to process a form here
	}
	homeContent := fmt.Sprintf(htmlContent, "Home", "<h1>Hello, welcome to the homepage</h1>")
	_, _ = w.Write([]byte(homeContent))
}

func about(w http.ResponseWriter, r *http.Request) {
	aboutContent := `
<h2>About</h2>
<div>We are a small web shop doing great things!</div>`

	aboutContent = fmt.Sprintf(htmlContent, "About us", aboutContent) // parse the template
	_, _ = w.Write([]byte(aboutContent))
}

func contact(w http.ResponseWriter, r *http.Request) {
	contactContent := `
<h2>Contact</h2>
<div>send as an email on akukopd@test.com</div>`

	contactContent = fmt.Sprintf(htmlContent, "Contact", contactContent)
	_, _ = w.Write([]byte(contactContent))
}
