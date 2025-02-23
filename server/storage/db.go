package storage

import (
	"fmt"
	"hexxcore/config"
	"hexxcore/types"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresStorage() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Envs.DB_URL))
	fmt.Println(config.Envs.DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&types.User{}, &types.Schedule{}, &types.Subject{}, &types.Attendance{}); err != nil {
		log.Fatal(err)
	}
}

func InsertPredefinedSchedule(db *gorm.DB) error {
	var count int64
	db.Model(&types.Schedule{}).Count(&count)

	if count > 0 {
		log.Println("Schedule table already populated, skipping insertion.")
		return nil
	}

	// Predefined schedule data

	// Bulk insert
	if err := db.Create(&config.PredefinedSchedule).Error; err != nil {
		return fmt.Errorf("failed to insert schedule data: %w", err)
	}

	log.Println("Predefined schedule inserted successfully.")
	return nil
}
