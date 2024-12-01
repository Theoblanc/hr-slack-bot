package repository

import (
	"time"

	"github.com/TheoBlanc/AttendEase/internal/domain/attendance"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) attendance.Repository {
	return &AttendanceRepository{db: db}
}

func (r *AttendanceRepository) Save(a *attendance.Attendance) error {
	return r.db.Save(a).Error
}

func (r *AttendanceRepository) FindByID(id uint) (*attendance.Attendance, error) {
	var att attendance.Attendance
	result := r.db.First(&att, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &att, nil
}

func (r *AttendanceRepository) FindTodayAttendance(employeeID uint) (*attendance.Attendance, error) {
	var att attendance.Attendance
	today := time.Now().Format("2006-01-02")

	result := r.db.Where("employee_id = ? AND DATE(created_at) = ?", employeeID, today).
		First(&att)

	if result.Error != nil {
		return nil, result.Error
	}
	return &att, nil
}

func (r *AttendanceRepository) FindByDateRange(employeeID uint, start, end time.Time) ([]attendance.Attendance, error) {
	var attendances []attendance.Attendance

	result := r.db.Where("employee_id = ? AND created_at BETWEEN ? AND ?",
		employeeID, start, end).
		Order("created_at DESC").
		Find(&attendances)

	if result.Error != nil {
		return nil, result.Error
	}
	return attendances, nil
}
