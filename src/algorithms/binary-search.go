package algorithms

import "slices"

// the arr needs to be sorted
func BinarySearch(arr []int, element int) int {
	slices.Sort(arr)

	// left is inclusive and right is exclusive
	// [l,r)

	l := 0
	r := len(arr)

	for l < r {
		mid := l + (r-l)/2

		if arr[mid] == element {
			return mid
		}

		// element is on the second half
		if element > arr[mid] {
			l = mid + 1
		} else {
			// element is on the first half
			r = mid // since r is exclusive, we don't need to add -1
		}
	}

	return -1
}
