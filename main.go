package main

import (
	"html/template"
	"log"
	"net/http"
)

type Focus struct {
	Focus string
}

func handle(w http.ResponseWriter, r *http.Request) {
	focus := &Focus{Focus: "Hello"}
	tmpl, err := template.ParseFiles("template/focus.tmpl")
	if err != nil {
		log.Fatalf("Bad Parse: %v", err)
	}

	tmpl.Execute(w, focus)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}
