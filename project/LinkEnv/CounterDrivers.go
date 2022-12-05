package LinkEnv

import (
	"LinksShortner/project/DBEnv"
	"fmt"
	"log"
)

// Counter специальная структура, позволяющая обрабатывать счетчики из таблицы counters
// Id идентификатор счетчика в БД
// FkLinkId foreign key на Id ссылки в таблице links
// Counter счетчик, значение которого должно увеличиваться каждый раз, когда
// пользователь переходит по короткой ссылке
type Counter struct {
	Id       int
	FkLinkId int
	Counter  int
}

// NewCounter функция, возвращает объект типа Counter, для создания необходим ключ FkLinkId
func NewCounter(FkLinkId int) *Counter {
	return &Counter{Id: 0, FkLinkId: FkLinkId, Counter: 0}
}

// CheckRow - метод для объекта Counter, проверяет, существует ли счетчик для данной ссылки
// в БД, если существует, возвращает true, иначе - false
func (c *Counter) CheckRow() bool {
	var check int

	DB := DBEnv.NewBase(DBEnv.SETTINGS)
	qr := fmt.Sprintf("SELECT count(id) FROM counters WHERE %s = '%s'", "fk_link_id", c.FkLinkId)
	row, err := DB.DataBase.Query(qr)

	if err != nil {
		log.Println(err)
	}

	for row.Next() {
		if err = row.Scan(&check); err != nil {
			log.Println(err)
		}
	}
	defer DB.DataBase.Close()

	return check != 0
}

// WriteRow - метод для объекта Counter, создает запись в таблице counters,
// обязательный аргумент - значение ключа-ссылки на таблицу link (FkLinkId)
func (c *Counter) WriteRow() {
	DB := DBEnv.NewBase(DBEnv.SETTINGS)
	defer DB.DataBase.Close()
	qr := fmt.Sprintf("INSERT INTO counters (fk_link_id, counter) VALUES(%d, %d)", c.FkLinkId, c.Counter)
	DB.DataBase.Exec(qr)
}

// GetRow - метод для объекта Counter, возвращает значения всех полей объекта,
// запрос производиться по значению fk_link_id
func (c *Counter) GetRow() {
	DB := DBEnv.NewBase(DBEnv.SETTINGS)
	defer DB.DataBase.Close()
	qr := fmt.Sprintf("SELECT * FROM counter WHERE fk_link_id = %d LIMIT(1)", c.FkLinkId)
	row, err := DB.DataBase.Query(qr)
	if err != nil {
		log.Println(err)
	}
	for row.Next() {
		row.Scan(&c.Id, &c.FkLinkId, &c.Counter)
	}
}

// CntIncrement - метод для объекта Counter, увеличивает на 1 значение счетчика.
// Запись в БД не производиться, необходимо использовать соответствующий метод отдельно.
func (c *Counter) CntIncrement() {
	c.Counter += 1
}
