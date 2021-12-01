package ConcurrencySort

import (
	"sort"
	"sync"
)
func mergeSlice(first, second []int) []int {
	result := make([]int, len(first) + len(second))
	i, j, k := 0, 0, 0
	for i < len(first) && j < len(second) {
		if first[i] <= second[j] {
			result[k] = first[i]
			i++
		} else {
			result[k] = second[j]
			j++
		}
		k++
	}
	for i < len(first) {
		result[k] = first[i]
		i++
		k++
	}
	for j < len(second) {
		result[k] = second[j]
		j++
		k++
	}
	return result
}
func mergeAlgorithm(slices [][]int) []int  {
	var step = 1
	for step < len(slices) {
		for i := 0; i + step < len(slices); i += 2 * step {
			slices[i] = mergeSlice(slices[i], slices[i + step])
		}
		step *= 2
	}
	return slices[0]
}
func Sort(slice *[]int, threadN int) {
	var wg sync.WaitGroup
	miniSlices := make([][]int, threadN)
	for i := 0; i < threadN; i++ {
		i := i
		left := (len(*slice) * i) / threadN
		right := (len(*slice) * (i + 1)) / threadN
		miniSlices[i] = (*slice)[left:right]
		wg.Add(1)
		go func() {
			sort.Ints(miniSlices[i])
			wg.Done()
		}()
	}
	wg.Wait()
	*slice = mergeAlgorithm(miniSlices)
}




























