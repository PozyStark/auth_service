package db

import (
	"time"
	"gorm.io/gorm"
)

type GormDbConnection struct {
	dbConnection  *gorm.DB
	gormConfig    *gorm.Config
	sessionConfig *gorm.Session
	sqlDialect    gorm.Dialector
}

func NewGormDbConnection(sqlDialect gorm.Dialector) (dbConnection *GormDbConnection) {
	return &GormDbConnection{
		sqlDialect: sqlDialect,
		gormConfig: &gorm.Config{},
	}
}

func NewGormDbConnectionWithConfig(
	sqlDialect gorm.Dialector,
	sessionConfig *gorm.Session,
	gormConfig *gorm.Config,
) (dbConnection *GormDbConnection) {
	return &GormDbConnection{
		sqlDialect:    sqlDialect,
		sessionConfig: sessionConfig,
		gormConfig:    gormConfig,
	}
}

func (c *GormDbConnection) Open() (dbConnection *GormDbConnection, err error) {
	c.dbConnection, err = gorm.Open(c.sqlDialect, c.gormConfig)
	return c, err
}

func (c *GormDbConnection) GetConnection() (*gorm.DB) {
	return c.dbConnection
}

func (c *GormDbConnection) Session() (dbSession *gorm.DB) {
	return c.dbConnection.Session(c.sessionConfig)
}

func (c *GormDbConnection) SessionWithConfig(sessionConfig *gorm.Session) (dbSession *gorm.DB) {
	return c.dbConnection.Session(sessionConfig)
}

func (c *GormDbConnection) Close() (err error) {
	sqlStmt, err := c.dbConnection.DB()
	if err != nil {
		return err
	}
	return sqlStmt.Close()
}

func (c *GormDbConnection) SetMaxOpenConnections(limit int) (err error) {
	sqlStmt, err := c.dbConnection.DB()
	if err != nil {
		return err
	}
	sqlStmt.SetMaxOpenConns(limit)
	return nil
}

func (c *GormDbConnection) SetMaxIdleConnections(limit int) (err error) {
	sqlStmt, err := c.dbConnection.DB()
	if err != nil {
		return err
	}
	sqlStmt.SetMaxIdleConns(limit)
	return nil
}

func (c *GormDbConnection) SetMaxConnectionLifetime(duration time.Duration) (err error) {
	sqlStmt, err := c.dbConnection.DB()
	if err != nil {
		return err
	}
	sqlStmt.SetConnMaxLifetime(duration)
	return nil
}

func (c *GormDbConnection) SetMaxConnectionIdletime(duration time.Duration) (err error) {
	sqlStmt, err := c.dbConnection.DB()
	if err != nil {
		return err
	}
	sqlStmt.SetConnMaxIdleTime(duration)
	return nil
}
