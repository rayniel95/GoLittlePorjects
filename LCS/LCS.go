package lcs

import (
	"bytes"
	"fmt"
)

func toLCS(small string, large string) string {
	best := ""
	lcs([]byte(small), []byte(large), &best, []byte{}, 0, 0, 0)
	fmt.Println(best)
	return best
}

// bestString debe ser una referencia mientras actualString si necesito que
// sea una copia en cada llamado
func lcs(
	small []byte, large []byte, bestString *string, actualString []byte,
	actualIndex int, actualLargeIndex int, rec int,
) {
	if len(actualString) > len(*bestString) {
		*bestString = string(actualString)
	}

	for index := actualIndex; index < len(small); index++ {

		if bytes.Contains(large[actualLargeIndex:], []byte{small[index]}) {
			lcs(
				small, large, bestString,
				append(actualString, small[index]), index+1,
				len(large[:actualLargeIndex])+bytes.IndexByte(
					large[actualLargeIndex:], small[index]), rec+1,
			)
		}
	}
}
