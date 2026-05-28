package repository

import (
	"auth-service/internal/models"
	"context"
)

type IPermissionRepository interface {
	Save(model *models.Permission) (error)
	SaveByMap(valuesMap map[string]any) (error)
	UpdateById(id uint, valuesMap map[string]any) (error)
	UpdateByModel(model *models.Permission) (error)
	GetById(ctx context.Context, id uint) (models.Permission, error)
	SafeDeleteById(id uint) (error)
	DeleteById(id uint) (error)
}