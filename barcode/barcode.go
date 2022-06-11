package barcode

func sumIndex(array []int, index int) int {
	if index < 0 {
		return 0
	}
	sum := 0
	for ind, value := range array {
		if ind > index {
			break
		}
		sum += value
	}
	return sum
}

func barcode(array []int, sum int, limit int, times *int, index int) {
	if index >= len(array) {
		if sumIndex(array, len(array)-1) == sum {
			*times++
		}
		return
	}
	newArray := make([]int, len(array))
	copy(newArray, array)
	for cant := 1; cant <= limit && sumIndex(array, index-1)+cant <= sum; cant++ {
		newArray[index] = cant
		barcode(newArray, sum, limit, times, index+1)
	}
}

func solveBarcode(sum int, cantity int, limit int) int {
	times := 0
	barcode(make([]int, cantity), sum, limit, &times, 0)
	return times
}
