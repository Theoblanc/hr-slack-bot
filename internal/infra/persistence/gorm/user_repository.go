package repository

import (
	"github.com/TheoBlanc/AttendEase/internal/domain/employee"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) employee.Repository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(employee *employee.Employee) error {
	return r.db.Create(employee).Error
}

func (r *EmployeeRepository) FindByID(id uint) (*employee.Employee, error) {
	var e employee.Employee
	err := r.db.First(&e, id).Error
	if err != nil {
		return nil, err
	}
	return &e, nil
}
