package domain

type Repository interface {
	ReplacePacks([]Pack) error
	GetAllPacks() ([]Pack, error)
}
