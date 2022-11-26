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
			page, err := ReadHTML("project/HTMLForms/Title.html")
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
		page, err := ReadHTML("project/HTMLForms/Short.html")
		if err != nil {
			log.Fatal(err)
		}
		switch request.Method {
		case "POST":
			originLink := request.FormValue("origin_link")

			short := LinkEnv.GenerateShort()
			link := LinkEnv.NewLink(originLink)
			link.Short = short
			link.WriteToBD()
			if _, err := fmt.Fprintf(writer, page, short); err != nil {
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

		fmt.Fprintf(writer, page, "yo yo yo")
	})

	http.HandleFunc("/NVSL/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			original := LinkEnv.SearchInDB(request.RequestURI)
			if original == "" {
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
			} else {
				http.Redirect(writer, request, original, http.StatusSeeOther)
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

func ReadHTML(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
