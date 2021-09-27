package permutations

func permuting(original []int, selected []bool, permuted []int, call int) int {
	if call >= len(original) {
		return 1
	}
	times := 0
	for index := 0; index < len(original); index++ {
		if !selected[index] {
			selected[index] = true
			permuted[index] = original[call]
			times += permuting(original, selected, permuted, call+1)
			selected[index] = false
		}
	}
	return times
}
