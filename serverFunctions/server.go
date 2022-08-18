package serverfunctions

import (
	"fmt"
	"net/http"

	projectforms "github.com/nviktorovich/ShorterLinks/projectForms"
)

type LongLinkHandler struct{}

func (h *LongLinkHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := projectforms.ReadForm("longLingForm.html")
	if err != nil {
		http.Error(w, "ошибка. Не удалось прочитать необходимую форму", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, string(data))
}
