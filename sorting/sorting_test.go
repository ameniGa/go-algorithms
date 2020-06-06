package sorting_test

import (
	"github.com/ameniGa/go-algorithms/sorting/insertion"
	"math/rand"
	"testing"
	"time"
)

type algo struct {
	name string
	f    func(list insertion.Sortable)
}

func getAlgorithms() []algo {
	return []algo{
		{
			name: "insertion",
			f:    insertion.Sort,
		},
	}
}

func generateRandomList(size int) insertion.Sortable {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	mockSortable := insertion.SortableMock{
		GetLen: func() int {
			return len(slice)
		},
		GetCompare: func(i, j interface{}) bool {
			return i.(int) > j.(int)
		},
		GetTranspose: func(i int, j interface{}) {
			slice[i] = j.(int)
		},
		GetElement: func(i int) interface{} {
			return slice[i]
		},
	}
	return mockSortable
}

func Benchmark_Sort(b *testing.B) {
	for _, sort := range getAlgorithms() {
		b.Run(sort.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
                sort.f(generateRandomList(2000))
			}
		})
	}
}