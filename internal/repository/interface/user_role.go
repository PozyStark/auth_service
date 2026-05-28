package repository

import (
	"auth-service/internal/models"
	"context"
)

type IUserRoleRepository interface {
	Save(model *models.UserRole) (error)
	SaveByMap(valuesMap map[string]any) (error)
	UpdateById(id uint, valuesMap map[string]any) (error)
	UpdateByModel(model *models.UserRole) (error)
	GetById(ctx context.Context, id uint) (models.UserRole, error)
	GetUserRolesByUserId(ctx context.Context, userId string) ([]models.Role, error)
	GetUserRolesByUserIdAsStringSlice(ctx context.Context, userId string) ([]string, error)
	GetUserRolesByUsername(ctx context.Context, userName string) ([]models.Role, error)
	SafeDeleteById(id uint) (error)
	DeleteById(id uint) (error)
}