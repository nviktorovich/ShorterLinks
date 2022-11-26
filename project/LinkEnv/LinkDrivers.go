package LinkEnv

import (
	"LinksShortner/project/DBEnv"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	PREFIX = "http://localhost:8080/NVSL/"
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
	short := GenerateShort()
	return &Link{Id: 0, Original: original, Short: short}
}

// GenerateShort функция, предназначена для генерации короткой строки, строка
// состоит из префикса "NVSL" и рандомной цифровой части
func GenerateShort() string {
	source := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(source)
	return PREFIX + strconv.Itoa(randomNumber.Intn(999999999))
}

// WriteToBD метод струкруты Link, предназначен для создания записи в базе
// данных. Id записи не передается и создается, непосредственно во время записи в
// БД
func (l *Link) WriteToBD() {
	DB := DBEnv.NewBase(DBEnv.SETTINGS)

	_, err := DB.DataBase.Exec("insert into links (original,short) values ($1, $2)",
		l.Original,
		l.Short,
	)
	if err != nil {
		DB.Err = err
		log.Print(err)
	}
	defer DB.Close()
}

// SearchInDB принимает на вход строку (короткую ссылку short) и возвращает
// строку (исходную ссылку original). Поиск осуществляется по совпадению короткой
// ссылки с полем short в таблице link. Поиск ведется до первого совпадения, установлен параметр LIMIT 1.
// внутри функции создается объект типа Link - oneLink.
func SearchInDB(short string) (original string) {
	DB := DBEnv.NewBase(DBEnv.SETTINGS)
	rows, err := DB.DataBase.Query("SELECT * FROM links WHERE short = $1 LIMIT 1", short)

	if err != nil {
		DB.Err = err
		log.Print(err)
	}

	defer DB.Close()

	oneLink := &Link{}
	for rows.Next() {
		err = rows.Scan(&oneLink.Id, &oneLink.Original, &oneLink.Short)
		if err != nil {
			DB.Err = err
			log.Print(err)
			continue
		}
	}

	return oneLink.Original

}
