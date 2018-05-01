package main

import (
	"net/http"
)

var mainHandler = http.FileServer(http.Dir("site"))

func main() {
	http.Handle("/", mainHandler)
	http.ListenAndServe(":3000", nil)
}
