package config

import (
	"github.com/joho/godotenv"
	"os"
)

type ServerConfig struct {
	Port string
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
		Port: os.Getenv("PORT"),
	}
	dbConfig := DBConfig{
		Conn: os.Getenv("DB_CONNECTION"),
	}
	APPConfig := AppConfig{
		Env:          os.Getenv("ENV"),
		ServerConfig: &srvConfig,
		DBConfig:     &dbConfig,
	}
	return &APPConfig
}
