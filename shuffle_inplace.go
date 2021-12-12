package weightedshuffle

import (
	"fmt"
	"math/rand"
	"reflect"
)

// ShuffleInplace randomly reorders the slice with the preference to put first items with
// higher weight.
//
// The first played out position is 0th and the probability to take the position is
// proportional to the item's weight. Then 1th position is played out, and so on
// up until the end.
//
// Out-of-order factor of this method is 31% (where 0% means no randomness
// and just sort by weight, 50% means purely random order and 100% means
// the reverse order).
// The "out-of-order factor" is the share of pairs which are in wrong order,
// relatively to each other when weights are just a sequence [ 0 .. len(slice) ).
//
// T: O(n^2)
// S: O(1)
//
// It is not recommended to use this method if there are more than 100 items in the slice.
// Use method Reorder, instead.
func ShuffleInplace(
	slice interface{},
	weight func(i int) float64,
	randSource rand.Source,
) {
	length := reflect.ValueOf(slice).Len()
	swap := reflect.Swapper(slice)
	if length <= 1 {
		return
	}
	var randFloat64 func() float64
	var randShuffle func(n int, swap func(i, j int))
	if randSource != nil {
		randFloat64 = rand.New(randSource).Float64
		randShuffle = rand.New(randSource).Shuffle
	} else {
		randFloat64 = rand.Float64
		randShuffle = rand.Shuffle
	}

	weightSum := float64(0)
	for idx := 0; idx < length; idx++ {
		w := weight(idx)
		if w < 0 {
			panic(fmt.Errorf("negative weight at index %d", idx))
		}
		weightSum += w
	}

	for baseIdx := 0; baseIdx < length; baseIdx++ {
		randWeightSum := randFloat64() * weightSum
		randWeightSumOrig := randWeightSum
		for swapIdx := baseIdx; swapIdx < length; swapIdx++ {
			w := weight(swapIdx)
			randWeightSum -= w
			if randWeightSum < 0 {
				swap(baseIdx, swapIdx)
				weightSum -= w
				if weightSum < 0 {
					panic(fmt.Errorf("internal error: %f < 0", weightSum))
				}
				break
			}
		}
		if randWeightSum == randWeightSumOrig {
			// only zero-weight items left, just use shuffle:
			randShuffle(length-baseIdx, func(i, j int) {
				swap(baseIdx+i, baseIdx+j)
			})
			break
		}
	}
}
