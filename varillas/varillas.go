package varillas

func rebuild(array []int) int {
	for cantity := len(array) - 1; cantity > 0; cantity-- {
		result := subsetSum(array, cantity, make([]int, len(array)), 0)
		if len(result) > 0 {
			return cantity
		}
	}
	return -1
}

func subsetSum(array []int, k int, assign []int, index int) []int {
	if index >= len(array) {
		if subsetEqualSum(array, assign, k) {
			return assign
		}
		return []int{}
	}
	newAssign := make([]int, len(array))
	copy(newAssign, assign)
	for subset := 0; subset < k; subset++ {
		newAssign[index] = subset
		result := subsetSum(array, k, newAssign, index+1)
		if len(result) > 0 {
			return result
		}
	}
	return []int{}
}

func subsetEqualSum(array []int, assign []int, k int) bool {
	sum := make([]int, k)

	for index, val := range assign {
		sum[val] += array[index]
	}
	for _, val := range sum {
		if val != sum[0] {
			return false
		}
	}
	return true
}
