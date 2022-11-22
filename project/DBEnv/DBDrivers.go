package DBEnv

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const SETTINGS = "user=postgres password=2005760 dbname=postgres sslmode=disable"

// Base структура, предназначенная для работы с конкретным объектом базой данных
// DataBase - объект базы данных
// Err - ошибка
type Base struct {
	DataBase *sql.DB
	Err      error
}

func NewBase(settings string) *Base {
	db, err := sql.Open("postgres", settings)
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

//db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)",
//        "Apple", 72000)
