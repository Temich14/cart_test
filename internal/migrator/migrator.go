package migrator

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

type Migrator struct {
	migrationsDir string
	dbURL         string
}

func NewMigrator(migrationsDir, dbURL string) *Migrator {
	return &Migrator{
		migrationsDir: migrationsDir,
		dbURL:         dbURL,
	}
}

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
