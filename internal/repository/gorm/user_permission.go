package repository

import (
	db "auth-service/internal/db/gorm"
	"auth-service/internal/models"
	"context"
	"database/sql"
	"time"
)

type UserPermissionRepository struct {
	model        *models.UserPermission
	dbConnection *db.GormDbConnection
}

func NewGormUserPermissionRepository(dbConnection *db.GormDbConnection) UserPermissionRepository {
	return UserPermissionRepository{dbConnection: dbConnection}
}

func (r UserPermissionRepository) Save(model *models.UserPermission) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(model)
	return err.Error
}

func (r UserPermissionRepository) SaveByMap(valuesMap map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(valuesMap)
	return err.Error
}

func (r UserPermissionRepository) UpdateById(id uint, values map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(values)
	return err.Error
}

func (r UserPermissionRepository) UpdateByModel(model *models.UserPermission) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", model.ID).Updates(model)
	return err.Error
}

func (r UserPermissionRepository) GetById(
	ctx context.Context,
	id uint,
) (models.UserPermission, error) {
	var result models.UserPermission
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Debug().Model(r.model).First(&result, "id=?", id)
	return result, err.Error
}

func (r UserPermissionRepository) GetUserPermissionsByUserId(
	ctx context.Context,
	userId string,
) ([]models.Permission, error) {

	var result []models.Permission
	session := r.dbConnection.Session().WithContext(ctx)

	stmt := `
		SELECT 
			"@permissionsTable"."id" AS id,
			"@permissionsTable"."name",
			"@permissionsTable"."description"
		FROM @userPermissionsTable
		JOIN @permissionsTable ON "@userPermissionsTable"."permission_id" = "@permissionsTable"."id"
		WHERE "@userPermissionsTable"."user_id" = @userId
	`

	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("userId", userId),
	).Find(&result)

	return result, err.Error
}

func (r UserPermissionRepository) GetAllUserPermissionsByUserId(
	ctx context.Context,
	userId string,
) ([]models.Permission, error) {

	var result []models.Permission
	session := r.dbConnection.Session().WithContext(ctx)

	stmt := `
		SELECT 
			"@permissionsTable"."id" AS id,
			"@permissionsTable"."name",
			"@permissionsTable"."description"
		FROM @rolePermissionsTable
		JOIN 
			@permissionsTable ON "@permissionsTable"."id" = "@rolePermissionsTable"."permission_id"
		WHERE role_id IN (
			SELECT role_id FROM @userRolesTable WHERE user_id = @userId
		)
		UNION
		SELECT 
			"@permissionsTable"."id" AS id,
			"@permissionsTable"."name",
			"@permissionsTable"."description"
		FROM @groupPermissionsTable 
		JOIN 
			@permissionsTable ON "@permissionsTable"."id" = "@groupPermissionsTable"."permission_id"
		WHERE group_id IN (
			SELECT group_id FROM @userGroupsTable WHERE user_id = @userId
		)
		UNION
		SELECT 
			"@permissionsTable"."id" AS id,
			"@permissionsTable"."name",
			"@permissionsTable"."description"
		FROM @userPermissionsTable
		JOIN 
			@permissionsTable ON "@permissionsTable"."id" = "@userPermissionsTable"."permission_id"
		WHERE "@userPermissionsTable"."user_id" = @userId 
		ORDER BY id;
	`

	err := session.Raw(
		stmt, models.GetTableNameMap(), sql.Named("userId", userId),
	).Find(&result)

	return result, err.Error
}

func (r UserPermissionRepository) GetAllUserPermissionsByUserIdAsStringSlice(
	ctx context.Context,
	userId string,
) ([]string, error) {

	var result []string
	session := r.dbConnection.Session().WithContext(ctx)

	stmt := `
		SELECT 
			"@permissionsTable"."name"
		FROM @rolePermissionsTable
		JOIN 
			@permissionsTable ON "@permissionsTable"."id" = "@rolePermissionsTable"."permission_id"
		WHERE role_id IN (
			SELECT role_id FROM @userRolesTable WHERE user_id = @userId
		)
		UNION
		SELECT 
			"@permissionsTable"."name"
		FROM @groupPermissionsTable
		JOIN 
			@permissionsTable ON "@permissionsTable"."id" = "@groupPermissionsTable"."permission_id"
		WHERE group_id IN (
			SELECT group_id FROM @userGroupsTable WHERE user_id = @userId
		)
		UNION
		SELECT
			"@permissionsTable"."name"
		FROM @userPermissionsTable
		JOIN 
			@permissionsTable ON "@permissionsTable"."id" = "@userPermissionsTable"."permission_id"
		WHERE "@userPermissionsTable"."user_id" = @userId 
		ORDER BY name;
	`

	err := session.Raw(
		stmt, models.GetTableNameMap(), sql.Named("userId", userId),
	).Find(&result)

	return result, err.Error
}

func (r UserPermissionRepository) GetUserPermissionsByUsername(
	ctx context.Context,
	userName string,
) ([]models.Permission, error) {

	var result []models.Permission
	session := r.dbConnection.Session().WithContext(ctx)

	stmt := `
		SELECT
			"@permissionsTable"."id" AS id,
			"@permissionsTable"."name",
			"@permissionsTable"."description"
		FROM @userPermissionsTable
		JOIN 
			@permissionsTable ON "@permissionsTable"."id" = "@userPermissionsTable"."permission_id"
		WHERE "@userPermissionsTable"."user_id" = (
			SELECT "@userTable"."id" FROM @userTable WHERE "@userTable"."username" = @userName
		)
	`

	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("userName", userName),
	).Find(&result)

	return result, err.Error
}

func (r UserPermissionRepository) GetAllUserPermissionsByUsername(
	ctx context.Context,
	userName string,
) ([]models.Permission, error) {

	var result []models.Permission
	session := r.dbConnection.Session().WithContext(ctx)

	stmt := `
		WITH temp_user AS (
			SELECT "@userTable"."id" AS id
			FROM @userTable
			WHERE "@userTable"."username" = @userName
		)
		SELECT 
			"@permissionsTable"."id" AS id,
			"@permissionsTable"."name",
			"@permissionsTable"."description"
		FROM @rolePermissionsTable
		JOIN 
			@permissionsTable ON "@permissionsTable"."id" = "@rolePermissionsTable"."permission_id"
		WHERE role_id IN (
			SELECT role_id FROM @userRolesTable 
			WHERE "@userRolesTable"."user_id" = (SELECT id FROM temp_user)
		)
		UNION
		SELECT 
			"@permissionsTable"."id" AS id,
			"@permissionsTable"."name",
			"@permissionsTable"."description"
		FROM @groupPermissionsTable
		JOIN 
			@permissionsTable ON "@permissionsTable"."id" = "@groupPermissionsTable"."permission_id"
		WHERE group_id IN (
			SELECT group_id FROM @userGroupsTable 
			WHERE "@userGroupsTable"."user_id" = (SELECT id FROM temp_user)
		)
		UNION
		SELECT 
			"@permissionsTable"."id" AS id,
			"@permissionsTable"."name",
			"@permissionsTable"."description"
		FROM @userPermissionsTable
		JOIN 
			@permissionsTable ON "@permissionsTable"."id" = "@userPermissionsTable"."permission_id"
		WHERE "@userPermissionsTable"."user_id" = (SELECT id FROM temp_user) 
		ORDER BY id;
	`

	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("userName", userName),
	).Find(&result)

	return result, err.Error
}

func (r UserPermissionRepository) SafeDeleteById(id uint) error {
	session := r.dbConnection.Session()
	safeDeleteFields := make(map[string]any)
	safeDeleteFields["IsActive"] = false
	safeDeleteFields["DeletedAt"] = time.Now()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(safeDeleteFields)
	return err.Error
}

func (r UserPermissionRepository) DeleteById(id uint) error {
	session := r.dbConnection.Session()
	err := session.Debug().Where("id=?", id).Unscoped().Delete(r.model)
	return err.Error
}
