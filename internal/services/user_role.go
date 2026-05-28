package services

import (
	"auth-service/internal/models"
	repository "auth-service/internal/repository/interface"
	"context"
	// "time"
)

type UserRoleService struct {
	repository repository.IUserRoleRepository
}

func NewUserRoleService(repository repository.IUserRoleRepository) *UserRoleService {
	return &UserRoleService{
		repository: repository,
	}
}

func (s UserRoleService) Save(model *models.UserRole) error {
	return s.repository.Save(model)
}

func (s UserRoleService) SaveByMap(valueMap map[string]any) error {
	return s.repository.SaveByMap(valueMap)
}

func (s UserRoleService) UpdateById(id uint, values map[string]any) error {
	return s.repository.UpdateById(id, values)
}

func (s UserRoleService) UpdateByModel(model *models.UserRole) error {
	return s.repository.UpdateByModel(model)
}

func (s UserRoleService) GetById(ctx context.Context, id uint) (models.UserRole, error) {
	return s.repository.GetById(ctx, id)
}

func (s UserRoleService) GetUserRolesByUserId(ctx context.Context, userId string) ([]models.Role, error) {
	return s.repository.GetUserRolesByUserId(ctx, userId)
}

func (s UserRoleService) GetUserRolesByUserIdAsStringSlice(ctx context.Context, userId string) ([]string, error) {
	// time.Sleep(time.Millisecond*200)
	return s.repository.GetUserRolesByUserIdAsStringSlice(ctx, userId)
}

func (s UserRoleService) SafeDeleteById(id uint) error {
	return s.repository.SafeDeleteById(id)
}

func (s UserRoleService) DeleteById(id uint) error {
	return s.repository.DeleteById(id)
}
