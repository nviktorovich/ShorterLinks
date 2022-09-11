package main

import (
	"net/http"

	"github.com/nviktorovich/ShorterLinks/programs"
)

func main() {
	hndlr := &programs.TitleHandler{}
	http.Handle("/hello", hndlr)
	http.ListenAndServe(":8080", nil)
}
