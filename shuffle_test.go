package weightedshuffle

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type funcInfo struct {
	Name string
	Func func(
		slice interface{},
		weight func(i int) float64,
		randSource rand.Source,
	)
}

var shuffleFuncs = []funcInfo{
	{"Shuffle", Shuffle},
	{"ShuffleInplace", ShuffleInplace},
}

func TestReoderRandomness(t *testing.T) {
	for _, shuffleFunc := range shuffleFuncs {
		t.Run(shuffleFunc.Name, func(t *testing.T) {
			amountOfItems := 10000
			s := make([]float64, amountOfItems)
			for i := 0; i < amountOfItems; i++ {
				s[i] = float64(i)
			}

			shuffleFunc.Func(s, func(i int) float64 {
				return s[i]
			}, rand.NewSource(0))

			assert.Equal(t, float64(0), s[amountOfItems-1], s)

			countMap := map[float64]int{}
			for _, v := range s {
				countMap[v]++
				assert.Equal(t, 1, countMap[v])
			}
			assert.Len(t, countMap, amountOfItems)

			leftSum, rightSum := float64(0), float64(0)
			for idx, v := range s {
				if idx < amountOfItems/2 {
					leftSum += v
				} else {
					rightSum += v
				}
			}
			assert.InDelta(t, 1.8, leftSum/rightSum, 0.2, s)

			outOfOrderPairCount := 0
			for idx0, v0 := range s {
				for idx1, v1 := range s {
					if idx0 == idx1 {
						continue
					}

					if (idx0 < idx1) == (v0 < v1) {
						outOfOrderPairCount++
					}
				}
			}

			theoreticalMaxOutOfOrderPairCount := amountOfItems * (amountOfItems - 1)
			assert.InDelta(t, 0.27, float64(outOfOrderPairCount)/float64(theoreticalMaxOutOfOrderPairCount), 0.04, s)
		})
	}
}

func TestReoderZeroWeight(t *testing.T) {
	for _, shuffleFunc := range shuffleFuncs {
		t.Run(shuffleFunc.Name, func(t *testing.T) {
			amountOfItems := 1000
			s0 := make([]float64, amountOfItems)
			s1 := make([]float64, amountOfItems)
			for i := 0; i < amountOfItems; i++ {
				if i%2 == 0 {
					s0[i] = float64(i) + 1
					s1[i] = float64(i) + 1
				} else {
					s0[i] = -float64(i)
					s1[i] = -float64(i)
				}
			}

			shuffleFunc.Func(s0, func(i int) float64 {
				if s0[i] < 0 {
					return 0
				}
				return s0[i]
			}, rand.NewSource(0))
			shuffleFunc.Func(s1, func(i int) float64 {
				if s1[i] < 0 {
					return 0
				}
				return s1[i]
			}, rand.NewSource(int64(amountOfItems)))

			for idx, v := range s0 {
				if idx < amountOfItems/2 {
					assert.Positive(t, v, fmt.Sprintf("%d", idx))
				} else {
					assert.Negative(t, v, fmt.Sprintf("%d", idx))
				}
			}

			greaterCount := 0
			for idx, v0 := range s0[amountOfItems/2:] {
				v1 := s1[amountOfItems/2+idx]
				if v1 > v0 {
					greaterCount++
				}
			}

			assert.InDelta(t, 0.5, float64(greaterCount)/float64(amountOfItems/2), 0.05, fmt.Sprintf("\n%v\n%v\n", s0, s1))
		})
	}
}

func BenchmarkShuffle(b *testing.B) {
	r := rand.NewSource(0)

	for _, amountOfItems := range []int{0, 1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("%d", amountOfItems), func(b *testing.B) {
			s := make([]float64, amountOfItems)
			for i := 0; i < amountOfItems; i++ {
				s[i] = float64(i)
			}

			for _, shuffleFunc := range shuffleFuncs {
				if shuffleFunc.Name == "ShuffleInplace" && amountOfItems > 10000 {
					b.Skip()
					return
				}
				b.Run(shuffleFunc.Name, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for idx := 0; idx < b.N; idx++ {
						shuffleFunc.Func(s, func(i int) float64 {
							return s[i]
						}, r)
					}
				})
			}
		})
	}
}
