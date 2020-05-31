package sorting_test

import (
	"github.com/ameniGa/go-algorithms/sorting/insertion"
	"math/rand"
	"testing"
	"time"
)

type algo struct {
	name string
	f    func([]int) []int
}

func getAlgorithms() []algo {
	return []algo{
		{
			name: "insertion",
			f:    insertion.Sort,
		},
	}
}

func generateRandomSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func Benchmark_Sort(b *testing.B) {
	for _, sort := range getAlgorithms() {
		b.Run(sort.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
                sort.f(generateRandomSlice(2000))
			}
		})
	}
}