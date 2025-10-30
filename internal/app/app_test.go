package app

import (
	"fmt"
	"testing"

	"github.com/stolendog/order-packs-optimizer/internal/domain"
)

type mockCalculator struct{}

func (m *mockCalculator) Calculate(orderQuantity int, availablePacks []int) (map[int]int, error) {
	return nil, nil
}

type mockRepository struct {
	storeError error
	getError   error
	packs      []domain.Pack
}

func (m *mockRepository) SetPacks(packs []domain.Pack) error {
	if m.storeError != nil {
		return m.storeError
	}
	m.packs = packs
	return nil
}

func (m *mockRepository) GetAllPacks() ([]domain.Pack, error) {
	if m.getError != nil {
		return nil, m.getError
	}
	return m.packs, nil
}

// this is not full coverage unit tests for app layer just basic show of unit test structure and usage of mocks
func TestApp(t *testing.T) {
	repo := &mockRepository{}
	calculator := &mockCalculator{}
	app := NewApp(calculator, repo)

	t.Run("TestStorePackList", func(t *testing.T) {
		packs := []domain.Pack{
			{Size: 100},
			{Size: 200},
		}
		err := app.StorePackList(packs)
		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}
	})

	t.Run("TestStorePackList_Invalid", func(t *testing.T) {
		packs := []domain.Pack{
			{Size: 100},
			{Size: 100}, // duplicate size for invalid case
		}
		err := app.StorePackList(packs)
		if err == nil {
			t.Errorf("expected error for duplicate pack sizes but got none")
		}
	})

	t.Run("TestGetAllPacks", func(t *testing.T) {
		_, err := app.GetAllPacks()
		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}
	})

	t.Run("TestGetAllPacks_RepoError", func(t *testing.T) {
		repo.getError = fmt.Errorf("repository error")
		_, err := app.GetAllPacks()
		if err == nil {
			t.Errorf("expected error from repository but got none")
		}
	})
}
