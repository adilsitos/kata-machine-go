package sorting

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// split the array into halves - log n operation
	arr1 := MergeSort(arr[:len(arr)/2])
	arr2 := MergeSort(arr[len(arr)/2:])
	return Merge(arr1, arr2)

}

func Merge(arr1, arr2 []int) []int {
	newArr := []int{}
	i := 0
	j := 0

	// go over all the elements into the arrays
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			newArr = append(newArr, arr1[i])
			i++
		} else {
			newArr = append(newArr, arr2[j])
			j++
		}
	}

	for ; i < len(arr1); i++ {
		newArr = append(newArr, arr1[i])
	}

	for ; j < len(arr2); j++ {
		newArr = append(newArr, arr2[j])
	}

	return newArr
}
