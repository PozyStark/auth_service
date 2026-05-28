package repository

import (
	db "auth-service/internal/db/gorm"
	"auth-service/internal/models"
	"context"
	"database/sql"
	"time"
)

type UserGroupRepository struct {
	model        *models.UserGroup
	dbConnection *db.GormDbConnection
}

func NewGormUserGroupRepository(dbConnection *db.GormDbConnection) UserGroupRepository {
	return UserGroupRepository{dbConnection: dbConnection}
}

func (r UserGroupRepository) Save(model *models.UserGroup) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(model)
	return err.Error
}

func (r UserGroupRepository) SaveByMap(valuesMap map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Create(valuesMap)
	return err.Error
}

func (r UserGroupRepository) UpdateById(id uint, values map[string]any) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(values)
	return err.Error
}

func (r UserGroupRepository) UpdateByModel(model *models.UserGroup) error {
	session := r.dbConnection.Session()
	err := session.Debug().Model(r.model).Where("id=?", model.ID).Updates(model)
	return err.Error
}

func (r UserGroupRepository) GetById(
	ctx context.Context,
	id uint,
) (models.UserGroup, error) {
	var result models.UserGroup
	session := r.dbConnection.Session().WithContext(ctx)
	err := session.Debug().Model(r.model).First(&result, "id=?", id)
	return result, err.Error
}

func (r UserGroupRepository) GetUserGroupsByUserId(
	ctx context.Context,
	userId string,
) ([]models.Group, error) {

	session := r.dbConnection.Session().WithContext(ctx)
	var result []models.Group

	stmt := `
		SELECT 
			"@groupsTable"."id",
			"@groupsTable"."name",
			"@groupsTable"."description"
		FROM @groupsTable
		JOIN 
			@userGroupsTable ON "@userGroupsTable"."group_id" = "@groupsTable"."id"
		WHERE
			"@userGroupsTable"."user_id" = @userId
		ORDER BY id
	`

	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("userId", userId),
	).Find(&result)

	return result, err.Error
}

func (r UserGroupRepository) GetUserGroupsByUserIdAsStringSlice(
	ctx context.Context,
	userId string,
) ([]string, error) {

	session := r.dbConnection.Session().WithContext(ctx)
	var result []string

	stmt := `
		SELECT 
			"@groupsTable"."name"
		FROM @groupsTable
		JOIN 
			@userGroupsTable ON "@userGroupsTable"."group_id" = "@groupsTable"."id"
		WHERE
			"@userGroupsTable"."user_id" = @userId
		ORDER BY name
	`
	
	err := session.Raw(
		stmt, models.GetTableNameMap(), sql.Named("userId", userId),
	).Find(&result)

	return result, err.Error
}

func (r UserGroupRepository) GetUserGroupsByUsername(
	ctx context.Context,
	userName string,
) ([]models.Group, error) {

	var result []models.Group
	session := r.dbConnection.Session().WithContext(ctx)

	stmt := `
		SELECT 
			"@groupsTable"."id",
			"@groupsTable"."name",
			"@groupsTable"."description"
		FROM @groupsTable
		JOIN 
			@userGroupsTable ON "@userGroupsTable"."group_id" = "@groupsTable"."id"
		WHERE
			"@userGroupsTable"."user_id" = (
				SELECT "@userTable"."id" FROM @userTable WHERE "@userTable"."username" = @userName
			)
	`

	err := session.Debug().Raw(
		stmt, models.GetTableNameMap(), sql.Named("userName", userName),
	).Find(&result)

	return result, err.Error
}

func (r UserGroupRepository) SafeDeleteById(id uint) error {
	session := r.dbConnection.Session()
	safeDeleteFields := make(map[string]any)
	safeDeleteFields["IsActive"] = false
	safeDeleteFields["DeletedAt"] = time.Now()
	err := session.Debug().Model(r.model).Where("id=?", id).Updates(safeDeleteFields)
	return err.Error
}

func (r UserGroupRepository) DeleteById(id uint) error {
	session := r.dbConnection.Session()
	err := session.Debug().Where("id=?", id).Unscoped().Delete(r.model)
	return err.Error
}
