package domain

import "testing"

func TestPack(t *testing.T) {
	t.Run("Valid Pack Creation", func(t *testing.T) {
		pack, err := NewPack(500)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if pack.Size != 500 {
			t.Errorf("expected pack size 500, got %v", pack.Size)
		}
	})

	t.Run("Invalid Pack Creation - Zero Size", func(t *testing.T) {
		_, err := NewPack(0)
		if err == nil {
			t.Fatalf("expected error for zero size, got none")
		}
	})

	t.Run("Invalid Pack Creation - Negative Size", func(t *testing.T) {
		_, err := NewPack(-100)
		if err == nil {
			t.Fatalf("expected error for negative size, got none")
		}
	})
}

func TestPackList(t *testing.T) {
	t.Run("Valid PackList Creation", func(t *testing.T) {
		packs := []Pack{
			{Size: 100},
			{Size: 200},
			{Size: 500},
		}
		packList, err := NewPackList(packs)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(packList.Packs) != 3 {
			t.Errorf("expected 3 packs in the list, got %v", len(packList.Packs))
		}
	})

	t.Run("Invalid PackList Creation - Duplicate Sizes", func(t *testing.T) {
		packs := []Pack{
			{Size: 100},
			{Size: 200},
			{Size: 100}, // duplicate
		}
		_, err := NewPackList(packs)
		if err == nil {
			t.Fatalf("expected error for duplicate pack sizes, got none")
		}
	})
}
