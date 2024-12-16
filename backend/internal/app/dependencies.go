package app

import (
	"eCommerce/pkg/config"
	"eCommerce/pkg/database/databaseConnection"
	"eCommerce/pkg/httpClient"
	"gorm.io/gorm"
	"net/http"
)

type Dependencies struct {
	DB         *gorm.DB
	HttpClient *http.Client
	Config     *config.Config
}

func GetDependencies() (*Dependencies, error) {
	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	// postgres connection
	newPostgresDB := databaseConnection.NewPostgresConnection(getConfig)
	// redis connection
	getDB, err := newPostgresDB.Connect()
	if err != nil {
		return nil, err
	}

	// http client connection
	clientHttp := httpClient.NewHttpConnect()

	return &Dependencies{
		DB:         getDB,
		HttpClient: clientHttp,
		Config:     getConfig,
	}, nil
}
