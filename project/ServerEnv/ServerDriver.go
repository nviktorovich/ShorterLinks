package ServerEnv

import (
	"LinksShortner/project/Configuration"
	"LinksShortner/project/LinkEnv"
	"fmt"
	"log"
	"net/http"
	"os"
)

func RunServer() {

	http.HandleFunc(Configuration.MainPage, func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			page, err := GetStringFromFile(Configuration.TitlePath)
			if err != nil {
				log.Fatal(err)
			}
			page = fmt.Sprintf(page, "http://"+Configuration.Address+Configuration.ShortPage)
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

	http.HandleFunc(Configuration.ShortPage, func(writer http.ResponseWriter, request *http.Request) {
		page, err := GetStringFromFile(Configuration.ShortPath)
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
				page = fmt.Sprintf(page, Configuration.Address+linkObj.Short)
				if _, err := fmt.Fprintln(writer, page); err != nil { // здесь работаем с параметром %s
					log.Fatal(err)
				}
			} else {
				linkObj.Short = LinkEnv.GenerateShort()
				LinkEnv.WriteRowToDB(linkObj)
				LinkEnv.GetRowFromDB(linkObj)
				// создаем новый объект типа Counter, и делаем запись в БД, значение счетчика равно 0
				counterObj := LinkEnv.NewCounter(linkObj.Id)
				LinkEnv.WriteRowToDB(counterObj)
				page = fmt.Sprintf(page, Configuration.Address+linkObj.Short)

				if _, err := fmt.Fprintln(writer, page); err != nil {
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

	http.HandleFunc(Configuration.RedirectPage, func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			linkObj := LinkEnv.NewLink("")
			linkObj.Short = request.RequestURI
			res := LinkEnv.CheckRowInDB(linkObj)
			countObj := LinkEnv.NewCounter(0)
			if res {
				LinkEnv.GetRowFromDB(linkObj)
				countObj.FkLinkId = linkObj.Id

				if LinkEnv.CheckRowInDB(countObj) {
					LinkEnv.GetRowFromDB(countObj)
					countObj.CntIncrement()

				} else {
					LinkEnv.WriteRowToDB(countObj)
					LinkEnv.GetRowFromDB(countObj)
					countObj.CntIncrement()
				}

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

	if err := http.ListenAndServe(Configuration.Address, nil); err != nil {
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
