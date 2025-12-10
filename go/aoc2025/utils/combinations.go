package utils

func Combinations[T any](items []T, k int) [][]T {
	if k <= 0 || k > len(items) {
		return [][]T{}
	}

	result := [][]T{}
	combination := make([]T, k)
	indices := make([]int, k)

	for i := range indices {
		indices[i] = i
	}

	for {
		for i, idx := range indices {
			combination[i] = items[idx]
		}
		combCopy := make([]T, k)
		copy(combCopy, combination)
		result = append(result, combCopy)

		i := k - 1
		for i >= 0 && indices[i] == len(items)-k+i {
			i--
		}

		if i < 0 {
			break
		}

		indices[i]++
		for j := i + 1; j < k; j++ {
			indices[j] = indices[j-1] + 1
		}
	}

	return result
}

func Subsets[T any](items []T) [][]T {
	result := [][]T{}
	for size := 1; size <= len(items); size++ {
		result = append(result, Combinations(items, size)...)
	}
	return result
}
