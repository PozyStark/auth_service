package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository/interface"
	"context"
)

type UserPermissionService struct {
	repository repository.IUserPermissionRepository
}

func NewUserPermissionService(repository repository.IUserPermissionRepository) (*UserPermissionService) {
	return &UserPermissionService{
		repository: repository,
	}
}

func (s UserPermissionService) Save(model *models.UserPermission) (error) {
	return s.repository.Save(model)
}

func (s UserPermissionService) SaveByMap(valueMap map[string]any) (error) {
	return s.repository.SaveByMap(valueMap)
}

func (s UserPermissionService) UpdateById(id uint, values map[string]any) (error) {
	return s.repository.UpdateById(id, values)
}

func (s UserPermissionService) UpdateByModel(model *models.UserPermission) (error) {
	return s.repository.UpdateByModel(model)
}

func (s UserPermissionService) GetById(ctx context.Context, id uint) (models.UserPermission, error) {
	return s.repository.GetById(ctx, id)
}

func (s UserPermissionService) GetUserPermissionsByUserId(ctx context.Context, userId string) ([]models.Permission, error) {
	return s.repository.GetUserPermissionsByUserId(ctx, userId)
}

func (s UserPermissionService) GetUserPermissionsByUsername(ctx context.Context, userName string) ([]models.Permission, error) {
	return s.repository.GetUserPermissionsByUsername(ctx, userName)
}

func (s UserPermissionService) GetAllUserPermissionsByUserId(ctx context.Context, userId string) ([]models.Permission, error) {
	return s.repository.GetAllUserPermissionsByUserId(ctx, userId)
}

func (s UserPermissionService) GetAllUserPermissionsByUserIdAsStringSlice(ctx context.Context, userId string) ([]string, error) {
	return s.repository.GetAllUserPermissionsByUserIdAsStringSlice(ctx, userId)
}

func (s UserPermissionService) GetAllUserPermissionsByUsername(ctx context.Context, userName string) ([]models.Permission, error) {
	return s.repository.GetAllUserPermissionsByUsername(ctx, userName)
}

func (s UserPermissionService) SafeDeleteById(id uint) (error) {
	return s.repository.SafeDeleteById(id)
}

func (s UserPermissionService) DeleteById(id uint) (error) {
	return s.repository.DeleteById(id)
}