package LinkEnv

// DBCommunicator интерфейс который содержит в себе общие методы для работы
// различных структур с базами данных предназначен для упрощения взаимодействия,
// чтобы не вникать в суть реализации методов
type DBCommunicator interface {
	CheckRow() bool
	WriteRow()
	GetRow()
}

// CheckRowInDB функция для обощения методов, удовлетворяющих интерфейсу DBCommunicator, реализует метод CheckRow
func CheckRowInDB(i DBCommunicator) bool {
	return i.CheckRow()
}

// GetRowFromDB функция для обощения методов, удовлетворяющих интерфейсу DBCommunicator, реализует метод WriteRow
func GetRowFromDB(i DBCommunicator) {
	i.GetRow()
}

// WriteRowToDB функция для обобщения методов, удовлетворяющих интерфесу DBCommunicator, реализует метод WriteRow
func WriteRowToDB(i DBCommunicator) {
	i.WriteRow()
}
