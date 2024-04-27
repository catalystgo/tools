package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSorting(t *testing.T) {
	t.Parallel()

	t.Run("sort_slice", func(t *testing.T) {
		t.Parallel()

		ints := []int64{5, 4, 3, 2, 1}
		floats := []float64{5, 4, 3, 2, 1}
		strings := []string{"d", "c", "b", "a"}

		SortSlice(ints)
		SortSlice(floats)
		SortSlice(strings)

		require.Equal(t, []int64{1, 2, 3, 4, 5}, ints)
		require.Equal(t, []float64{1, 2, 3, 4, 5}, floats)
		require.Equal(t, []string{"a", "b", "c", "d"}, strings)
	})

	t.Run("sort_desc_order", func(t *testing.T) {
		t.Parallel()

		ints := []int64{5, 4, 3, 2, 1}
		floats := []float64{5, 4, 3, 2, 1}
		strings := []string{"d", "c", "b", "a"}

		SortSliceWithCmpFunc(ints, DescOrder[int64])
		SortSliceWithCmpFunc(floats, DescOrder[float64])
		SortSliceWithCmpFunc(strings, DescOrder[string])

		require.Equal(t, []int64{5, 4, 3, 2, 1}, ints)
		require.Equal(t, []float64{5, 4, 3, 2, 1}, floats)
		require.Equal(t, []string{"d", "c", "b", "a"}, strings)
	})

	t.Run("sort_slice_return", func(t *testing.T) {
		t.Parallel()

		ints := []int64{5, 4, 3, 2, 1}
		floats := []float64{5, 4, 3, 2, 1}
		strings := []string{"d", "c", "b", "a"}

		intsResult := SortSlice(ints)
		floatsResult := SortSlice(floats)
		stringsResult := SortSlice(strings)

		require.Equal(t, []int64{1, 2, 3, 4, 5}, ints)
		require.Equal(t, []float64{1, 2, 3, 4, 5}, floats)
		require.Equal(t, []string{"a", "b", "c", "d"}, strings)

		require.Equal(t, &ints[0], &intsResult[0])
		require.Equal(t, &floatsResult[0], &floatsResult[0])
		require.Equal(t, &stringsResult[0], &stringsResult[0])
	})

	t.Run("sort_desc_order_return", func(t *testing.T) {
		t.Parallel()

		ints := []int64{5, 4, 3, 2, 1}
		floats := []float64{5, 4, 3, 2, 1}
		strings := []string{"d", "c", "b", "a"}

		intsResult := SortSliceWithCmpFunc(ints, DescOrder[int64])
		floatsResult := SortSliceWithCmpFunc(floats, DescOrder[float64])
		stringsResult := SortSliceWithCmpFunc(strings, DescOrder[string])

		require.Equal(t, []int64{5, 4, 3, 2, 1}, ints)
		require.Equal(t, []float64{5, 4, 3, 2, 1}, floats)
		require.Equal(t, []string{"d", "c", "b", "a"}, strings)
		require.Equal(t, &ints[0], &intsResult[0])
		require.Equal(t, &floatsResult[0], &floatsResult[0])
		require.Equal(t, &stringsResult[0], &stringsResult[0])
	})
}
