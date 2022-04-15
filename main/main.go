package main

import (
	"forum"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./pages")))
	http.HandleFunc("/login", forum.Login)
	http.HandleFunc("/register", forum.Register)
	http.ListenAndServe(":5000", nil)
}
