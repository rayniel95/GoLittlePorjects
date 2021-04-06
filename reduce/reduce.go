package reduce

import "fmt"

func isReduced(array []int) bool {
	for firstIndex := 0; firstIndex < len(array); firstIndex++ {
		secondIndex := firstIndex + 1
		for ; secondIndex < len(array) &&
			array[firstIndex] == array[secondIndex]; secondIndex++ {
		}
		if secondIndex-firstIndex > 2 {
			return false
		}
	}
	return true
}

func reduce(array []int, index int) []int {
	if isReduced(array) {
		fmt.Println(array)
		return array
	}
	best := make([]int, len(array))
	startLooking := 0

	if index < len(array) {
		looking := index - 1
		for ; looking >= 0 && array[index] == array[looking]; looking-- {
		}

		if (looking < index-1 && looking+1-index == 1 &&
			array[index] == array[index+1]) || (looking+1 < index-1) {

			startLooking = looking + 1
		}
	}
	for looking := startLooking; looking < len(array); looking++ {
		cell := looking + 1
		for ; cell < len(array) && array[looking] == array[cell]; cell++ {
		}

		if cell-looking > 2 {
			copyArray := make([]int, len(array))
			copy(copyArray, array)
			newArray := append(copyArray[:looking], copyArray[cell:]...)
			fmt.Println(array)
			fmt.Println(newArray)

			fmt.Println()
			result := reduce(newArray, looking)
			if len(result) < len(best) {
				best = result
			}
		}
	}
	return best
}
