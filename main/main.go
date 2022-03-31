package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./pages")))
	http.ListenAndServe(":5000", nil)
}
