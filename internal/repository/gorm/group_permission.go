package repository

import (
	db "auth-service/internal/db/gorm"
	"auth-service/internal/models"
	"context"
	"database/sql"
	"time"
)

type GroupPermissionRepository struct {
	model        *models.GroupPermission
	dbConnection *db.GormDbConnection
}

func NewGormGroupPermissionRepository(dbConnection *db.GormDbConnection) GroupPermissionRepository {
	return GroupPermissionRepository{dbConnection: dbConnection}
}

func (r GroupPermissionRepository) Save(model *models.GroupPermission) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(model)
	return err.Error
}

func (r GroupPermissionRepository) SaveByMap(valuesMap map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(valuesMap)
	return err.Error
}

func (r GroupPermissionRepository) UpdateById(id uint, values map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(values)
	return err.Error
}

func (r GroupPermissionRepository) UpdateByModel(model *models.GroupPermission) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", model.ID).Updates(model)
	return err.Error
}

func (r GroupPermissionRepository) GetById(
	ctx context.Context,
	id uint,
) (models.GroupPermission, error) {
	var result models.GroupPermission
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Debug().Model(r.model).First(&result, "id=?", id)
	return result, err.Error
}

func (r GroupPermissionRepository) GetGroupPermissionsByGroupId(
	ctx context.Context,
	groupId uint,
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
			@groupPermissionsTable ON "@groupPermissionsTable"."permission_id" = "@permissionsTable"."id"
		WHERE 
			"@groupPermissionsTable"."group_id" = @groupId
	`
	
	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("groupId", groupId),
	).Find(&result)

	return result, err.Error
}

func (r GroupPermissionRepository) GetGroupPermissionsByGroupName(
	ctx context.Context,
	groupName string,
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
			@groupPermissionsTable ON "@groupPermissionsTable"."permission_id" = "@permissionsTable"."id"
		WHERE 
			"@groupPermissionsTable"."group_id" = (
				SELECT "@groupsTable"."id" FROM @groupsTable WHERE "@groupsTable"."name" = @groupName
			)
	`

	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("groupName", groupName),
	).Find(&result)

	return result, err.Error
}

func (r GroupPermissionRepository) SafeDeleteById(id uint) error {
	session := r.dbConnection.Session()
	safeDeleteFields := make(map[string]any)
	safeDeleteFields["IsActive"] = false
	safeDeleteFields["DeletedAt"] = time.Now()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(safeDeleteFields)
	return err.Error
}

func (r GroupPermissionRepository) DeleteById(id uint) error {
	session := r.dbConnection.Session()
	err := session.Debug().Where("id=?", id).Unscoped().Delete(r.model)
	return err.Error
}
