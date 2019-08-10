package main

import (
	"os"
	"log"
	"text/template"
)

var t *template.Template

func init () {
	t = template.Must(template.ParseGlob("./templates/*"))
}

func main () {
	err := t.Execute(os.Stdout, 42)
	if err != nil {
		log.Fatalln(err)
	}
}