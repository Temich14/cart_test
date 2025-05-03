package config

import (
	"github.com/joho/godotenv"
	"log"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
