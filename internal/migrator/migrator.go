package migrator

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
	migrationsDir string
	dbURL         string
}

// NewMigrator создает новый экземпляр Migrator.
// migrationsDir — путь к директории с миграциями (например, "./migrations").
// dbURL — строка подключения к БД.
func NewMigrator(migrationsDir, dbURL string) *Migrator {
	return &Migrator{
		migrationsDir: migrationsDir,
		dbURL:         dbURL,
	}
}

// MustApplyMigrations применяет все доступные миграции к базе данных.
// Завершает выполнение программы (panic), если происходит ошибка, кроме случая,
// когда нет новых миграций (ErrNoChange).
func (m *Migrator) MustApplyMigrations() {
	mig, err := migrate.New("file://"+m.migrationsDir, m.dbURL)
	if err != nil {
		panic(err)
	}

	if err = mig.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		panic(err)
	}
}
