package calculator

import (
	"reflect"
	"testing"
)

func TestDPCalculator_InputValidation(t *testing.T) {
	testCases := []struct {
		name           string
		orderQuantity  int
		availablePacks []int
		expectError    bool
	}{
		{
			name:           "Negative order quantity",
			orderQuantity:  -5,
			availablePacks: []int{100, 200},
			expectError:    true,
		},
		{
			name:           "Zero order quantity",
			orderQuantity:  0,
			availablePacks: []int{100, 200},
			expectError:    true,
		},
		{
			name:           "Empty available packs",
			orderQuantity:  100,
			availablePacks: []int{},
			expectError:    true,
		},
		{
			name:           "Pack size zero",
			orderQuantity:  100,
			availablePacks: []int{0, 200},
			expectError:    true,
		},
		{
			name:           "Pack size negative",
			orderQuantity:  100,
			availablePacks: []int{-50, 200},
			expectError:    true,
		},
	}

	calculator := NewPackDPCalculator()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := calculator.Calculate(tc.orderQuantity, tc.availablePacks)
			if tc.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("did not expect error but got: %v", err)
			}
		})
	}
}

func TestDPCalculator_BasicCases(t *testing.T) {
	testCases := []struct {
		name           string
		orderQuantity  int
		availablePacks []int
		expected       map[int]int
	}{
		{
			name:           "Exact match with single pack size",
			orderQuantity:  200,
			availablePacks: []int{200},
			expected:       map[int]int{200: 1},
		},
		{
			name:           "Overshoot with single pack size",
			orderQuantity:  150,
			availablePacks: []int{200},
			expected:       map[int]int{200: 1},
		},
		{
			name:           "Multiple pack sizes optimal solution",
			orderQuantity:  300,
			availablePacks: []int{100, 200},
			expected:       map[int]int{200: 1, 100: 1},
		},
		{
			name:           "Order quantity of 1",
			orderQuantity:  1,
			availablePacks: []int{1, 5, 10},
			expected:       map[int]int{1: 1},
		},
		{
			name:           "Order quantity of 1 with larger packs",
			orderQuantity:  1,
			availablePacks: []int{5, 10, 20},
			expected:       map[int]int{5: 1},
		},
	}

	calculator := NewPackDPCalculator()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := calculator.Calculate(tc.orderQuantity, tc.availablePacks)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !mapsEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestDPCalculator_MinimizeTotalItems(t *testing.T) {
	testCases := []struct {
		name           string
		orderQuantity  int
		availablePacks []int
		expected       map[int]int
	}{
		{
			name:           "Minimize total items",
			orderQuantity:  251,
			availablePacks: []int{250, 500, 1000},
			expected:       map[int]int{500: 1},
		},
		{
			name:           "Rule 2 over Rule 3 - fewer items wins",
			orderQuantity:  501,
			availablePacks: []int{250, 500, 1000},
			expected:       map[int]int{250: 1, 500: 1}, // 750 total items vs 1000 with single 1000 pack
		},
	}

	calculator := NewPackDPCalculator()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := calculator.Calculate(tc.orderQuantity, tc.availablePacks)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !mapsEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}

}

func TestDPCalculator_EdgeCases(t *testing.T) {
	testCases := []struct {
		name           string
		orderQuantity  int
		availablePacks []int
		expected       map[int]int
	}{
		{
			name:           "Edge case large order quantity",
			orderQuantity:  500000,
			availablePacks: []int{23, 31, 53},
			expected:       map[int]int{23: 2, 31: 7, 53: 9429},
		},
		{
			name:           "Duplicate pack sizes",
			orderQuantity:  100,
			availablePacks: []int{100, 100, 200, 200},
			expected:       map[int]int{100: 1},
		},
	}
	calculator := NewPackDPCalculator()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := calculator.Calculate(tc.orderQuantity, tc.availablePacks)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !mapsEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func mapsEqual(a, b map[int]int) bool {
	return reflect.DeepEqual(a, b)
}
