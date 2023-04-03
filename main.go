package main

import (
	"html/template"
	"log"
	"os"
)

type Focus struct {
	Focus string
}

func main() {
	focus := &Focus{Focus: "Hello"}
	t, err := template.ParseFiles("template/focus.tmpl")
	if err != nil {
		log.Fatalf("Bad Parse: %v", err)
	}

	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("create file: %v", err)
	}

	err = t.Execute(f, focus)
	if err != nil {
		log.Fatalf("Bad execute: %v", err)
	}
}
