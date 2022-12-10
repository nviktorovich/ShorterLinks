package LinkEnv

import (
	"LinksShortner/project/Configuration"
	"LinksShortner/project/DBEnv"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// Link специальная структура, позволяющая обрабатывать ссылки
// Id идентификатор ссылки в БД
// Original исходная ссылка
// Short сокращенная ссылка
type Link struct {
	Id       int
	Original string
	Short    string
}

// NewLink функция, возвращает объект типа Link, для создания необходима оригинальная ссылка
func NewLink(original string) *Link {
	return &Link{Id: 0, Original: original, Short: ""}
}

// GenerateShort функция, предназначена для генерации короткой строки, строка
// состоит из префикса "NVSL" и рандомной цифровой части данная реализация
// является критически ошибочной, поскольку в теории возможна генерация
// нескольких одинаковых коротких цифровых составляющих, но для учебного проекта,
// подойдет.
func GenerateShort() string {
	source := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(source)
	return Configuration.Prefix + strconv.Itoa(randomNumber.Intn(999999999))
}

// CheckRow - метод для объекта Link, проверяет, существует ли переданная ссылка
// в БД, если существует, возвращает true, иначе - false
func (l *Link) CheckRow() bool {
	var check bool
	// в объектке Link могут содержаться обе ссылки, и оригинальная и короткая.
	// Поэтому необходимо проверять оба варианта
	if l.Original != "" {
		return DBCheckQuery("original", l.Original)
	} else if l.Short != "" {
		return DBCheckQuery("short", l.Short)
	}
	return check
}

// GetRow - метод, который возвращает заполненную структуру Link
func (l *Link) GetRow() {

	if l.Original != "" {

		row := DBEnv.DoQuery(fmt.Sprintf("SELECT * FROM links WHERE original = '%s' LIMIT(1)", l.Original))
		for row.Next() {
			row.Scan(&l.Id, &l.Original, &l.Short)
		}
	} else if l.Short != "" {
		row := DBEnv.DoQuery(fmt.Sprintf("SELECT * FROM links WHERE short = '%s' LIMIT(1)", l.Short))

		for row.Next() {
			row.Scan(&l.Id, &l.Original, &l.Short)
		}
	}
}

// WriteRow - метод, который создает запись в БД, для создания записи необходимо
// сгенерировать короткую ссылку
func (l *Link) WriteRow() {

	DBEnv.DoExec(fmt.Sprintf("INSERT INTO links (original, short) VALUES('%s', '%s')", l.Original, l.Short))
}

// DBCheckQuery на вход подается два параметра - строки, название поля и значение
// поля. На выходе булево значение. Если в БД существует поле с таким значением,
// то возвращается true, если нет - false
func DBCheckQuery(fieldName, fieldValue string) bool {
	var i int

	row := DBEnv.DoQuery(fmt.Sprintf("SELECT count(id) FROM links WHERE %s = '%s'", fieldName, fieldValue))
	for row.Next() {
		err := row.Scan(&i)

		if err != nil {
			log.Println(err)
		}
	}

	return i != 0

}
