package services

import (
	"auth-service/internal/models"
	repository "auth-service/internal/repository/interface"
	"context"
)

type RolePermissionService struct {
	repository repository.IRolePermissionRepository
}

func NewRolePermissionService(repository repository.IRolePermissionRepository) *RolePermissionService {
	return &RolePermissionService{
		repository: repository,
	}
}

func (s RolePermissionService) Save(model *models.RolePermission) error {
	return s.repository.Save(model)
}

func (s RolePermissionService) SaveByMap(valueMap map[string]any) error {
	return s.repository.SaveByMap(valueMap)
}

func (s RolePermissionService) UpdateById(id uint, values map[string]any) error {
	return s.repository.UpdateById(id, values)
}

func (s RolePermissionService) UpdateByModel(model *models.RolePermission) error {
	return s.repository.UpdateByModel(model)
}

func (s RolePermissionService) GetById(ctx context.Context, id uint) (models.RolePermission, error) {
	return s.repository.GetById(ctx, id)
}

func (s RolePermissionService) SafeDeleteById(id uint) error {
	return s.repository.SafeDeleteById(id)
}

func (s RolePermissionService) DeleteById(id uint) error {
	return s.repository.DeleteById(id)
}
