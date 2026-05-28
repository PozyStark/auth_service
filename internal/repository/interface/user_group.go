package repository

import (
	"auth-service/internal/models"
	"context"
)

type IUserGroupRepository interface {
	Save(model *models.UserGroup) (error)
	SaveByMap(valuesMap map[string]any) (error)
	UpdateById(id uint, valuesMap map[string]any) (error)
	UpdateByModel(model *models.UserGroup) (error)
	GetById(ctx context.Context, id uint) (models.UserGroup, error)
	GetUserGroupsByUserId(ctx context.Context, userId string) ([]models.Group, error)
	GetUserGroupsByUserIdAsStringSlice(ctx context.Context, userId string) ([]string, error)
	GetUserGroupsByUsername(ctx context.Context, userName string) ([]models.Group, error)
	SafeDeleteById(id uint) (error)
	DeleteById(id uint) (error)
}