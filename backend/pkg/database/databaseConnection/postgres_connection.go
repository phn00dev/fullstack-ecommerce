package databaseConnection

import (
	"eCommerce/pkg/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnect struct {
	config *config.Config
}

func NewPostgresConnection(config *config.Config) *PostgresConnect {
	return &PostgresConnect{config: config}
}

func (p *PostgresConnect) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		p.config.DbConfig.DbHost,
		p.config.DbConfig.DbUser,
		p.config.DbConfig.DbPassword,
		p.config.DbConfig.DbName,
		p.config.DbConfig.DbPort,
		p.config.DbConfig.DbSslMode,
		p.config.DbConfig.DbTimeZone,
	)
	log.Println("database password")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
