package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository/interface"
	"context"
)

type RoleService struct {
	repository repository.IRoleRepository
}

func NewRoleService(repository repository.IRoleRepository) (*RoleService) {
	return &RoleService{
		repository: repository,
	}
}

func (s RoleService) Save(model *models.Role) (error) {
	return s.repository.Save(model)
}

func (s RoleService) SaveByMap(valueMap map[string]any) (error) {
	return s.repository.SaveByMap(valueMap)
}

func (s RoleService) UpdateById(id uint, values map[string]any) (error) {
	return s.repository.UpdateById(id, values)
}

func (s RoleService) UpdateByModel(model *models.Role) (error) {
	return s.repository.UpdateByModel(model)
}

func (s RoleService) GetById(ctx context.Context, id uint) (models.Role, error) {
	return s.repository.GetById(ctx, id)
}

func (s RoleService) SafeDeleteById(id uint) (error) {
	return s.repository.SafeDeleteById(id)
}

func (s RoleService) DeleteById(id uint) (error) {
	return s.repository.DeleteById(id)
}