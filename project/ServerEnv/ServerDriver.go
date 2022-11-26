package ServerEnv

import (
	"LinksShortner/project/LinkEnv"
	"fmt"
	"net/http"
)

func RunServer() {

	http.HandleFunc("/Hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello")
	})

	http.HandleFunc("/NVSL/", func(writer http.ResponseWriter, request *http.Request) {
		original := LinkEnv.SearchInDB(request.RequestURI)
		fmt.Println(original)
		http.Redirect(writer, request, original, http.StatusSeeOther)
	})

	http.ListenAndServe(":8080", nil)
}
