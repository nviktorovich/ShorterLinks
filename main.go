package main

import (
	"net/http"

	serverfunctions "github.com/nviktorovich/ShorterLinks/serverFunctions"
)

func main() {
	titlePageHandler := &serverfunctions.WellcomeHandler{}
	shortLinkGenerateHandler := &serverfunctions.LinkGenerateHandler{}
	http.Handle("/title", titlePageHandler)
	http.Handle("/send", shortLinkGenerateHandler)
	http.ListenAndServe(":8080", nil)
}
