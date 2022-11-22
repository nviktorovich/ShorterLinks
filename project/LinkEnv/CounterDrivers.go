package LinkEnv

// Counter специальная структура, позволяющая обрабатывать счетчики переходов
// id идентификатор счетчика в БД
// fk_link_id соответствие между счетчиком и ссылкой в Link
// counter счетчик переходов
type Counter struct {
	id         int
	fk_link_id int
	counter    int
}
