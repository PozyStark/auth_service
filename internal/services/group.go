package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repository/interface"
	"context"
)

type GroupService struct {
	repository repository.IGroupRepository
}

func NewGroupService(repository repository.IGroupRepository) (*GroupService) {
	return &GroupService{
		repository: repository,
	}
}

func (s GroupService) Save(model *models.Group) (error) {
	return s.repository.Save(model)
}

func (s GroupService) SaveByMap(valueMap map[string]any) (error) {
	return s.repository.SaveByMap(valueMap)
}

func (s GroupService) UpdateById(id uint, values map[string]any) (error) {
	return s.repository.UpdateById(id, values)
}

func (s GroupService) UpdateByModel(model *models.Group) (error) {
	return s.repository.UpdateByModel(model)
}

func (s GroupService) GetById(ctx context.Context, id uint) (models.Group, error) {
	return s.repository.GetById(ctx, id)
}

func (s GroupService) SafeDeleteById(id uint) (error) {
	return s.repository.SafeDeleteById(id)
}

func (s GroupService) DeleteById(id uint) (error) {
	return s.repository.DeleteById(id)
}