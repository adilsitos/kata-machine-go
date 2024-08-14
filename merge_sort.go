package katamachinegolang

func MergeSort(sl []int) []int {
	if len(sl) < 2 {
		return sl
	}

	// divide the input in two halves
	sl1 := MergeSort(sl[:len(sl)/2])
	sl2 := MergeSort(sl[len(sl)/2:])

	return merge(sl1, sl2)

}

func merge(sl1, sl2 []int) []int {
	var newArr []int
	i := 0
	j := 0

	for i < len(sl1) && j < len(sl2) {
		if sl1[i] <= sl2[j] {
			newArr = append(newArr, sl1[i])
			i++
		} else {
			newArr = append(newArr, sl2[j])
			j++
		}
	}

	for ; i < len(sl1); i++ {
		newArr = append(newArr, sl1[i])
	}

	for ; j < len(sl2); j++ {
		newArr = append(newArr, sl2[j])
	}

	return newArr
}
