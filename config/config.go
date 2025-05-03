package config

import (
	"dibimbing_golang_ticketing/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func buildDSN() string {
	dsn := GetEnv("DB_DSN", "")
	if dsn != "" {
		return dsn
	}
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3306")
	user := GetEnv("DB_USER", "root")
	password := GetEnv("DB_PASSWORD", "")
	dbname := GetEnv("DB_NAME", "travel_booking")
	return user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func InitDB() *gorm.DB {
	dsn := buildDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	if err := db.AutoMigrate(&entity.User{}, &entity.Event{}, &entity.Ticket{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	return db
}
