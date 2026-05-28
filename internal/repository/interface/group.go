package repository

import (
	"auth-service/internal/models"
	"context"
)

type IGroupRepository interface {
	Save(model *models.Group) (error)
	SaveByMap(valuesMap map[string]any) (error)
	UpdateById(id uint, valuesMap map[string]any) (error)
	UpdateByModel(model *models.Group) (error)
	GetById(ctx context.Context, id uint) (models.Group, error)
	SafeDeleteById(id uint) (error)
	DeleteById(id uint) (error)
}