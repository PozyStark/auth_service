package services

import (
	"auth-service/internal/models"
	repository "auth-service/internal/repository/interface"
	"context"
)

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s UserService) Save(model *models.User) error {
	return s.repository.Save(model)
}

func (s UserService) SaveByMap(valueMap map[string]any) error {
	return s.repository.SaveByMap(valueMap)
}

func (s UserService) UpdateById(id string, values map[string]any) error {
	return s.repository.UpdateById(id, values)
}

func (s UserService) UpdateByModel(model *models.User) error {
	return s.repository.UpdateByModel(model)
}

func (s UserService) GetUserByUserId(ctx context.Context, userId string) (models.User, error) {
	return s.repository.GetUserByUserId(ctx, userId)
}

func (s UserService) GetUserByUsername(ctx context.Context, userName string) (models.User, error) {
	return s.repository.GetUserByUsername(ctx, userName)
}

func (s UserService) GetUserByUsernameAndPassword(ctx context.Context, userName string, password string) (models.User, error) {
	return s.repository.GetUserByUsernameAndPassword(ctx, userName, password)
}

func (s UserService) SafeDeleteById(userId string) error {
	return s.repository.SafeDeleteById(userId)
}

func (s UserService) DeleteById(id string) error {
	return s.repository.DeleteById(id)
}
