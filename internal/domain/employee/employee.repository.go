package employee

type Repository interface {
	Create(user *Employee) error
	FindByID(id uint) (*Employee, error)
}
