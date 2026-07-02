package main

import (
	"fmt"
	"os"
	"strings"
	"text/template" // good option for command line application, you can choose html/template for web app
)

// First you need prepare the data

type EmailData struct {
	RecipientName string
	SenderName    string
	Subject       string
	Body          string
	Items         []string
	UnreadCount   int
	//Sample interface{}
	Sample any
}

func main() {

	fmt.Println("--- Text template example ---")

	// then you need to customize your template
	emailTemplate := `
Subject: {{ .Subject }}

{{.Body}}

{{if .Items}}
	Related Items:
{{range .Items}}
	- {{.}}
{{end}}
{{end}}

{{if gt .UnreadCount 0}}
You have {{.UnderCount}} unreads.
{{else}}
You have no messages.
{{end}}


- Thanks
{{.SenderName}}
`
	// You build the template object
	tmpl, err := template.New("email-message").Parse(emailTemplate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// You pass the data value that you've created
	data := EmailData{
		RecipientName: "Alice",
		SenderName:    "Bob's Auto-Responder",
		Subject:       "Your Weekly Update",
		Body:          "Here is the update you requested. We hope you find it useful.",
		Items:         []string{"Report A", "Document B", "Summary C"},
		UnreadCount:   0,
	}

	// then you execute
	err = tmpl.Execute(os.Stdout, data) // put it on the console
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// another way to output
	var output strings.Builder

	err = tmpl.Execute(&output, data) // During the execution, choosing the buffer created before
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(strings.ToUpper(output.String()))
}
