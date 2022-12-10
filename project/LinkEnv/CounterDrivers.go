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

	row := DBEnv.DoQuery(fmt.Sprintf("SELECT count(id) FROM counters WHERE %s = %d", "fk_link_id", c.FkLinkId))
	for row.Next() {
		if err := row.Scan(&check); err != nil {
			log.Println(err)
		}
	}

	return check != 0
}

// WriteRow - метод для объекта Counter, создает запись в таблице counters,
// обязательный аргумент - значение ключа-ссылки на таблицу link (FkLinkId)
func (c *Counter) WriteRow() {
	DBEnv.DoExec(fmt.Sprintf("INSERT INTO counters (fk_link_id, counter) VALUES(%d, %d)", c.FkLinkId, c.Counter))
}

// GetRow - метод для объекта Counter, возвращает значения всех полей объекта,
// запрос производиться по значению fk_link_id
func (c *Counter) GetRow() {

	row := DBEnv.DoQuery(fmt.Sprintf("SELECT * FROM counters WHERE fk_link_id = %d LIMIT(1)", c.FkLinkId))

	for row.Next() {
		err := row.Scan(&c.Id, &c.FkLinkId, &c.Counter)
		if err != nil {
			log.Println(err)
		}
	}
}

// CntIncrement - метод для объекта Counter, увеличивает на 1 значение счетчика.
func (c *Counter) CntIncrement() {
	c.Counter += 1
	DBEnv.DoExec(fmt.Sprintf("UPDATE counters SET counter = %d WHERE id = %d", c.Counter, c.Id))
}
