package compact

func compact(s string, startIndex int, endIndex int) ([]byte, int, int) {
	patternTimes, firstIndexPat, secondIndexPat := searchPattern(s, startIndex,
		endIndex)

	if patternTimes == -1 {
		return []byte(s)[startIndex:endIndex], startIndex, endIndex
	}
	newPat, firsIndexNew, secondIndexNew := compact(s, firstIndexPat,
		secondIndexPat)

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

	patternLength := endIndexPat - startIndexPat
	patternRepeat := true

	for times := 1; true; times++ {
		for index := startIndex; index < endIndexPat; index++ {
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

}
