package collections

// MapSlice is a generic function that takes a slice of elements of type X
// and transforms it into a slice of elements of type Y by applying
// the function f to each element
func MapSlice[X any, Y any](input []X, f func(X) Y) []Y {
	result := make([]Y, len(input))
	for idx, element := range input {
		result[idx] = f(element)
	}
	return result
}

// MapSliceWithIndex is a generic function that takes a slice of elements of type X
// and transforms it into a slice of elements of type Y by applying
// the function f to each element, including the index
func MapSliceWithIndex[X any, Y any](input []X, f func(int, X) Y) []Y {
	result := make([]Y, len(input))
	for idx, element := range input {
		result[idx] = f(idx, element)
	}
	return result
}

// MapAndFilterSliceWithIndex is a generic function for filtering and transforming elements of a slice
func MapAndFilterSliceWithIndex[X any, Y any](input []X, f func(int, X) Y, filter func(int, X) bool) []Y {
	var result []Y
	for idx, element := range input {
		if filter(idx, element) {
			result = append(result, f(idx, element))
		}
	}
	return result
}

// CreateSliceFromDict creates a slice from a hash table
func CreateSliceFromDict[K comparable, X any, Y any](input map[K]X, f func(K, X) Y) []Y {
	result := make([]Y, len(input))
	idx := 0
	for key := range input {
		result[idx] = f(key, input[key])
		idx++
	}
	return result
}

// CreateDictFromSlice is a generic function for creating a hash table from a flat list
func CreateDictFromSlice[K comparable, X any, Y any](input []X, keyFunc func(*X) K, valueFunc func(*X) Y) map[K]Y {
	result := make(map[K]Y, len(input))
	for idx := range input {
		result[keyFunc(&input[idx])] = valueFunc(&input[idx])
	}
	return result
}

// GroupByToDict is a generic function for grouping data into a hash table
func GroupByToDict[K comparable, X any, Y any](input []X, keyFunc func(*X) K, valueFunc func(*X) Y) map[K][]Y {
	result := make(map[K][]Y, len(input))

	for idx := range input {
		key := keyFunc(&input[idx])
		value := valueFunc(&input[idx])
		if keyValue, exists := result[key]; exists {
			result[key] = append(keyValue, value)
		} else {
			result[key] = []Y{value}
		}
	}
	return result
}

// ForEachSlice is a generic function for modifying elements of a slice
func ForEachSlice[X any](input []X, f func(int, *X) error) error {
	for idx := range input {
		if err := f(idx, &input[idx]); err != nil {
			return err
		}
	}
	return nil
}

// GetSymmetricallyDifferentKeys returns the symmetric difference of keys of hash tables
// Unique keys of each hash table are returned separately
func GetSymmetricallyDifferentKeys[K comparable, X any](a, b map[K]X) ([]K, []K) {
	var left, right []K
	for key := range a {
		if _, ok := b[key]; !ok {
			left = append(left, key)
		}
	}
	for key := range b {
		if _, ok := a[key]; !ok {
			right = append(right, key)
		}
	}

	return left, right
}
