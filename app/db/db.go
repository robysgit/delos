package db

import (
	entities "delos/app/db/entities"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=test dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database = db
	db.AutoMigrate(&entities.ApiHistoryEntity{})
	db.AutoMigrate(&entities.FarmEntity{})
	db.AutoMigrate(&entities.PondEntity{})
	db.AutoMigrate(&entities.FarmVPondEntity{})

	pDb, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	pDb.SetMaxIdleConns(5)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	pDb.SetMaxOpenConns(20)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	pDb.SetConnMaxLifetime(time.Minute * 5)
}

func GetDB() *gorm.DB {
	return database
}
