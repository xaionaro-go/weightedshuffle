package weightedrandsort

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

var reoderFuncs = []funcInfo{
	{"Reorder", Reorder},
	{"ReorderInplace", ReorderInplace},
}

func TestReoder(t *testing.T) {
	for _, reorderFunc := range reoderFuncs {
		t.Run(reorderFunc.Name, func(t *testing.T) {
			amountOfItems := 10000
			s := make([]float64, amountOfItems)
			for i := 0; i < amountOfItems; i++ {
				s[i] = float64(i)
			}

			reorderFunc.Func(s, func(i int) float64 {
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

func BenchmarkReoder(b *testing.B) {
	r := rand.NewSource(0)

	for _, amountOfItems := range []int{0, 1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("%d", amountOfItems), func(b *testing.B) {
			s := make([]float64, amountOfItems)
			for i := 0; i < amountOfItems; i++ {
				s[i] = float64(i)
			}

			for _, reorderFunc := range reoderFuncs {
				if reorderFunc.Name == "ReorderInplace" && amountOfItems > 10000 {
					b.Skip()
					return
				}
				b.Run(reorderFunc.Name, func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for idx := 0; idx < b.N; idx++ {
						reorderFunc.Func(s, func(i int) float64 {
							return s[i]
						}, r)
					}
				})
			}
		})
	}
}
