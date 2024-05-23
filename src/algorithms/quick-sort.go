package algorithms

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivot := Partition(arr, lo, hi)
	qs(arr, lo, pivot-1)
	qs(arr, pivot+1, hi)
}

func Partition(arr []int, lo, hi int) int {
	pivot := arr[hi]

	idx := lo - 1

	for i := lo; i < hi; i++ {
		// insert all values smaller or equal to pivot
		// before the pivot
		if arr[i] <= pivot {
			idx++

			tmp := arr[idx]
			arr[idx] = arr[i]
			arr[i] = tmp
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

func QuickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}
