package steers

import "fmt"

func steers(sticks int) int {
	if sticks == 0 {
		return 0
	}
	steer := 0
	myArray := make([]int8, sticks, sticks)
	myIndex := 0

	steersA(sticks, false, &steer, myArray, myIndex)
	return steer
}

// TODO - see overload functions
func steersA(sticks int, vertical bool, numberSteers *int, myArray []int8, index int) {
	if sticks == 0 {
		*numberSteers++
		fmt.Println(myArray)
		return
	}
	temp := myArray[index]
	if vertical {
		myArray[index] = 0
		steersA(sticks-1, false, numberSteers, myArray, index+1)
		myArray[index] = temp
		return
	}
	myArray[index] = 0
	steersA(sticks-1, false, numberSteers, myArray, index+1)
	myArray[index] = 1
	steersA(sticks-1, true, numberSteers, myArray, index+1)
	myArray[index] = temp
}
