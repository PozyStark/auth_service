package repository

import (
	db "auth-service/internal/db/gorm"
	"auth-service/internal/models"
	"context"
	"database/sql"
	"time"
)

type RolePermissionRepository struct {
	model        *models.RolePermission
	dbConnection *db.GormDbConnection
}

func NewGormRolePermissionRepository(dbConnection *db.GormDbConnection) RolePermissionRepository {
	return RolePermissionRepository{dbConnection: dbConnection}
}

func (r RolePermissionRepository) Save(model *models.RolePermission) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(model)
	return err.Error
}

func (r RolePermissionRepository) SaveByMap(valuesMap map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(valuesMap)
	return err.Error
}

func (r RolePermissionRepository) UpdateById(id uint, values map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(values)
	return err.Error
}

func (r RolePermissionRepository) UpdateByModel(model *models.RolePermission) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", model.ID).Updates(model)
	return err.Error
}

func (r RolePermissionRepository) GetById(ctx context.Context, id uint) (models.RolePermission, error) {
	var result models.RolePermission
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Debug().Model(r.model).First(&result, "id=?", id)
	return result, err.Error
}

func (r RolePermissionRepository) GetRolePermissionsByRoleId(
	ctx context.Context,
	roleId uint,
) ([]models.Permission, error) {

	var result []models.Permission
	session := r.dbConnection.Session().WithContext(ctx)

	stmt := `
		SELECT 
			"@permissionsTable"."id",
			"@permissionsTable"."name",
			"@permissionsTable"."description"
		FROM
			@permissionsTable
		JOIN 
			@rolePermissionsTable ON "@rolePermissionsTable"."permission_id" = "@permissionsTable"."id"
		WHERE 
			"@rolePermissionsTable"."role_id" = @roleId
	`

	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("roleId", roleId),
	).Find(&result)

	return result, err.Error
}

func (r RolePermissionRepository) GetRolePermissionsByRoleName(
	ctx context.Context,
	roleName string,
) ([]models.Permission, error) {

	var result []models.Permission
	session := r.dbConnection.Session().WithContext(ctx)

	stmt := `
		SELECT 
			"@permissionsTable"."id",
			"@permissionsTable"."name",
			"@permissionsTable"."description"
		FROM
			@permissionsTable
		JOIN 
			@rolePermissionsTable ON "@rolePermissionsTable"."permission_id" = "@permissionsTable"."id"
		WHERE 
			"@rolePermissionsTable"."role_id" = (
				SELECT "@rolesTable"."id" FROM @rolesTable WHERE "@rolesTable"."name" = @roleName
			)
	`

	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("roleName", roleName),
	).Find(&result)

	return result, err.Error
}

func (r RolePermissionRepository) SafeDeleteById(id uint) error {
	session := r.dbConnection.Session()
	safeDeleteFields := make(map[string]any)
	safeDeleteFields["IsActive"] = false
	safeDeleteFields["DeletedAt"] = time.Now()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(safeDeleteFields)
	return err.Error
}

func (r RolePermissionRepository) DeleteById(id uint) error {
	session := r.dbConnection.Session()
	err := session.Debug().Where("id=?", id).Unscoped().Delete(r.model)
	return err.Error
}
