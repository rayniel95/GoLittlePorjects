package bigsubpalindrome

func isPal(array []byte) bool {
	pal := true
	for index := 0; index < len(array)/2; index++ {
		if array[index] != array[len(array)-1-index] {
			pal = false
			break
		}
	}
	return pal
}

func bigSubPal(array []byte) []byte {
	biggest := make([]byte, 0)
	for firstIndex := 0; firstIndex < len(array); firstIndex++ {
		for secondIndex := firstIndex + 1; secondIndex < len(array); secondIndex++ {
			if isPal(array[firstIndex:secondIndex+1]) && 1+secondIndex-firstIndex > len(biggest) {
				biggest = array[firstIndex : secondIndex+1]
			}
		}
	}
	return biggest
}
