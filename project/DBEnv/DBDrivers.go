package DBEnv

import (
	"LinksShortner/project/Configuration"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

//const (
//	SETTINGS    = "user=postgres password=2005760 dbname=postgres sslmode=disable"
//	DRIVER_NAME = "postgres"
//)

// Base структура, предназначенная для работы с конкретным объектом базой данных
// DataBase - объект базы данных
// Err - ошибка
type Base struct {
	DataBase *sql.DB
	Err      error
}

func NewBase() *Base {
	db, err := sql.Open(Configuration.DriverName, Configuration.DBInit)
	return &Base{
		DataBase: db,
		Err:      err,
	}
}

// Close метод, предназначенный для закрытия объекта базы данных
func (b *Base) Close() {
	err := b.DataBase.Close()
	b.Err = err
}

// DoQuery функция, предназначеная для того, чтобы получать из базы
// данных строки, удовлетворяющие запросу rqst
func DoQuery(rqst string) *sql.Rows {
	DB := NewBase()
	defer DB.Close()
	rows, err := DB.DataBase.Query(rqst)
	if err != nil {
		log.Println(err)
	}

	return rows
}

// DoExec функция, предназначенная для того, чтобы записывать в БД
// строку с переданными в аргументе значениями
func DoExec(rqst string) {
	DB := NewBase()
	defer DB.Close()

	_, err := DB.DataBase.Exec(rqst)
	if err != nil {
		log.Println(err)
	}
}
