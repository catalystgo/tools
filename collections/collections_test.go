package collections

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Struct1 struct {
	Key  int64
	Name string
}

type Struct2 struct {
	Value string
}

func TestHelpers(t *testing.T) {
	t.Parallel()

	t.Run("map_slice", func(t *testing.T) {
		t.Parallel()

		input := []Struct1{
			{
				Key:  1,
				Name: "name_1",
			},
			{
				Key:  2,
				Name: "name_2",
			},
		}

		result := MapSlice(input, func(value Struct1) Struct2 {
			return Struct2{
				Value: value.Name,
			}
		})

		expectedResult := []Struct2{
			{
				Value: "name_1",
			},
			{
				Value: "name_2",
			},
		}

		require.Equal(t, result, expectedResult)

	})

	t.Run("map_slice_with_index", func(t *testing.T) {
		t.Parallel()

		input := []Struct1{
			{
				Key:  1,
				Name: "name_1",
			},
			{
				Key:  2,
				Name: "name_2",
			},
		}

		result := MapSliceWithIndex(input, func(idx int, value Struct1) int {
			return idx
		})

		expectedResult := []int{
			0, 1,
		}

		require.Equal(t, result, expectedResult)
	})

	t.Run("map_and_filter_slice_with_index", func(t *testing.T) {
		t.Parallel()

		input := []Struct1{
			{
				Key:  1,
				Name: "name_1",
			},
			{
				Key:  2,
				Name: "name_2",
			},
			{
				Key:  3,
				Name: "name_3",
			},
			{
				Key:  4,
				Name: "name_4",
			},
			{
				Key:  5,
				Name: "name_5",
			},
			{
				Key:  6,
				Name: "name_6",
			},
			{
				Key:  7,
				Name: "name_7",
			},
		}

		result := MapAndFilterSliceWithIndex(
			input,
			func(idx int, value Struct1) int {
				return idx
			},
			func(idx int, value Struct1) bool {
				return idx%2 == 0 || value.Name == "name_6"
			})

		expectedResult := []int{
			0, 2, 4, 5, 6,
		}

		require.Equal(t, expectedResult, result)

	})

	t.Run("create_map_from_slice", func(t *testing.T) {
		t.Parallel()

		input := []Struct1{
			{
				Key:  1,
				Name: "name_1",
			},
			{
				Key:  2,
				Name: "name_2",
			},
		}

		result := CreateDictFromSlice(
			input,
			func(value *Struct1) int64 { return value.Key },
			func(value *Struct1) Struct2 {
				return Struct2{
					Value: value.Name,
				}
			},
		)

		expectedResult := map[int64]Struct2{
			1: {
				Value: "name_1",
			},
			2: {
				Value: "name_2",
			},
		}

		require.Equal(t, result, expectedResult)

	})

	t.Run("group_by", func(t *testing.T) {
		t.Parallel()

		input := []Struct1{
			{
				Key:  1,
				Name: "name_1",
			},
			{
				Key:  2,
				Name: "name_2",
			},
			{
				Key:  1,
				Name: "name_3",
			},
			{
				Key:  4,
				Name: "name_4",
			},
			{
				Key:  2,
				Name: "name_5",
			},
		}

		result := GroupByToDict(
			input,
			func(value *Struct1) int64 { return value.Key },
			func(value *Struct1) string {
				return value.Name
			},
		)

		expectedResult := map[int64][]string{
			1: {"name_1", "name_3"},
			2: {"name_2", "name_5"},
			4: {"name_4"},
		}

		require.Equal(t, result, expectedResult)
	})

	t.Run("for_each", func(t *testing.T) {
		t.Parallel()

		input := []Struct1{
			{
				Key:  1,
				Name: "name_1",
			},
			{
				Key:  2,
				Name: "name_2",
			},
		}

		err := ForEachSlice(
			input,
			func(idx int, value *Struct1) error {
				value.Name = fmt.Sprintf("modified_%s", value.Name)
				return nil
			},
		)

		require.NoError(t, err)

		expectedResult := []Struct1{
			{
				Key:  1,
				Name: "modified_name_1",
			},
			{
				Key:  2,
				Name: "modified_name_2",
			},
		}

		require.Equal(t, input, expectedResult)
	})

	t.Run("error/for_each", func(t *testing.T) {
		t.Parallel()

		input := []Struct1{
			{
				Key:  1,
				Name: "name_1",
			},
			{
				Key:  2,
				Name: "name_2",
			},
		}
		err := ForEachSlice(
			input,
			func(idx int, value *Struct1) error {
				return errors.New("some error")
			},
		)

		require.EqualError(t, err, errors.New("some error").Error())

		expectedResult := []Struct1{
			{
				Key:  1,
				Name: "name_1",
			},
			{
				Key:  2,
				Name: "name_2",
			},
		}

		require.Equal(t, input, expectedResult)
	})

	t.Run("create_slice_from_map", func(t *testing.T) {
		t.Parallel()

		input := map[int64]Struct1{
			1: {
				Key:  1,
				Name: "name_1",
			},
			2: {
				Key:  2,
				Name: "name_2",
			},
			3: {
				Key:  3,
				Name: "name_3",
			},
		}

		result := CreateSliceFromDict(input, func(index int64, value Struct1) Struct2 {
			return Struct2{
				Value: value.Name,
			}
		})

		require.Equal(t, len(result), 3)
	})

	t.Run("get_excepted_keys", func(t *testing.T) {
		t.Parallel()

		input1 := map[int64]bool{
			1: true,
			2: true,
			3: true,
			4: true,
		}

		input2 := map[int64]bool{
			3: true,
			4: true,
			5: true,
			6: true,
			7: true,
		}

		left, right := GetSymmetricallyDifferentKeys(input1, input2)
		SortSlice(left)
		SortSlice(right)

		require.Equal(t, []int64{1, 2}, left)
		require.Equal(t, []int64{5, 6, 7}, right)
	})

	t.Run("split_into_batches", func(t *testing.T) {
		t.Parallel()

		testsCases := []struct {
			name      string
			input     []int
			batchSize int
			expected  [][]int
		}{
			{
				name:      "empty",
				input:     []int{},
				batchSize: 2,
				expected:  [][]int{},
			},
			{
				name:      "single_batch",
				input:     []int{1, 2, 3, 4, 5},
				batchSize: 5,
				expected:  [][]int{{1, 2, 3, 4, 5}},
			},
			{
				name:      "multiple_batches_less_than_batch_size",
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				batchSize: 3,
				expected:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}},
			},
			{
				name:      "multiple_batches_equal_to_batch_size",
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				batchSize: 2,
				expected:  [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			},
			{
				name:      "multiple_batches_greater_than_batch_size",
				input:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				batchSize: 4,
				expected:  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10}},
			},
		}

		for _, tt := range testsCases {
			t.Run(tt.name, func(t *testing.T) {
				result := SplitIntoBatches(tt.input, tt.batchSize)
				require.Equal(t, tt.expected, result)
			})
		}
	})
}
