package permutations

import "fmt"

func permuting(original []int, selected []bool, permuted []int, call int) int {
	if call >= len(original) {
		fmt.Println(permuted)
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

func permuting2(original []int, selected []bool, permuted []int, call int) int {
	if call >= len(permuted) {
		fmt.Println(permuted)
		return 1
	}
	times := 0
	for index := 0; index < len(original); index++ {
		if !selected[index] {
			selected[index] = true
			permuted[call] = original[index]
			times += permuting2(original, selected, permuted, call+1)
			selected[index] = false
		}
	}
	return times
}

func permutingWithRep(original []int, permuted []int, call int) int {
	if call >= len(permuted) {
		fmt.Println(permuted)
		return 1
	}
	times := 0
	for index := 0; index < len(original); index++ {
		permuted[call] = original[index]
		times += permutingWithRep(original, permuted, call+1)
	}
	return times
}
