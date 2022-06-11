package compact

import (
	"fmt"
	"strconv"
)

func compact(s string) string {
	value, _, _, _ := compactA(s, 0, len(s)-1)
	return string(value)
}

func compactA(s string, startIndex int, endIndex int) ([]byte, int, int, int) {
	patternTimes, firstIndexPat, secondIndexPat := searchPattern(s, startIndex,
		endIndex)
	fmt.Println(patternTimes, firstIndexPat, secondIndexPat)
	if patternTimes == -1 {
		return []byte(s)[startIndex : endIndex+1], startIndex, endIndex, -1
	}
	newPat, startIndexNewPat, endIndexNewPat, timesNewPat := compactA(
		s, firstIndexPat, secondIndexPat)
	// there have three cases to differentiate
	if firstIndexPat == startIndexNewPat && secondIndexPat == endIndexNewPat {
		part1 := append([]byte("("), newPat...)
		part2 := append(part1, []byte(")")...)
		part3 := append(part2, strconv.Itoa(patternTimes)...)
		return part3, firstIndexPat, secondIndexPat, patternTimes
	}
	newPatLength := endIndexNewPat - startIndexNewPat + 1
	part1 := append([]byte("("), []byte(s)[firstIndexPat:startIndexNewPat]...)
	fmt.Println(string(part1))
	part2 := append(part1, newPat...)
	fmt.Println(string(part2))
	part3 := append(part2, []byte(s)[startIndexNewPat+newPatLength*timesNewPat:secondIndexPat+1]...)
	fmt.Println(string(part3))
	part4 := append(part3, []byte(")")...)
	fmt.Println(string(part4))
	part5 := append(part4, strconv.Itoa(patternTimes)...)
	fmt.Println(string(part5))
	fmt.Println()
	return part5, firstIndexPat, secondIndexPat, patternTimes
}

func searchPattern(s string, startIndex int, endIndex int) (int, int, int) {
	for firstIndex := 0; firstIndex < endIndex; firstIndex++ {
		for secondIndex := endIndex; secondIndex > firstIndex; secondIndex-- {
			times := checkPatternTimes(s, startIndex, endIndex, firstIndex, secondIndex)

			if times > 1 {
				return times, firstIndex, secondIndex
			}
		}
	}
	return -1, -1, -1
}

func checkPatternTimes(s string, startIndex int, endIndex int, startIndexPat int,
	endIndexPat int) int {

	patternLength := endIndexPat - startIndexPat + 1
	patternRepeat := true

	for times := 1; true; times++ {
		for index := startIndexPat; index <= endIndexPat; index++ {
			if times*patternLength+index > endIndex ||
				s[index] != s[times*patternLength+index] {

				patternRepeat = false
				break
			}
		}
		if !patternRepeat {
			return times
		}
	}
	return -1
}
