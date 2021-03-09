package aplanando

import "fmt"

type mov struct {
	stack   int
	cantity int
}

func aplana(stacks []int) int {
	best := []mov{}
	aplanando(stacks, []mov{}, arraySum(stacks)/len(stacks), &best, 0)
	return len(best)
}

func aplanando(stacks []int, movs []mov, min int, best *[]mov, indiceAnt int) {
	if isAplaned(stacks) {
		if len(movs) < len(*best) {
			*best = movs
		}
		return
	} else if len(*best) > 0 && len(movs) > len(*best) { // poda
		return
	}
	fmt.Println(stacks)
	fmt.Println(movs)

	for stack := 0; stack < len(stacks); stack++ {
		if stack == 0 || stack == len(stacks)-1 {
			ady := 1
			if stack == len(stacks)-1 {
				ady = len(stacks) - 2
			} // creo es posible realizar otra poda en este for, se puede mover toda las cajas hasta que stacks[stack] == min
			for boxs := 1; stacks[stack]-boxs >= min; boxs++ {
				fmt.Println(boxs)
				fmt.Println()
				newStacks := make([]int, len(stacks))
				copy(newStacks, stacks)
				newStacks[stack] -= boxs
				newStacks[ady] += boxs
				aplanando(newStacks, append(movs, mov{stack, boxs}), min, best, stack+1)
			}
		} else {
			for boxs := 1; stacks[stack]-2*boxs >= min; boxs++ {
				fmt.Println(boxs)
				fmt.Println()
				newStacks := make([]int, len(stacks))
				copy(newStacks, stacks)
				newStacks[stack] -= 2 * boxs
				newStacks[stack-1] += boxs
				newStacks[stack+1] += boxs
				aplanando(newStacks, append(movs, mov{stack, boxs}), min, best, stack+1)
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
	for _, elem := range array {
		if elem < min {
			min = elem
		}
	}
	return min
}

func arraySum(array []int) int {
	sum := 0
	for _, elem := range array {
		sum += elem
	}
	return sum
}
