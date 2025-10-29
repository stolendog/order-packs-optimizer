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
