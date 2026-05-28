package repository

import (
	db "auth-service/internal/db/gorm"
	"auth-service/internal/models"
	"context"
	"time"
)

type UserRepository struct {
	model        *models.User
	dbConnection *db.GormDbConnection
}

func NewGormUserRepository(dbConnection *db.GormDbConnection) UserRepository {
	return UserRepository{dbConnection: dbConnection}
}

func (r UserRepository) Save(model *models.User) error {
	session := r.dbConnection.Session()
	err := session.Model(r.model).Create(model)
	return err.Error
}

func (r UserRepository) SaveByMap(valuesMap map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Model(r.model).Create(valuesMap)
	return err.Error
}

func (r UserRepository) UpdateById(id string, values map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(values)
	return err.Error
}

func (r UserRepository) UpdateByModel(model *models.User) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", model.ID).Updates(model)
	return err.Error
}

func (r UserRepository) GetUserByUserId(
	ctx context.Context,
	userId string,
) (models.User, error) {
	var result models.User
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Debug().Model(r.model).First(&result, "id=?", userId)
	return result, err.Error
}

func (r UserRepository) GetUserByUsername(
	ctx context.Context,
	userName string,
) (models.User, error) {
	var result models.User
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Debug().Model(r.model).First(&result, "username=?", userName)
	return result, err.Error
}

func (r UserRepository) GetUserByUsernameAndPassword(
	ctx context.Context,
	userName string,
	password string,
) (models.User, error) {
	var result models.User
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Model(r.model).Where(
		&models.User{Username: userName, Password: password},
	).First(&result)
	return result, err.Error
}

func (r UserRepository) SafeDeleteById(id string) error {
	session := r.dbConnection.Session()
	safeDeleteFields := make(map[string]any)
	safeDeleteFields["IsActive"] = false
	safeDeleteFields["DeletedAt"] = time.Now()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(safeDeleteFields)
	return err.Error
}

func (r UserRepository) DeleteById(id string) error {
	session := r.dbConnection.Session()
	err := session.Debug().Where("id=?", id).Unscoped().Delete(r.model)
	return err.Error
}
