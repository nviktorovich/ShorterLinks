package main

import (
	"net/http"

	serverfunctions "github.com/nviktorovich/ShorterLinks/serverFunctions"
)

func main() {
	titlePageHandler := &serverfunctions.LongLinkHandler{}
	http.Handle("/title", titlePageHandler)
	http.ListenAndServe(":8080", nil)
}
