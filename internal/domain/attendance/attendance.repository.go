package attendance

type Repository interface {
	Create(attendance *Attendance) error
	GetByID(id uint) (*Attendance, error)
	Update(attendance *Attendance) error
	Delete(id uint) error
}
