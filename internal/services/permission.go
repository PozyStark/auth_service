package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository/interface"
	"context"
)

type PermissionService struct {
	repository repository.IPermissionRepository
}

func NewPermissionService(repository repository.IPermissionRepository) (*PermissionService) {
	return &PermissionService{
		repository: repository,
	}
}

func (s PermissionService) Save(model *models.Permission) (error) {
	return s.repository.Save(model)
}

func (s PermissionService) SaveByMap(valueMap map[string]any) (error) {
	return s.repository.SaveByMap(valueMap)
}

func (s PermissionService) UpdateById(id uint, values map[string]any) (error) {
	return s.repository.UpdateById(id, values)
}

func (s PermissionService) UpdateByModel(model *models.Permission) (error) {
	return s.repository.UpdateByModel(model)
}

func (s PermissionService) GetById(ctx context.Context, id uint) (models.Permission, error) {
	return s.repository.GetById(ctx, id)
}

func (s PermissionService) SafeDeleteById(id uint) (error) {
	return s.repository.SafeDeleteById(id)
}

func (s PermissionService) DeleteById(id uint) (error) {
	return s.repository.DeleteById(id)
}