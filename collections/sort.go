package collections

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// SortSlice is a generic function for sorting slices in ascending order
func SortSlice[K constraints.Ordered](slice []K) []K {
	SortSliceWithCmpFunc(slice, AscOrder[K])
	return slice
}

// SortSliceWithCmpFunc is a generic function for sorting lists with customizable comparison
func SortSliceWithCmpFunc[K any](slice []K, comparisonFunc func(k1, k2 K) bool) []K {
	sort.Slice(slice, func(i, j int) bool {
		return comparisonFunc(slice[i], slice[j])
	})
	return slice
}

// AscOrder is a comparison function for sorting in ascending order
func AscOrder[K constraints.Ordered](v1, v2 K) bool {
	return v1 < v2
}

// DescOrder is a comparison function for sorting in descending order
func DescOrder[K constraints.Ordered](v1, v2 K) bool {
	return v1 > v2
}
