package leave

import "github.com/TheoBlanc/AttendEase/internal/domain/employee"

type LeaveBalance struct {
	EmployeeID         uint              `gorm:"primaryKey"`
	Employee           employee.Employee `gorm:"foreignKey:EmployeeID"`
	AnnualLeaveBalance int               `gorm:"not null"`
	SickLeaveBalance   int               `gorm:"not null"`
	OtherLeaveBalance  int               `gorm:"not null"`
}
