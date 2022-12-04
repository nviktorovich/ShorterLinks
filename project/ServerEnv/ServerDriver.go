package ServerEnv

import (
	"LinksShortner/project/LinkEnv"
	"fmt"
	"log"
	"net/http"
	"os"
)

func RunServer() {

	http.HandleFunc("/OriginalToShort", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			page, err := GetStringFromFile("project/HTMLForms/Title.html")
			if err != nil {
				log.Fatal(err)
			}
			if _, err := fmt.Fprintln(writer, page); err != nil {
				log.Fatal(err)
			}
		default:
			err := fmt.Errorf("Метод %s не поддерживается\npage not found %d",
				request.Method,
				http.StatusNotFound,
			)
			if _, err := fmt.Fprintln(writer, err); err != nil {
				log.Fatal(err)
			}

		}

	})

	http.HandleFunc("/Short", func(writer http.ResponseWriter, request *http.Request) {
		page, err := GetStringFromFile("project/HTMLForms/Short.html")
		if err != nil {
			log.Fatal(err)
		}
		switch request.Method {
		case "POST":
			originalLink := request.FormValue("origin_link")
			if originalLink == "" {
				if _, err := fmt.Fprintln(writer, "отсутствует содержимое поля 'original_link', короткая ссылка не может быть создана"); err != nil {
					log.Fatal(err)
				}
			}

			linkObj := LinkEnv.NewLink(originalLink)
			res := LinkEnv.CheckRowInDB(linkObj)
			if res {
				LinkEnv.GetRowFromDB(linkObj)
				if _, err := fmt.Fprintf(writer, page, linkObj.Short); err != nil {
					log.Fatal(err)
				}
			} else {
				linkObj.Short = LinkEnv.GenerateShort()
				LinkEnv.WriteRowToDB(linkObj)
				LinkEnv.GetRowFromDB(linkObj)
				if _, err := fmt.Fprintf(writer, page, linkObj.Short); err != nil {
					log.Fatal(err)
				}
			}

		default:
			err := fmt.Errorf("Метод %s не поддерживается\npage not found %d",
				request.Method,
				http.StatusNotFound,
			)
			if _, err := fmt.Fprintln(writer, err); err != nil {
				log.Fatal(err)
			}
		}
	})

	http.HandleFunc("/NVSL/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			linkObj := LinkEnv.NewLink("")
			linkObj.Short = request.RequestURI
			res := LinkEnv.CheckRowInDB(linkObj)
			if res {
				LinkEnv.GetRowFromDB(linkObj)
				http.Redirect(writer, request, linkObj.Original, http.StatusSeeOther)
				// добавить функционал наращивания значения счетчика
			} else {
				err := fmt.Errorf(
					"указанной короткой ссылке '%s' не соответствует ни одна оригинальная ссылка из БД\npage not found %d",
					request.RequestURI,
					http.StatusNotFound,
				)
				if _, err := fmt.Fprintln(writer, err); err != nil {
					log.Fatal(err)
				}
				if err != nil {
					return
				}
			}

		default:
			err := fmt.Errorf("Метод %s не поддерживается\npage not found %d",
				request.Method,
				http.StatusNotFound,
			)
			if _, err := fmt.Fprintln(writer, err); err != nil {
				log.Fatal(err)
			}
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// GetStringFromFile ReadHTML функция, предназначена для чтения файла (по
// указаному пути - path)и возврата его содержимого в виде строки
func GetStringFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
