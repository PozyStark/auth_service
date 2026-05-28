package db

import (
	"auth-service/internal/config"
	"database/sql"
	"time"
	_ "github.com/lib/pq"
)

type DbConnection struct {
	dbConnection  *sql.DB
	driverName string
	dataSource string
}

func NewDbConnection(driverName string, dataSource string) (*DbConnection) {
	return &DbConnection{
		driverName: driverName,
		dataSource: dataSource,
	}
}

func NewDbConnetion(driverName string, dbConfig config.DbConfig) (*DbConnection) {
	return &DbConnection{
		driverName: driverName,
		dataSource: GetDataSource(driverName, dbConfig),
	}
}

func (c *DbConnection) Open() (dbConnection *DbConnection, err error) {
	c.dbConnection, err = sql.Open(c.driverName, c.dataSource)
	return c, err
}

func (c *DbConnection) Close() (err error) {
	return c.dbConnection.Close()
}

func (c *DbConnection) GetConnection() (*sql.DB) {
	return c.dbConnection
}

func (c *DbConnection) SetMaxOpenConnections(limit int) {
	c.dbConnection.SetMaxOpenConns(limit)
}

func (c *DbConnection) SetMaxIdleConnections(limit int) {
	c.dbConnection.SetMaxIdleConns(limit)
}

func (c *DbConnection) SetMaxConnectionLifetime(duration time.Duration) {
	c.dbConnection.SetConnMaxLifetime(duration)
}

func (c *DbConnection) SetMaxConnectionIdletime(duration time.Duration) {
	c.dbConnection.SetConnMaxIdleTime(duration)
}