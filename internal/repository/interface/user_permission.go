package repository

import (
	"auth-service/internal/models"
	"context"
)

type IUserPermissionRepository interface {
	Save(model *models.UserPermission) (error)
	SaveByMap(valuesMap map[string]any) (error)
	UpdateById(id uint, valuesMap map[string]any) (error)
	UpdateByModel(model *models.UserPermission) (error)
	GetById(ctx context.Context, id uint) (models.UserPermission, error)
	GetUserPermissionsByUserId(ctx context.Context, userId string) ([]models.Permission, error)
	GetAllUserPermissionsByUserId(ctx context.Context, userId string) ([]models.Permission, error)
	GetAllUserPermissionsByUserIdAsStringSlice(ctx context.Context, userId string) ([]string, error)
	GetUserPermissionsByUsername(ctx context.Context, userName string) ([]models.Permission, error)
	GetAllUserPermissionsByUsername(ctx context.Context, userName string) ([]models.Permission, error)
	SafeDeleteById(id uint) (error)
	DeleteById(id uint) (error)
}