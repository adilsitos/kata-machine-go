package algorithms

// the greater element stays at the end
func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				aux := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = aux
			}
		}
	}

	return arr
}
