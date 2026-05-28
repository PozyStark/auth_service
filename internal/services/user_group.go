package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository/interface"
	"context"
)

type UserGroupService struct {
	repository repository.IUserGroupRepository
}

func NewUserGroupService(repository repository.IUserGroupRepository) (*UserGroupService) {
	return &UserGroupService{
		repository: repository,
	}
}

func (s UserGroupService) Save(model *models.UserGroup) (error) {
	return s.repository.Save(model)
}

func (s UserGroupService) SaveByMap(valueMap map[string]any) (error) {
	return s.repository.SaveByMap(valueMap)
}

func (s UserGroupService) UpdateById(id uint, values map[string]any) (error) {
	return s.repository.UpdateById(id, values)
}

func (s UserGroupService) UpdateByModel(model *models.UserGroup) (error) {
	return s.repository.UpdateByModel(model)
}

func (s UserGroupService) GetById(ctx context.Context, id uint) (models.UserGroup, error) {
	return s.repository.GetById(ctx, id)
}

func (s UserGroupService) GetUserGroupsByUserId(ctx context.Context, userId string) ([]models.Group, error) {
	return s.repository.GetUserGroupsByUserId(ctx, userId)
}

func (s UserGroupService) GetUserGroupsByUserIdAsStringSlice(ctx context.Context, userId string) ([]string, error) {
	return s.repository.GetUserGroupsByUserIdAsStringSlice(ctx, userId)
}

func (s UserGroupService) SafeDeleteById(id uint) (error) {
	return s.repository.SafeDeleteById(id)
}

func (s UserGroupService) DeleteById(id uint) (error) {
	return s.repository.DeleteById(id)
}