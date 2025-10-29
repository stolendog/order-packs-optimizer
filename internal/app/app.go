package app

import "github.com/stolendog/order-packs-optimizer/internal/domain"

type Calculator interface {
	Calculate(orderQuantity int, availablePacks []int) (map[int]int, error)
}

// For simple calculator app maybe this is overkill but wanted to keep more closely to clean architecture principles

type App struct {
	calculator Calculator

	repo domain.Repository
}

func NewApp(calculator Calculator, repo domain.Repository) *App {
	return &App{
		calculator: calculator,
		repo:       repo,
	}
}

func (a *App) StorePackList(packs []domain.Pack) error {
	packList, err := domain.NewPackList(packs)
	if err != nil {
		return err
	}

	err = a.repo.SetPacks(packList.Packs)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) GetAllPacks() ([]domain.Pack, error) {
	packs, err := a.repo.GetAllPacks()
	if err != nil {
		return nil, err
	}

	return packs, nil
}

func (a *App) CalculatePacks(orderQuantity int) (*domain.PackingResult, error) {
	packs, err := a.GetAllPacks()
	if err != nil {
		return nil, err
	}

	var packSizes []int
	for _, pack := range packs {
		packSizes = append(packSizes, int(pack.Size))
	}

	calculatedPacks, err := a.calculator.Calculate(orderQuantity, packSizes)
	if err != nil {
		return nil, err
	}

	return domain.NewPackingResult(calculatedPacks), nil
}
