package serverfunctions

import (
	"fmt"
	"net/http"

	projectforms "github.com/nviktorovich/ShorterLinks/projectForms"
)

type WellcomeHandler struct{}
type LinkGenerateHandler struct{}

func (h *WellcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := projectforms.ReadForm("WellcomeForm.html")
	if err != nil {
		http.Error(w, "ошибка. Не удалось прочитать необходимую форму", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, string(data))
}

func (h *LinkGenerateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "здесь должна быть короткая ссылка")
}
