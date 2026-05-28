package repository

import (
	"auth-service/internal/models"
	"context"
)

type IRegistryToken interface {
	Save(model *models.RegistryToken) error
	SaveByMap(values map[string]any) error
	UpdateById(tokenId string, values map[string]any) error
	UpdateByModel(model *models.RegistryToken) error
	UpdateByUserId(userId string, values map[string]any) error
	GetByTokenId(ctx context.Context, tokenId string) (models.RegistryToken, error)
	GetActiveByUserId(ctx context.Context, userId string) ([]models.RegistryToken, error)
	SafeDeleteByTokenId(tokenId string) error
	DeleteByTokenId(tokenId string) error
}
