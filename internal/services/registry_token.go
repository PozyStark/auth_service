package services

import (
	"auth-service/internal/models"
	repository "auth-service/internal/repository/interface"
	"context"
)

type RegistryTokenService struct {
	repository repository.IRegistryToken
}

func NewRegistryTokenService(repository repository.IRegistryToken) *RegistryTokenService {
	return &RegistryTokenService{
		repository: repository,
	}
}

func (s RegistryTokenService) Save(model *models.RegistryToken) error {
	return s.repository.Save(model)
}

func (s RegistryTokenService) SaveByMap(values map[string]any) error {
	return s.repository.SaveByMap(values)
}

func (s RegistryTokenService) UpdateById(tokenId string, values map[string]any) error {
	return s.repository.UpdateById(tokenId, values)
}

func (s RegistryTokenService) UpdateByModel(model *models.RegistryToken) error {
	return s.repository.UpdateByModel(model)
}

func (s RegistryTokenService) UpdateByUserId(userId string, values map[string]any) error {
	return s.repository.UpdateByUserId(userId, values)
}

func (s RegistryTokenService) GetByTokenId(ctx context.Context, tokenId string) (models.RegistryToken, error) {
	return s.repository.GetByTokenId(ctx, tokenId)
}

func (s RegistryTokenService) GetActiveByUserId(ctx context.Context, userId string) ([]models.RegistryToken, error) {
	return s.repository.GetActiveByUserId(ctx, userId)
}

func (s RegistryTokenService) SafeDeleteByTokenId(tokenId string) error {
	return s.repository.SafeDeleteByTokenId(tokenId)
}

func (s RegistryTokenService) DeleteByTokenId(tokenId string) error {
	return s.repository.DeleteByTokenId(tokenId)
}
