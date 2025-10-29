package infra

import (
	"sync"

	"github.com/stolendog/order-packs-optimizer/internal/domain"
)

type MemoryRepository struct {
	packs []domain.Pack
	mutex sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		packs: []domain.Pack{},
	}
}

func (r *MemoryRepository) SetPacks(packs []domain.Pack) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.packs = packs
	return nil
}

func (r *MemoryRepository) GetAllPacks() ([]domain.Pack, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := make([]domain.Pack, len(r.packs))
	copy(result, r.packs)

	return result, nil
}
