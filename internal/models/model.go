package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	TABLE_PREFIX = "authentication_service"
)

type Role struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:always"` // При ручной миграции выкл генерацию ID <-:create
	Name        string `gorm:"unique;size:256;not null"`
	Description string `gorm:"size:256"`
}

func (Role) TableName() string {
	return TABLE_PREFIX + "_roles"
}

type Group struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:always"` // При ручной миграции выкл генерацию ID <-:create
	Name        string `gorm:"unique;size:256;not null"`
	Description string `gorm:"size:256"`
}

func (Group) TableName() string {
	return TABLE_PREFIX + "_groups"
}

type Permission struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:always"` // При ручной миграции выкл генерацию ID <-:create
	Name        string `gorm:"unique;size:256;not null"`
	Description string `gorm:"size:256"`
}

func (Permission) TableName() string {
	return TABLE_PREFIX + "_permissions"
}

type UserGroup struct {
	ID      uint   `gorm:"index;primaryKey;autoIncrement:always"` // При ручной миграции выкл генерацию ID <-:create
	UserID  string `gorm:"index;not null;type:uuid"`
	GroupID uint   `gorm:"not null"`
	User    User
	Group   Group
}

func (UserGroup) TableName() string {
	return TABLE_PREFIX + "_user_groups"
}

type UserRole struct {
	ID     uint   `gorm:"index;primaryKey;autoIncrement:always"` // При ручной миграции выкл генерацию ID <-:create
	UserID string `gorm:"index;not null;type:uuid"`
	RoleID uint   `gorm:"not null"`
	User   User
	Role   Role
}

func (UserRole) TableName() string {
	return TABLE_PREFIX + "_user_roles"
}

type UserPermission struct {
	ID           uint   `gorm:"index;primaryKey;autoIncrement:always"` // При ручной миграции выкл генерацию ID <-:create
	UserID       string `gorm:"index;not null;type:uuid"`
	PermissionID uint   `gorm:"not null"`
	User         User
	Permission   Permission
}

func (UserPermission) TableName() string {
	return TABLE_PREFIX + "_user_permissions"
}

type RolePermission struct {
	ID           uint `gorm:"primaryKey;autoIncrement:always"` // При ручной миграции выкл генерацию ID <-:create
	RoleID       uint `gorm:"index;not null"`
	PermissionID uint `gorm:"not null"`
	Role         Role
	Permission   Permission
}

func (RolePermission) TableName() string {
	return TABLE_PREFIX + "_role_permissions"
}

type GroupPermission struct {
	ID           uint `gorm:"primaryKey;autoIncrement:always"` // При ручной миграции выкл генерацию ID <-:create
	GroupID      uint `gorm:"index;not null"`
	PermissionID uint `gorm:"not null"`
	Group        Group
	Permission   Permission
}

func (GroupPermission) TableName() string {
	return TABLE_PREFIX + "_group_permissions"
}

type RegistryToken struct {
	ID        string    `gorm:"index;primaryKey;type:uuid;default:uuidv7()"`
	UserID    string    `gorm:"index;not null;type:uuid"`
	Jti       string    `gorm:"not null;type:uuid"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	ExpireAt  time.Time `gorm:"not null"`
	Active    bool      `gorm:"not null;default:true"`
}

func (RegistryToken) TableName() string {
	return TABLE_PREFIX + "_registry_tokens"
}

type User struct {
	ID          string    `gorm:"index;primaryKey;type:uuid;default:uuidv7()"` // При ручной миграции выкл генерацию ID <-:create
	Username    string    `gorm:"index;unique;size:256;not null"`
	PhoneNumber string    `gorm:"index;unique;not null"`
	Email       string    `gorm:"index;unique;size:256"`
	FirstName   string    `gorm:"size:256"`
	MiddleName  string    `gorm:"size:256"`
	LastName    string    `gorm:"size:256"`
	Password    string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:null"`
	UpdatedAt   time.Time `gorm:"default:null"`
	DeletedAt   time.Time `gorm:"default:null"`
	LastLogin   time.Time `gorm:"default:null"`
	Superuser   bool      `gorm:"default:false"`
	Active      bool      `gorm:"default:true"`
	Age         int
}

func (User) TableName() string {
	return TABLE_PREFIX + "_users"
}

func GetTableNameMap() map[string]any {

	return map[string]any{
		"rolesTable":            gorm.Expr(Role{}.TableName()),
		"groupsTable":           gorm.Expr(Group{}.TableName()),
		"permissionsTable":      gorm.Expr(Permission{}.TableName()),
		"userRolesTable":        gorm.Expr(UserRole{}.TableName()),
		"userGroupsTable":       gorm.Expr(UserGroup{}.TableName()),
		"userPermissionsTable":  gorm.Expr(UserPermission{}.TableName()),
		"rolePermissionsTable":  gorm.Expr(RolePermission{}.TableName()),
		"groupPermissionsTable": gorm.Expr(GroupPermission{}.TableName()),
		"registryTokenTable":    gorm.Expr(RegistryToken{}.TableName()),
		"userTable":             gorm.Expr(User{}.TableName()),
	}

}
