package app

// DBCloser интерфейс для закрытия соединения с базой данных
type DBCloser interface {
	CloseDB() error
}
