package repository

import (
	"auth-service/internal/models"
	"context"
)

type IUserRepository interface {
	Save(model *models.User) (error)
	SaveByMap(valuesMap map[string]any) (error)
	UpdateById(id string, valuesMap map[string]any) (error)
	UpdateByModel(model *models.User) (error)
	GetUserByUserId(ctx context.Context, userId string) (models.User, error)
	GetUserByUsername(ctx context.Context, userName string) (models.User, error)
	GetUserByUsernameAndPassword(ctx context.Context, userName string, password string) (models.User, error)
	SafeDeleteById(id string) (error)
	DeleteById(id string) (error)
}