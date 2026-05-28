package db

import (
	"auth-service/internal/config"
	"fmt"
)

const (
	POSTGRES_DRIVER = "postgres"
)

func GetDataSource(driverName string, dbConfig config.DbConfig) string {

	host := dbConfig.DbHost
	port := dbConfig.DbPort
	user := dbConfig.DbUser
	password := dbConfig.DbPassword
	dbname := dbConfig.DbName
	sslmode := dbConfig.SslMode

	var driver string

	switch driverName {
	case POSTGRES_DRIVER:
		driver = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode,
		)
	}
	return driver
}
