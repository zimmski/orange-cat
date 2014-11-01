package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Template(filepath string, port int) func(*http.ResponseWriter) {
	templateStr := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<title>%[1]s</title>
</head>
<body>
<h1>Hello, world! port: %[2]d</h1>
</body>`, filepath, port)

	template, err := template.New("template").Parse(templateStr)

	if err != nil {
		panic(err)
	}

	return func(w *http.ResponseWriter) {
		err := template.Execute(*w, nil)
		if err != nil {
			panic(err)
		}
	}
}
