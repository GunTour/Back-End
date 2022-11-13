package database

import (
	"GunTour/config"

	ar "GunTour/features/admin/repository"
	rr "GunTour/features/ranger/repository"

	gr "GunTour/features/google/repository"
	ur "GunTour/features/users/repository"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// FUNC TO INITIALIZE DATABASE CONFIG
func InitDB(c *config.AppConfig) *gorm.DB {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Error("db config error: ", err.Error())
		return nil
	}

	return db
}

// FUNC TO MIGRATE TABLE TO DATABASE
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&ur.User{})
	db.AutoMigrate(&rr.Ranger{})
	db.AutoMigrate(&ur.Booking{})
	db.AutoMigrate(&ur.Product{})
	db.AutoMigrate(&ur.BookingProduct{})
	db.AutoMigrate(&ar.Climber{})
	db.AutoMigrate(&ar.Pesan{})
	db.AutoMigrate(&gr.Code{})
}
