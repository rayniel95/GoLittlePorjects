package treemincost

import (
	"fmt"
	"math"
)

func noTaken(array []int) int {
	cantity := 0
	for _, value := range array {
		if value == -1 {
			cantity++
		}
	}
	return cantity
}

func selector(array []int, taken []int, index int, taked int, max int, level int, best *int) {
	if taked == max {
		sel(array, taken, level+1, best, max)
		return
	}
	for newIndex := index; newIndex < len(array); newIndex++ {
		if taken[newIndex] == -1 {
			taken[newIndex] = level
			selector(array, taken, newIndex+1, taked+1, max, level, best)
			taken[newIndex] = -1
		}
	}
}

func sel(array []int, taken []int, level int, best *int, taked int) {
	if noTaken(taken) == 0 {
		fmt.Println(taken)
		result := calculus(array, taken)
		if result < *best {
			*best = result
		}
		return
	}
	for take := 1; take <= int(math.Min(float64(noTaken(taken)), math.Pow(2, float64(taked)))); take++ {
		selector(array, taken, 0, 0, take, level, best)
	}
}

func calculus(array []int, taken []int) int {
	result := 0
	for index := 0; index < len(array); index++ {
		result += (array[index] * int(math.Pow(2, float64(taken[index]))))
	}
	return result
}

func constructTree(array []int) int {
	MaxUint := ^uint(0)
	MaxInt := int(MaxUint >> 1)

	best := &MaxInt
	levels := make([]int, len(array))
	for index := 0; index < len(levels); index++ {
		levels[index] = -1
	}
	sel(array, levels, 0, best, 0)
	return *best
}
