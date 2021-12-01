package ConcurrencySort

import (
	"math/rand"
	"sort"
	"testing"
)
const (
	sz = 100000
)
func BenchmarkSortFromGolangLibrary(b *testing.B) {
	slice := make([]int, sz)
	for i := 0; i < sz; i++ {
		slice[i] = rand.Int()
	}
	sort.Ints(slice)
}
func BenchmarkWithOneWorker(b *testing.B) {
	slice := make([]int, sz)
	for i := 0; i < sz; i++ {
		slice[i] = rand.Int()
	}
	Sort(&slice, 1)
}

func BenchmarkWithTwoWorker(b *testing.B) {
	slice := make([]int, sz)
	for i := 0; i < sz; i++ {
		slice[i] = rand.Int()
	}
	Sort(&slice, 2)
}

func BenchmarkWithFourWorker(b *testing.B) {
	slice := make([]int, sz)
	for i := 0; i < sz; i++ {
		slice[i] = rand.Int()
	}
	Sort(&slice, 4)
}

func BenchmarkWithEightWorker(b *testing.B) {
	slice := make([]int, sz)
	for i := 0; i < sz; i++ {
		slice[i] = rand.Int()
	}
	Sort(&slice, 8)
}