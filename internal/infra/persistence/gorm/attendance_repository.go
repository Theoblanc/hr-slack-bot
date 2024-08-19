package repository

import (
	"github.com/TheoBlanc/AttendEase/internal/domain/attendance"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) attendance.Repository {
	return &AttendanceRepository{db: db}
}

func (r *AttendanceRepository) Create(attendance *attendance.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *AttendanceRepository) GetByID(id uint) (*attendance.Attendance, error) {
	var a attendance.Attendance
	err := r.db.First(&a, id).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AttendanceRepository) Update(attendance *attendance.Attendance) error {
	return r.db.Save(attendance).Error
}

func (r *AttendanceRepository) Delete(id uint) error {
	return r.db.Delete(&attendance.Attendance{}, id).Error
}
