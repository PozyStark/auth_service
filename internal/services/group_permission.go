package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository/interface"
	"context"
)

type GroupPermissionService struct {
	repository repository.IGroupPermissionRepository
}

func NewGroupPermissionService(repository repository.IGroupPermissionRepository) (*GroupPermissionService) {
	return &GroupPermissionService{
		repository: repository,
	}
}

func (s GroupPermissionService) Save(model *models.GroupPermission) (error) {
	return s.repository.Save(model)
}

func (s GroupPermissionService) SaveByMap(valueMap map[string]any) (error) {
	return s.repository.SaveByMap(valueMap)
}

func (s GroupPermissionService) UpdateById(id uint, values map[string]any) (error) {
	return s.repository.UpdateById(id, values)
}

func (s GroupPermissionService) UpdateByModel(model *models.GroupPermission) (error) {
	return s.repository.UpdateByModel(model)
}

func (s GroupPermissionService) GetById(ctx context.Context, id uint) (models.GroupPermission, error) {
	return s.repository.GetById(ctx, id)
}

func (s GroupPermissionService) SafeDeleteById(id uint) (error) {
	return s.repository.SafeDeleteById(id)
}

func (s GroupPermissionService) DeleteById(id uint) (error) {
	return s.repository.DeleteById(id)
}