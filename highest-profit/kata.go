package kata

import (
	"sort"
)

func MinMax(arr []int) [2]int {
	sort.Ints(arr)
	var minmax [2]int
	minmax[0] = arr[0]
	minmax[1] = arr[len(arr)-1]

	return minmax
}
