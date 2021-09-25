package services

import (
	"fmt"
	"log"
	"os"

	"github.com/leeliwei930/go_commerce/ent"
)

type DBConfig struct {
	driver   string
	hostname string
	database string
	port     string
	username string
	password string
}

func (config *DBConfig) GetConnectionUrl() string {
	connectionURLFormat := "%s:%s@tcp(%s:%s)/%s?parseTime=True"
	return fmt.Sprintf(connectionURLFormat,
		config.username,
		config.password,
		config.hostname,
		config.port,
		config.database,
	)
}

func GetDefaultDBClient() *ent.Client {
	dbConfig := &DBConfig{
		driver:   os.Getenv("DATABASE_DRIVER"),
		hostname: os.Getenv("DATABASE_HOST"),
		database: os.Getenv("DATABASE_NAME"),
		port:     os.Getenv("DATABASE_PORT"),
		username: os.Getenv("DATABASE_USERNAME"),
		password: os.Getenv("DATABASE_PASSWORD"),
	}

	dbClient, err := ent.Open(dbConfig.driver, dbConfig.GetConnectionUrl())
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	return dbClient
}
