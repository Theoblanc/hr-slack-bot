package leave

import (
	"time"

	"github.com/TheoBlanc/AttendEase/internal/domain"
	"github.com/TheoBlanc/AttendEase/internal/domain/employee"
)

type LeaveRequest struct {
	domain.BaseModel
	EmployeeID uint
	Employee   employee.Employee `gorm:"foreignKey:EmployeeID"`
	LeaveType  string            `gorm:"size:20;not null"`
	StartDate  time.Time         `gorm:"not null"`
	Status     string            `gorm:"size:10;not null"`
	ApproverID *uint
	Approver   *employee.Employee `gorm:"foreignKey:ApproverID"`
}
