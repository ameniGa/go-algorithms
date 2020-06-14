package sorting_test

import (
	"github.com/ameniGa/go-algorithms/sorting"
	"math/rand"
	"testing"
	"time"
)

type algo struct {
	name sorting.Algorithm
	f    func(algo sorting.Algorithm, list *sorting.Sortable) ([]interface{}, error)
}

func getAlgorithms() []algo {
	return []algo{
		{
			name: sorting.Insertion,
			f:    sorting.Sort,
		},
	}
}

func generateRandomList(size int) *sorting.Sortable {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	sortable := sorting.Sortable{
		List: slice,
		Compare: func(i, j interface{}) bool {
			return i.(int) > j.(int)
		},
	}
	return &sortable
}

func Benchmark_Sort(b *testing.B) {
	for _, sort := range getAlgorithms() {
		b.Run(string(sort.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.f(sort.name, generateRandomList(2000))
			}
		})
	}
}