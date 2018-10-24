package arrayslices

func Sum(array []int) (sum int) {
	for _, number := range array {
		sum += number
	}
	return
}

func SumAll(slices [][]int) []int {
	sumSlice := make([]int, len(slices))
	for i, slice := range slices {
		sum := 0
		for _, number := range slice {
			sum += number
		}
		sumSlice[i] = sum
	}
	return sumSlice
}
