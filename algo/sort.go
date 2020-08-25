package algo

func mergeSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}

	var left []int
	var right []int
	for ind, elem := range items {
		if ind < len(items)/2 {
			left = append(left, elem)
		} else {
			right = append(right, elem)
		}
	}
	left = mergeSort(left)
	right = mergeSort(right)
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	for (len(left) > 0) && (len(right) > 0) {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for _, elem := range left {
		result = append(result, elem)
	}

	for _, elem := range right {
		result = append(result, elem)
	}
	return result
}

func quickSort(items []int, lo, hi int) {

	if lo < hi {
		p := partition(items, lo, hi)
		quickSort(items, lo, p-1)
		quickSort(items, lo+1, hi)
	}
}

func partition(items []int, lo, hi int) int {
	pivot := items[hi]

	i := lo
	for j := lo; j <= hi; j++ {
		if items[j] < pivot {
			items[i], items[j] = items[j], items[i]
			i++
		}
	}
	items[i], items[hi] = items[hi], items[i]
	return i
}
