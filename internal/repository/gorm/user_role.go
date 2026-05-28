package repository

import (
	"auth-service/internal/db/gorm"
	"auth-service/internal/models"
	"context"
	"database/sql"
	"time"
)

type UserRoleRepository struct {
	model        *models.UserRole
	dbConnection *db.GormDbConnection
}

func NewGormUserRoleRepository(dbConnection *db.GormDbConnection) UserRoleRepository {
	return UserRoleRepository{dbConnection: dbConnection}
}

func (r UserRoleRepository) Save(model *models.UserRole) (error) {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(model)
	return err.Error
}

func (r UserRoleRepository) SaveByMap(valuesMap map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(valuesMap)
	return err.Error
}

func (r UserRoleRepository) UpdateById(id uint, values map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(values)
	return err.Error
}

func (r UserRoleRepository) UpdateByIdReturn(id uint, values map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(values)
	return err.Error
}

func (r UserRoleRepository) UpdateByModel(model *models.UserRole) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", model.ID).Updates(model)
	return err.Error
}

func (r UserRoleRepository) GetById(
	ctx context.Context,
	id uint,
) (models.UserRole, error) {
	var result models.UserRole
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Debug().Model(r.model).First(&result, "id=?", id)
	return result, err.Error
}

func (r UserRoleRepository) GetUserRolesByUserId(
	ctx context.Context, 
	userId string,
) ([]models.Role, error) {

	session := r.dbConnection.Session().WithContext(ctx)
	var result []models.Role

	stmt := `
		SELECT 
			"@rolesTable"."id",
			"@rolesTable"."name",
			"@rolesTable"."description"
		FROM @rolesTable
		JOIN 
			@userRolesTable ON "@userRolesTable"."role_id" = "@rolesTable"."id"
		WHERE
			"@userRolesTable"."user_id" = @userId
		ORDER BY id
	`

	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("userId", userId),
	).Find(&result)

	return result, err.Error
}

func (r UserRoleRepository) GetUserRolesByUserIdAsStringSlice(
	ctx context.Context, 
	userId string,
) ([]string, error) {

	session := r.dbConnection.Session().WithContext(ctx)
	var result []string

	stmt := `
		SELECT
			"@rolesTable"."name"
		FROM @rolesTable
		JOIN 
			@userRolesTable ON "@userRolesTable"."role_id" = "@rolesTable"."id"
		WHERE
			"@userRolesTable"."user_id" = @userId
		ORDER BY name
	`

	err := session.Raw(
		stmt, models.GetTableNameMap(), sql.Named("userId", userId),
	).Find(&result)

	return result, err.Error
}

func (r UserRoleRepository) GetUserRolesByUsername(
	ctx context.Context,
	userName string,
) ([]models.Role, error) {
	var result []models.Role
	session := r.dbConnection.Session().WithContext(ctx)

	stmt := `
		SELECT
			"@rolesTable"."id",
			"@rolesTable"."name",
			"@rolesTable"."description"
		FROM @rolesTable
		JOIN 
			@userRolesTable ON "@userRolesTable"."role_id" = "@rolesTable"."id"
		WHERE
			"@userRolesTable"."user_id" = (
				SELECT "@userTable"."id" FROM @userTable WHERE "@userTable"."username" = @userName
			)
	`

	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("userName", userName),
	).Find(&result)

	return result, err.Error
}

func (r UserRoleRepository) SafeDeleteById(id uint) error {
	session := r.dbConnection.Session()
	safeDeleteFields := make(map[string]any)
	safeDeleteFields["IsActive"] = false
	safeDeleteFields["DeletedAt"] = time.Now()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(safeDeleteFields)
	return err.Error
}

func (r UserRoleRepository) DeleteById(id uint) error {
	session := r.dbConnection.Session()
	err := session.Debug().Where("id=?", id).Unscoped().Delete(r.model)
	return err.Error
}
