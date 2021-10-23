package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func isPalindrome(number int) bool {
	// your code goes here
	stringNumber := strconv.Itoa(number)
	for index := 0; index < len(stringNumber)/2; index++ {
		if stringNumber[index] != stringNumber[len(stringNumber)-1-index] {
			return false
		}
	}
	return true
}
func main() {
	var number int = rand.Intn(10000)
	if isPalindrome(number) {
		fmt.Printf("%d is a palindrome number", number)
	} else {
		fmt.Printf("%d is NOT a palindrome number", number)
	}
}
