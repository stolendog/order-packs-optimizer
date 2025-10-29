package domain

type Repository interface {
	SetPacks([]Pack) error
	GetAllPacks() ([]Pack, error)
}
