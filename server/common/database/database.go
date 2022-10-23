package database

import (
	"beluga/server/common/logger"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func GetDB() *gorm.DB {
	if globalDB == nil {
		log := logger.GetLogger()
		log.Fatal("Database is not init")
	}
	return globalDB
}

func InitDBPool(driver string, dsn string, maxIdleConns int, maxOpenConns int, maxLifeTime time.Duration) error {
	var db *gorm.DB
	var err error
	switch driver {
	case "mysql":
		db, err = InitMySql(dsn)
	case "postgres":
		db, err = InitPostgreSql(dsn)
	case "sqlite":
		db, err = InitSqLite(dsn)
	case "sqlserver":
		db, err = InitSqlServer(dsn)
	default:
		return errors.New("driver not support")
	}
	if err != nil {
		return err
	}
	dbPool, err := db.DB()
	dbPool.SetMaxIdleConns(maxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(maxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	dbPool.SetConnMaxLifetime(maxLifeTime)
	globalDB = db
	return err
}

func InitMySql(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func InitPostgreSql(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	return db, err
}

func InitSqLite(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	return db, err
}

func InitSqlServer(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	return db, err
}
