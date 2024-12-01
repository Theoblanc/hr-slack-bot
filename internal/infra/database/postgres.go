package database

import (
	"fmt"
	"log"

	"github.com/TheoBlanc/AttendEase/internal/domain/attendance"
	"github.com/TheoBlanc/AttendEase/internal/domain/employee"
	"github.com/TheoBlanc/AttendEase/internal/domain/leave"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func PostgresDBInitialize(dsn string) (*PostgresDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = migrateSchema(db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database schema: %w", err)
	}

	return &PostgresDB{DB: db}, nil
}

func migrateSchema(db *gorm.DB) error {
	err := db.AutoMigrate(
		&employee.Employee{},
		&attendance.Attendance{},
		&leave.LeaveRequest{},
		&leave.LeaveBalance{},
	)
	if err != nil {
		return err
	}

	log.Println("Database schema migrated successfully")
	return nil
}
