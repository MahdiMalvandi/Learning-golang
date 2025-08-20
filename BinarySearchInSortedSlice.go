package main
func BinarySearch(slice []int, target int) int {
	frs := 0               // first index
	lst := len(slice) - 1  // last index
	mid := (frs + lst) / 2 // mid index

	for frs <= lst {
		if slice[mid] == target {
			return mid
		} else if slice[mid] > target {
			lst = mid - 1
		} else if slice[mid] < target {
			frs = mid + 1
		}
		mid = (frs + lst) / 2
	}
	return -1
}

