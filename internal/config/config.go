package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// ServerConfig содержит конфигурацию HTTP-сервера.
type ServerConfig struct {
	Port   string // Порт сервера.
	Secret string // Секретный ключ JWT.
}

// DBConfig содержит параметры подключения к базе данных.
type DBConfig struct {
	Conn string // Строка подключения к базе данных PostgreSQL.
}

// AppConfig агрегирует все конфигурации приложения.
type AppConfig struct {
	Env          string        // Текущее окружение (DEV, PROD).
	ServerConfig *ServerConfig // Конфигурация сервера.
	DBConfig     *DBConfig     // Конфигурация базы данных.
}

// MustLoad загружает переменные окружения и возвращает объект конфигурации приложения.
// Завершает работу программы с ошибкой, если не удалось загрузить конфиг.
func MustLoad() *AppConfig {
	_ = godotenv.Load() // Загружает переменные окружения из .env файла если он есть.

	srvConfig := ServerConfig{
		Port:   os.Getenv("PORT"),
		Secret: os.Getenv("SECRET"),
	}

	dbConfig := DBConfig{
		Conn: buildDBConnectionString(),
	}

	appConfig := AppConfig{
		Env:          os.Getenv("ENV"),
		ServerConfig: &srvConfig,
		DBConfig:     &dbConfig,
	}

	return &appConfig
}

// buildDBConnectionString собирает строку подключения к PostgreSQL
// из переменных окружения: DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME.
func buildDBConnectionString() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
}
