package repository

type WARepository interface {
}

type waRepository struct {
}

func NewWARepository() WARepository {
	return &waRepository{}
}
