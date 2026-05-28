package repository

import (
	db "auth-service/internal/db/gorm"
	"auth-service/internal/models"
	"context"
	"time"
)

type RegistryTokenRepository struct {
	model        *models.RegistryToken
	dbConnection *db.GormDbConnection
}

func NewGormRegistryTokenRepository(dbConnection *db.GormDbConnection) RegistryTokenRepository {
	return RegistryTokenRepository{dbConnection: dbConnection}
}

func (r RegistryTokenRepository) Save(model *models.RegistryToken) error {
	session := r.dbConnection.Session()
	err := session.Model(r.model).Create(model)
	return err.Error
}

func (r RegistryTokenRepository) SaveByMap(valuesMap map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(valuesMap)
	return err.Error
}

func (r RegistryTokenRepository) UpdateById(id string, values map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(values)
	return err.Error
}

func (r RegistryTokenRepository) UpdateByModel(model *models.RegistryToken) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", model.ID).Updates(model)
	return err.Error
}

func (r RegistryTokenRepository) UpdateByUserId(userId string, values map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("user_id=?", userId).Updates(values)
	return err.Error
}

func (r RegistryTokenRepository) GetByTokenId(
	ctx context.Context,
	tokenId string,
) (models.RegistryToken, error) {
	var result models.RegistryToken
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Debug().Model(r.model).Where(&models.RegistryToken{ID: tokenId}).First(&result)
	return result, err.Error
}

func (r RegistryTokenRepository) GetActiveByUserId(
	ctx context.Context,
	userId string,
) ([]models.RegistryToken, error) {
	var result []models.RegistryToken
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Debug().Model(r.model).Where(
		"user_id = ? AND expire > ?", userId, time.Now().UTC(),
	).Find(&result)
	return result, err.Error
}

func (r RegistryTokenRepository) SafeDeleteByTokenId(tokenId string) error {
	session := r.dbConnection.Session()
	safeDeleteFields := make(map[string]any)
	safeDeleteFields["IsActive"] = false
	safeDeleteFields["DeletedAt"] = time.Now()
	err := session.Debug().Model(r.model).Where("id=?", tokenId).Updates(safeDeleteFields)
	return err.Error
}

func (r RegistryTokenRepository) DeleteByTokenId(tokenId string) error {
	session := r.dbConnection.Session()
	err := session.Debug().Where("id=?", tokenId).Unscoped().Delete(r.model)
	return err.Error
}
