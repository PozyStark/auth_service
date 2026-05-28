package db

import (
	"fmt"
	"gorm.io/gorm"
)


func MigrationTransaction(transaction *gorm.DB, model any) error {
	return transaction.Debug().AutoMigrate(model)
}

/*
	Автоматическая миграция проводится в рамках транзакции.
	Если миграция проходит не успешно это приводит к вызову panic и завершению работы сервиса
*/
func MustMakeMigrations(dbConnection *GormDbConnection, sessionConfig *gorm.Session, models...any) {
	session := dbConnection.GetConnection()
	err := session.Transaction(
		func(tx *gorm.DB) error {
			for i:=0; i<len(models);i++ {
				if err := MigrationTransaction(tx, models[i]); err != nil {
					return err
				}
			}
			return nil
		},
	)

	if err != nil {
		fmt.Printf("Ошибка при проведении миграции: %v", err)
		panic(1)
	}

}