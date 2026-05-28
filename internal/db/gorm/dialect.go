package db

import (
	"auth-service/internal/config"
	db "auth-service/internal/db/standart"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
	Пример диалектора для MySql
*/
func GetMySqlDialect(dbConfig config.DbConfig) {

}

func GetPostgressDialect(dbConfig config.DbConfig) gorm.Dialector {
	return postgres.Open(db.GetDataSource(db.POSTGRES_DRIVER, dbConfig))
}