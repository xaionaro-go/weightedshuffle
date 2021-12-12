package weightedshuffle

import (
	"math/rand"
	"reflect"
	"sort"
)

type randWeightSortWrapper struct {
	slice   interface{}
	swap    func(i, j int)
	length  int
	weights []float64
}

func (s *randWeightSortWrapper) Len() int {
	return s.length
}

func (s *randWeightSortWrapper) Less(i, j int) bool {
	return s.weights[i] > s.weights[j]
}

func (s *randWeightSortWrapper) Swap(i, j int) {
	s.swap(i, j)
	s.weights[i], s.weights[j] = s.weights[j], s.weights[i]
}

func newRandWeightSortWrapper(
	slice interface{},
	weight func(i int) float64,
	randSource rand.Source,
) (*randWeightSortWrapper, interface{}) {
	fullLength := reflect.ValueOf(slice).Len()
	swap := reflect.Swapper(slice)
	var randFloat64 func() float64
	if randSource != nil {
		randFloat64 = rand.New(randSource).Float64
	} else {
		randFloat64 = rand.Float64
	}

	length := fullLength
	var weightByIndex []float64
	for idx := 0; idx < length; idx++ {
		w := weight(idx)
		if w == 0 {
			swap(idx, length-1)
			length--
			idx--
			continue
		}
		weightByIndex = append(weightByIndex, w*randFloat64())
	}

	var unweightedSlice interface{}
	if length != fullLength {
		unweightedSlice = reflect.ValueOf(slice).Slice(length, fullLength).Interface()
	}

	return &randWeightSortWrapper{
		slice:   slice,
		swap:    swap,
		length:  length,
		weights: weightByIndex,
	}, unweightedSlice
}

// Shuffle randomly reorders the slice with the preference to put first items with
// higher weight.
//
// Basically, instead of this function: you may just calculate randomized
// weights and use sort.Slice (the time and space complexity will be the same).
//
// Out-of-order factor of this method is 25% (where 0% means no randomness
// and just sort by weight, 50% means purely random order and 100% means
// the reverse order).
// The "out-of-order factor" is the share of pairs which are in wrong order,
// relatively to each other when weights are just a sequence [ 0 .. len(slice) ).
//
// T: O(n * log(n))
// S: O(n)
func Shuffle(
	slice interface{},
	weight func(i int) float64,
	randSource rand.Source,
) {
	weightedSliceWrapper, unweightedSlice := newRandWeightSortWrapper(slice, weight, randSource)
	sort.Sort(weightedSliceWrapper)
	if unweightedSlice != nil {
		rand.Shuffle(reflect.ValueOf(unweightedSlice).Len(), reflect.Swapper(unweightedSlice))
	}
}
