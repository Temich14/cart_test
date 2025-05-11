package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type ServerConfig struct {
	Port   string
	Secret string
}
type DBConfig struct {
	Conn string
}
type AppConfig struct {
	Env          string
	ServerConfig *ServerConfig
	DBConfig     *DBConfig
}

func MustLoad() *AppConfig {
	_ = godotenv.Load()
	srvConfig := ServerConfig{
		Port:   os.Getenv("PORT"),
		Secret: os.Getenv("SECRET"),
	}
	dbConfig := DBConfig{
		Conn: buildDBConnectionString(),
	}
	APPConfig := AppConfig{
		Env:          os.Getenv("ENV"),
		ServerConfig: &srvConfig,
		DBConfig:     &dbConfig,
	}
	return &APPConfig
}
func buildDBConnectionString() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
}
