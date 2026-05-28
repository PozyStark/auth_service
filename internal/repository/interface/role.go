package repository

import (
	"auth-service/internal/models"
	"context"
)

type IRoleRepository interface {
	Save(model *models.Role) (error)
	SaveByMap(valuesMap map[string]any) (error)
	UpdateById(id uint, valuesMap map[string]any) (error)
	UpdateByModel(model *models.Role) (error)
	GetById(ctx context.Context, id uint) (models.Role, error)
	SafeDeleteById(id uint) (error)
	DeleteById(id uint) (error)
}