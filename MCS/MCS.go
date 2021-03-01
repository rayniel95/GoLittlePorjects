package mcs

import "fmt"

func toMCS(small string, large string) int {
	best := 0
	mcs([]byte(small), []byte(large), &best, 0, 0)
	fmt.Println(best)
	return len(large) - len(small) + len(small) - best
}

// bestString debe ser una referencia mientras actualString si necesito que
// sea una copia en cada llamado
func mcs(
	small []byte, large []byte, bestCount *int, actualCount int, actualIndex int,
) {
	if actualCount > *bestCount {
		*bestCount = actualCount
	}

	for index := actualIndex; index < len(small); index++ {
		if small[index] == large[index] {
			fmt.Println(small[index], large[index], index)
			mcs(small, large, bestCount, actualCount+1, index+1)
		}
	}
}
