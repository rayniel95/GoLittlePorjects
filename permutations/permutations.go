package permutations

import "fmt"

func permuting(original []int, selected []bool, permuted []int, call int) {
	if call >= len(original) {
		fmt.Println(permuted)
		return
	}
	for index := 0; index < len(original); index++ {
		if !selected[index] {
			selected[index] = true
			permuted[index] = original[index]
			permuting(original, selected, permuted, call+1)
			selected[index] = false
		}
	}
}
