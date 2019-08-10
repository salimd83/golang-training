package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})

	cookie, err := r.Cookie("visits")
	if err == http.ErrNoCookie {
		// http.SetCookie(w, &http.Cookie{
		// 	Name: "visits",
		// 	Value: "0",
		// })
		cookie = &http.Cookie{
			Name:  "visits",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Number of visits:", count)

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YUOR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
