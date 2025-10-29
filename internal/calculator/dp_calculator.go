package calculator

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

// 1. Only whole packs can be sent. Packs cannot be broken open.
// 2. Within the constraints of Rule 1 above, send out the least amount of items to fulfil the order.
// 3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each
// order.
// (Please note, rule #2 takes precedence over rule #3)

// Dynamic Programming based order packs calculator
// at the end reconstructs the optimal solution based on the DP table (minPacks vs lastPackUsed)
type DPCalculator struct{}

func NewPackDPCalculator() *DPCalculator {
	return &DPCalculator{}
}

func (pc *DPCalculator) Calculate(orderQuantity int, availablePacks []int) (map[int]int, error) {
	if orderQuantity <= 0 {
		return nil, errors.New("order quantity must be a positive integer")
	}

	if len(availablePacks) == 0 {
		return nil, errors.New("no available packs provided")
	}

	// Get unique, sorted pack sizes. Sorting largest to smallest can be a minor optimization.
	packSizeSet := make(map[int]bool)
	for _, packSize := range availablePacks {
		if packSize <= 0 {
			return nil, fmt.Errorf("pack size must be a positive integer: %v", packSize)
		}
		packSizeSet[packSize] = true
	}
	packSizes := make([]int, 0, len(packSizeSet))
	for size := range packSizeSet {
		packSizes = append(packSizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes))) // Largest to smallest, can be beneficial to take larger packs first

	// determine an appropriate limit for DP table
	maxPackSize := packSizes[0]          // largest pack size
	limit := orderQuantity + maxPackSize // limit set to orderQuantity + largest pack size = to have space for more items than orderQuantity when there is no exact match

	// using DP to find optimal solution (based on coin change problem)
	// minPacks[i] will store the minimum number of packs to achieve a total of i items
	minPacks := make([]int, limit+1)

	// lastPackUsed[i] will store the pack size used to achieve the optimal solution for total i
	// its needed to reconstruct the final result
	lastPackUsed := make([]int, limit+1)

	// init with infinity values
	for i := 1; i <= limit; i++ {
		minPacks[i] = math.MaxInt32
	}
	// 0 items requires 0 packs.
	minPacks[0] = 0

	// for each possible total number of items i
	// if pack sizes can be used to form that total, update minPacks and lastPackUsed
	for i := 1; i <= limit; i++ {
		// try to form it by adding each available pack size (from bigger to smaller)
		for _, size := range packSizes {
			if size <= i && minPacks[i-size] != math.MaxInt32 {
				// min(current minPacks[i], minPacks[i - size] + 1)
				if minPacks[i-size]+1 < minPacks[i] {
					minPacks[i] = minPacks[i-size] + 1
					lastPackUsed[i] = size
				}
			}
		}
	}

	// Rule #2: find the least amount of items >= orderQuantity.
	// iterate from the order quantity upwards and the first solvable total is our best total
	bestTotal := -1
	for i := orderQuantity; i <= limit; i++ {
		if minPacks[i] != math.MaxInt32 {
			bestTotal = i
			break // found the smallest total that meets the order
		}
	}

	if bestTotal == -1 {
		return nil, errors.New("cannot fulfill order with the given pack sizes")
	}

	// reconstruct the packs used from lastPackUsed

	packsUsed := make(map[int]int)
	currentTotal := bestTotal
	for currentTotal > 0 {
		packSize := lastPackUsed[currentTotal]
		packsUsed[packSize]++
		currentTotal -= packSize
	}

	return packsUsed, nil
}
