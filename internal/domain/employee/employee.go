package employee

import (
	"time"

	"github.com/TheoBlanc/AttendEase/internal/domain"
)

type Employee struct {
	domain.BaseModel
	Name          string    `gorm:"size:100;not null"`    // 이름
	Email         string    `gorm:"size:100;uniqueIndex"` // 이메일
	PhoneNumber   string    `gorm:"size:15"`              // 전화번호
	Position      string    `gorm:"size:50"`              // 직책 (예: 개발자, 디자이너 등)
	Department    string    `gorm:"size:50"`              // 부서 (예: 인사부, 개발팀 등)
	DateOfJoining time.Time `gorm:"not null"`             // 입사일
	Address       string    `gorm:"size:255"`             // 주소
	IsActive      bool      `gorm:"default:true"`         // 활성화 여부 (퇴사 여부 등)

}
