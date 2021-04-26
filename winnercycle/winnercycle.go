package winnercycle

func cashReceive(matrix [][]float32, cycle []int, currency int) float32 {
	cash := float32(1)
	start := currency
	for _, after := range cycle {
		cash *= matrix[start][after]
		start = after
	}
	cash *= matrix[start][currency]
	return cash
}

func findCycle(
	matrix [][]float32, taken []bool, actualCycle []int, best *[]int,
	currency int,
) {
	if cashReceive(matrix, actualCycle, currency) > cashReceive(matrix, *best, currency) {
		*best = make([]int, len(actualCycle))
		copy(*best, actualCycle)
	}
	for index := 0; index < len(matrix); index++ {
		if !taken[index] {
			taken[index] = true
			actualCycle = append(actualCycle, index)
			findCycle(matrix, taken, actualCycle, best, currency)
			taken[index] = false
			temp := make([]int, len(actualCycle)-1)
			copy(temp, actualCycle)
			actualCycle = temp
		}
	}
}

func winnerCycle(matrix [][]float32, currency int) []int {
	taken := make([]bool, len(matrix))
	taken[currency] = true
	best := make([]int, 0)
	findCycle(matrix, taken, make([]int, 0), &best, currency)
	return best
}
