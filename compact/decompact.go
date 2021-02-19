package compact

import (
	"bytes"
	"fmt"
	"strconv"
)

func decompact(s []byte) []byte {
	number := make([]byte, 0)
	leftPart := make([]byte, 0)
	rightPart := make([]byte, 0)

	var leftIndex int
	var rightIndex int

	for leftIndex = 0; leftIndex <= len(s); leftIndex++ {
		if leftIndex == len(s) {
			return leftPart
		}
		if s[leftIndex] == '(' {
			break
		}
		leftPart = append(leftPart, s[leftIndex])
	}
	for rightIndex = len(s) - 1; rightIndex > leftIndex && s[rightIndex] != ')'; rightIndex-- {
		if _, err := strconv.Atoi(string(s[rightIndex])); err == nil {
			number = append(number, s[rightIndex])
		} else {
			rightPart = append(rightPart, s[rightIndex])
		}
	}
	fmt.Println()
	fmt.Println(leftIndex)
	fmt.Println(rightIndex)
	fmt.Println(string(s))
	repeatTime, _ := strconv.Atoi(string(reverseByte(number)))
	temp := append(leftPart, bytes.Repeat(decompact(s[leftIndex+1:rightIndex]),
		repeatTime)...)

	return append(temp, reverseByte(rightPart)...)
}

func reverseByte(s []byte) []byte {
	for index := 0; index < len(s)/2; index++ {
		temp := s[index]
		s[index] = s[len(s)-1-index]
		s[len(s)-1-index] = temp
	}
	return s
}
