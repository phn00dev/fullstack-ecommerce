package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig     dbConfig     `json:"db_config"`
	RedisConnect redisConnect `json:"redis_connect"`
	HttpConfig   httpConfig   `json:"http_config"`
	FolderConfig folderConfig `json:"folder_config"`
}

type redisConnect struct {
	Username string `json:"username" env:"REDIS_USERNAME"`
	Password string `json:"password" env:"REDIS_PASSWORD"`
	Host     string `json:"host" env:"REDIS_HOST"`
	Port     string `json:"port" env:"REDIS_PORT"`
}

type dbConfig struct {
	DbHost     string `json:"db_host" env:"DB_HOST"`
	DbPort     string `json:"db_port" env:"DB_PORT"`
	DbUser     string `json:"db_user" env:"DB_USER"`
	DbPassword string `json:"db_password" env:"DB_PASSWORD"`
	DbName     string `json:"db_name" env:"DB_NAME"`
	DbSslMode  string `json:"db_sll_mode" env:"DB_SLL_MODE"`
	DbTimeZone string `json:"db_time_zone" env:"DB_TIME_ZONE"`
}
type httpConfig struct {
	HttpHost  string `json:"http_host" env:"HTTP_HOST"`
	HttpPort  string `json:"http_port" env:"HTTP_PORT"`
	AppName   string `json:"app_name" env:"APP_NAME"`
	AppHeader string `json:"app_header" env:"APP_HEADER"`
}

type folderConfig struct {
	PublicPath string `json:"public_path" env:"PUBLIC_PATH"`
	RootPath   string `json:"root_path" env:"ROOT_PATH"`
}

func GetConfig() (*Config, error) {
	if err := godotenv.Load("../.env"); err != nil {
		return nil, errors.New("error loading .env file")
	}
	dbCfg := dbConfig{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbSslMode:  os.Getenv("DB_SLL_MODE"),
		DbTimeZone: os.Getenv("DB_TIME_ZONE"),
	}

	httpCfg := httpConfig{
		HttpHost:  os.Getenv("HTTP_HOST"),
		HttpPort:  os.Getenv("HTTP_PORT"),
		AppName:   os.Getenv("APP_NAME"),
		AppHeader: os.Getenv("APP_HEADER"),
	}

	folderCfg := folderConfig{
		PublicPath: os.Getenv("PUBLIC_PATH"),
		RootPath:   os.Getenv("ROOT_PATH"),
	}

	redisCfg := redisConnect{
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
	}

	return &Config{
		DbConfig:     dbCfg,
		RedisConnect: redisCfg,
		HttpConfig:   httpCfg,
		FolderConfig: folderCfg,
	}, nil
}
