package repository

import (
	db "auth-service/internal/db/gorm"
	"auth-service/internal/models"
	"context"
	"time"
)

type PermissionRepository struct {
	model        *models.Permission
	dbConnection *db.GormDbConnection
}

func NewGormPermissionRepository(dbConnection *db.GormDbConnection) PermissionRepository {
	return PermissionRepository{dbConnection: dbConnection}
}

func (r PermissionRepository) Save(model *models.Permission) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(model)
	return err.Error
}

func (r PermissionRepository) SaveByMap(valuesMap map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(valuesMap)
	return err.Error
}

func (r PermissionRepository) UpdateById(id uint, values map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(values)
	return err.Error
}

func (r PermissionRepository) UpdateByModel(model *models.Permission) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", model.ID).Updates(model)
	return err.Error
}

func (r PermissionRepository) GetById(
	ctx context.Context,
	id uint,
) (models.Permission, error) {
	var result models.Permission
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Debug().Model(r.model).First(&result, "id=?", id)
	return result, err.Error
}

func (r PermissionRepository) SafeDeleteById(id uint) error {
	session := r.dbConnection.Session()
	safeDeleteFields := make(map[string]any)
	safeDeleteFields["IsActive"] = false
	safeDeleteFields["DeletedAt"] = time.Now()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(safeDeleteFields)
	return err.Error
}

func (r PermissionRepository) DeleteById(id uint) error {
	session := r.dbConnection.Session()
	err := session.Debug().Where("id=?", id).Unscoped().Delete(r.model)
	return err.Error
}
