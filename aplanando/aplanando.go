package aplanando

import "fmt"

type mov struct {
	stack   int
	cantity int
}

func aplana(stacks []int) int {
	best := []mov{}
	aplanando(stacks, []mov{}, arrayMin(stacks), &best)
	return len(best)
}

func aplanando(stacks []int, movs []mov, min int, best *[]mov) {
	if isAplaned(stacks) {
		if len(movs) < len(*best) {
			*best = movs
		}
		return
	}
	fmt.Println(movs)
	for stack := 0; stack < len(stacks); stack++ {
		if stack == 0 || stack == len(stacks)-1 {
			ady := 1
			if stack == len(stacks)-1 {
				ady = len(stacks) - 2
			}
			for boxs := 1; stacks[stack]-boxs > min && stacks[stack]-boxs > stacks[ady]; boxs++ {

				newStacks := make([]int, len(stacks))
				copy(newStacks, stacks)
				newStacks[stack] -= boxs
				newStacks[ady] += boxs
				aplanando(newStacks, append(movs, mov{stack, boxs}), min, best)
			}
		} else {
			for boxs := 1; stacks[stack]-2*boxs > min && (stacks[stack]-2*boxs > stacks[stack-1] ||
				stacks[stack]-2*boxs > stacks[stack+1]); boxs++ {

				newStacks := make([]int, len(stacks))
				copy(newStacks, stacks)
				newStacks[stack] -= 2 * boxs
				newStacks[stack-1] += boxs
				newStacks[stack+1] += boxs
				aplanando(newStacks, append(movs, mov{stack, boxs}), min, best)
			}
		}
	}
}

func isAplaned(stacks []int) bool {
	for index := 1; index < len(stacks); index++ {
		if stacks[index] != stacks[0] {
			return false
		}
	}
	return true
}

func arrayMin(array []int) int {
	min := array[0]
	for elem := range array {
		if elem < min {
			min = elem
		}
	}
	return min
}
