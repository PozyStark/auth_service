package repository

import (
	"auth-service/internal/models"
	"context"
)

type IGroupPermissionRepository interface {
	Save(model *models.GroupPermission) (error)
	SaveByMap(valuesMap map[string]any) (error)
	UpdateById(id uint, valuesMap map[string]any) (error)
	UpdateByModel(model *models.GroupPermission) (error)
	GetById(ctx context.Context, id uint) (models.GroupPermission, error)
	GetGroupPermissionsByGroupId(ctx context.Context, id uint) ([]models.Permission, error)
	GetGroupPermissionsByGroupName(ctx context.Context, groupName string) (models.Permission, error)
	SafeDeleteById(id uint) (error)
	DeleteById(id uint) (error)
}