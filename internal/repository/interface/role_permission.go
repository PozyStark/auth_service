package repository

import (
	"auth-service/internal/models"
	"context"
)

type IRolePermissionRepository interface {
	Save(model *models.RolePermission) (error)
	SaveByMap(valuesMap map[string]any) (error)
	UpdateById(id uint, valuesMap map[string]any) (error)
	UpdateByModel(model *models.RolePermission) (error)
	GetById(ctx context.Context, id uint) (models.RolePermission, error)
	GetRolePermissionsByRoleId(ctx context.Context, id uint) (models.Permission, error)
	GetRolePermissionsByRoleName(ctx context.Context, roleName string) ([]models.Permission, error)
	SafeDeleteById(id uint) error
	DeleteById(id uint) error
}
