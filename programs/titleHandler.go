package programs

import (
	"fmt"
	"net/http"
	"os"
)

const (
	logFile = "programs/net.log"
)

// TitleHandler простая структура для обработки титульной страницы
type TitleHandler struct{}

func (h *TitleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("forms/TitleForm.html")
	if err != nil {
		http.Error(w, "ошибка обработки html формы", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, string(data))
	case http.MethodPost:
		obj := NewLinkForm(r.FormValue("long_link"))
		fmt.Fprint(w, obj.OriginalLink)
	}

}

// LinkForm структура, содержащая информацию об обрабатываемой ссылке
// OriginalLink - оригинальная (исходная ссылка)
// ShortLink - сокращенная ссылка
// Counter - счетчик вызова короткой ссылки
type LinkForm struct {
	OriginalLink string
	ShortLink    string
	Counter      int
}

// NewLinkForm принримает на вход строку (ссылку) возвращает объект LinkForm
func NewLinkForm(link string) *LinkForm {
	return &LinkForm{OriginalLink: link}
}
