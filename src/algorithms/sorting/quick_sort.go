package sorting

func Qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivotIdx := Partition(arr, lo, hi)

	Qs(arr, lo, pivotIdx-1)
	Qs(arr, pivotIdx+1, hi)
}

func Partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	idx := lo

	// insert all the elements that are smaller than the pivot to the left
	for i := lo; i < hi; i++ {
		if arr[i] < pivot {
			arr[idx], arr[i] = arr[i], arr[idx]
			idx++
		}
	}

	arr[idx], arr[hi] = arr[hi], arr[idx]

	return idx
}
