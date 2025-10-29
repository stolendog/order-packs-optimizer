package domain

import "errors"

func NewPack(size int) (Pack, error) {
	if size <= 0 {
		return Pack{}, errors.New("pack size must be a positive integer")
	}
	return Pack{Size: size}, nil
}

// Pack represents a pack of a certain size. Can be extended with more fields if needed.
type Pack struct {
	Size int
}

// PackingResult represents the result of a packing calculation.
type PackingResult struct {
	PacksUsed map[int]int
}

func NewPackingResult(packsUsed map[int]int) *PackingResult {
	return &PackingResult{PacksUsed: packsUsed}
}

// PackList holds a list of available packs and ensures no duplicate sizes exist.
type PackList struct {
	Packs []Pack
}

// invariant: no duplicate pack sizes in the list
func NewPackList(packs []Pack) (*PackList, error) {
	seen := make(map[int]bool)

	for i := range packs {
		if seen[packs[i].Size] {
			return nil, errors.New("duplicate pack sizes are not allowed")
		}
		seen[packs[i].Size] = true
	}

	return &PackList{Packs: packs}, nil
}
