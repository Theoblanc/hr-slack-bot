package attendance

import "time"

type Repository interface {
	Save(*Attendance) error
	FindByID(id uint) (*Attendance, error)
	FindTodayAttendance(employeeID uint) (*Attendance, error)
	FindByDateRange(employeeID uint, start, end time.Time) ([]Attendance, error)
}
