package domain

import "errors"

// PackSize represents the size of a pack. Later could be extended with more attributes.
type PackSize int

func NewPack(size int) (Pack, error) {
	if size <= 0 {
		return Pack{}, errors.New("pack size must be a positive integer")
	}
	return Pack{Size: PackSize(size)}, nil
}

type Pack struct {
	Size PackSize
}

// PackingResult represents the result of a packing calculation.
type PackingResult struct {
	PacksUsed map[PackSize]int
}

// PackingOrder defines the interface for calculating the optimal packing of an order.
type PackingOrder interface {
	CalculatePacks(orderQuantity int, availablePacks []Pack) (PackingResult, error)
}

// PackList holds a list of available packs and ensures no duplicate sizes exist.
type PackList struct {
	Packs []Pack
}

// invariant: no duplicate pack sizes in the list
func NewPackList(packs []Pack) (*PackList, error) {
	seen := make(map[PackSize]bool)

	for i := range packs {
		if packs[i].Size <= 0 {
			return nil, errors.New("pack size must be a positive integer")
		}
		if seen[packs[i].Size] {
			return nil, errors.New("duplicate pack sizes are not allowed")
		}
		seen[packs[i].Size] = true
	}

	return &PackList{Packs: packs}, nil
}
