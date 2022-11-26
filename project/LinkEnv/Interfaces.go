package LinkEnv

// DBCommunicator интерфейс который содержит в себе общие методы для работы различных структур с базами данных
// предназначен для упрощения взаимодействия, чтобы не вникать в суть реализации методов
type DBCommunicator interface {
	WriteToBD()
}

func WriteRowToDB(i DBCommunicator) {
	i.WriteToBD()
}
