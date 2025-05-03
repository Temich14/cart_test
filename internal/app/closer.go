package app

type DBCloser interface {
	CloseDB() error
}
