package top

import "sort"

type SPair struct {
	char string
	count int
}
func TopWords(s string, n int) []SPair {
	// find the max of char in word and create slice with size max_of_char
	var mx = 0
	for _, ch := range s {
		if int(ch) > mx {
			mx = int(ch)
		}
	}
	slice_map := make([]SPair, mx + 1)
	for _, ch := range s {
		slice_map[ch].char = string(ch)
		slice_map[ch].count++
	}

	// firstly sort by counters descending, secondly sort by char ascending
	sort.Slice(slice_map, func(i, j int) bool {
		if slice_map[i].count == slice_map[j].count {
			return slice_map[i].char < slice_map[j].char
		}
		return slice_map[i].count > slice_map[j].count
	})
	for index, value := range slice_map {
		if value.count == 0 || index == n {
			return slice_map[:index]
		}
	}
	return slice_map
}
