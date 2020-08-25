package algo

func BinarySearch(items []int, len int, item int) int {
	left := 0
	right := len - 1

	for left <= right {
		middle := (left + right) / 2
		if items[middle] < item {
			left = middle + 1
		} else if items[middle] > item {
			right = middle - 1
		} else {
			return middle
		}
	}
	return -1
}
func BinarySearchNearest(items []int, len int, item int) int {
	if item < items[0] {
		return items[0]
	}
	if item > items[len-1] {
		return items[len-1]
	}
	left := 0
	right := len - 1

	for left <= right {
		middle := (left + right) / 2
		if items[middle] < item {
			left = middle + 1
		} else if items[middle] > item {
			right = middle - 1
		} else {
			return middle
		}
	}
	if (items[left] - item) < (item - items[right]) {
		return left
	}
	return right
}
