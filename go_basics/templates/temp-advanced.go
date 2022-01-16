package main

import (
	"os"
	"text/template"
)

type employee struct {
	Name        string
	Age         int
	Location    string
	IsPermanent bool
	Skills      []string
}

func main() {

	e1 := employee{
		Name:        "vishnu",
		Age:         29,
		Location:    "India",
		IsPermanent: true,
		Skills:      []string{"golang", "python", "shell"},
	}

	t1 := template.Must(template.New("foo.txt").Parse(`Hello {{ .Name }},
You are {{ .Age }} years old.
{{ if .IsPermanent -}}
  You are a permanent employee.
{{ else }}
  You are a contract employee.
{{ end -}}
Below are your skills:
{{- range $_, $val := .Skills }}
- {{ $val -}}
{{ end -}}
{{/* https://pkg.go.dev/text/template#hdr-Functions */}}
{{ if lt .Age 30 -}}
You are young!
{{ end -}}
`))
	t1.Execute(os.Stdout, e1)
}
