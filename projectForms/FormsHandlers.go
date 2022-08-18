package projectforms

import (
	"errors"
	"os"
)

const formPath = "projectforms"

// ReadForm - функция предназначена для чтения шаблонов HTML в папке
// projectforms. Возвращает []byte - содержимое шаблона или ошибку,
// если шаблон не удалось прочитать
func ReadForm(formName string) ([]byte, error) {
	data, err := os.ReadFile(formPath + "/" + formName)
	if err != nil {
		err = errors.New("ошибка чтения файла. Не удалось открыть файл: " + formName)
		return data, err
	}

	return data, nil
}
