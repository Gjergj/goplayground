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

func Binary_search_self(data []int, searchFor int) int {

	low := 0
	high := len(data)

	for range data {
		mid := (int)(high-low) / 2
		if searchFor == data[mid] {
			return mid
		} else if searchFor < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}
