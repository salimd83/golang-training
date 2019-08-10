package main

import (
	"net/http"
	"html/template"
	"log"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

func root(w http.ResponseWriter, r *http.Request) {
	err := t.Execute(w, "Welcome to the homepage")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func d(w http.ResponseWriter, r *http.Request) {
	err := t.Execute(w, "Woof Woof Woof")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	err := t.Execute(w, "Salim Dirani")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main () {
	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}