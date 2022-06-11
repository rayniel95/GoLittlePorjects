package neveras

func toMerge(nevs [][]int) int {
	MaxUint := ^uint(0)
	MaxInt := int(MaxUint >> 1)

	best := &MaxInt
	mergeNevs(nevs, 0, best)
	return *best
}

func mergeNevs(nevs [][]int, cost int, bestCost *int) {
	if len(nevs) <= 1 {
		if cost < *bestCost {
			*bestCost = cost
		}
		return
	}
	for index := 0; index < len(nevs); index++ {
		for _, dir := range []int{-1, 1} {
			if index+dir < len(nevs) && index+dir >= 0 {
				newNevs, actualCost := merge(nevs, index, dir)
				mergeNevs(newNevs, cost+actualCost, bestCost)
			}
		}
	}
}

// Coge el array que esta en index y lo anade al que esta en
// index + dir
func merge(nevs [][]int, index int, dir int) ([][]int, int) {
	copyNevs := copySlices(nevs)

	mergeNevs := append(nevs[index+dir], nevs[index]...)
	copyNevs[index+dir] = mergeNevs

	part1 := append(copyNevs[:index], copyNevs[index+1:]...)
	return part1, len(nevs[index+dir]) * sum(nevs[index])
}

func sum(array []int) int {
	sum := 0
	for index := 0; index < len(array); index++ {
		sum += array[index]
	}
	return sum
}

func copySlices(src [][]int) [][]int {
	new := make([][]int, len(src))
	for index := 0; index < len(src); index++ {
		new[index] = make([]int, len(src[index]))
	}
	for firstIndex := 0; firstIndex < len(src); firstIndex++ {
		for secondIndex := 0; secondIndex < len(src[firstIndex]); secondIndex++ {
			new[firstIndex][secondIndex] = src[firstIndex][secondIndex]
		}
	}
	return new
}
