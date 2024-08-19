package attendance

import (
	"time"

	"github.com/TheoBlanc/AttendEase/internal/domain"
	"github.com/TheoBlanc/AttendEase/internal/domain/employee"
)

type Attendance struct {
	domain.BaseModel
	EmployeeID    uint              // 외래 키로 Employee와 연결
	Employee      employee.Employee `gorm:"foreignKey:EmployeeID"`
	ClockIn       *time.Time
	ClockOut      *time.Time
	WorkHours     float64
	Location      string `gorm:"size:255"`
	IsLate        bool
	IsEarlyLeave  bool
	OvertimeHours float64
}
