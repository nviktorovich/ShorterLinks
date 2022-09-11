package programs

import (
	"errors"
	"os"
)

// WriteInfo - функция для записи info в файл-path
func WriteInfo(filePath, info string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		err = errors.New("не удалось открыть файл:" + filePath)
		return err
	}

	_, err = file.WriteString(info + "\n")
	if err != nil {
		err = errors.New("не удалось записать информацию в файл:" + info)
		return err
	}

	defer file.Close()

	return nil
}
