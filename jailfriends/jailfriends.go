package jailfriends

func jailfriendsA(assign []int, subsets int, index int, best *int, friendshipMatrix *friendships) {
	if index >= len(assign) {
		satiesfies := friendshipSatisfies(assign, friendshipMatrix)
		if allTaken(assign, subsets) && satiesfies > *best {
			*best = satiesfies
		}
		return
	}
	newAssign := make([]int, len(assign))
	copy(newAssign, assign)
	for subset := 0; subset < subsets; subset++ {
		newAssign[index] = subset
		jailfriendsA(newAssign, subsets, index+1, best, friendshipMatrix)
	}
}

func friendshipSatisfies(assign []int, friendshipMatrix *friendships) int {
	cantity := 0
	for firstIndex := 0; firstIndex < len(assign); firstIndex++ {
		for secondIndex := 0; secondIndex < len(assign); secondIndex++ {
			if assign[firstIndex] == assign[secondIndex] && friendshipMatrix.get(firstIndex, secondIndex) {
				cantity++
			}
		}
	}
	return cantity / 2
}

func allTaken(assign []int, k int) bool {
	taken := make([]bool, k)
	for _, val := range assign {
		taken[val] = true
	}
	for _, val := range taken {
		if !val {
			return false
		}
	}
	return true
}

type friendships struct {
	friends map[int]map[int]bool
}

func (f *friendships) get(i int, j int) bool {
	val, okext := (*f).friends[i]
	if okext {
		_, okin := val[j]
		if okin {
			return true
		}
	}
	return false
}

func (f *friendships) add(i int, j int) {
	_, ok := (*f).friends[i]
	if !ok {
		(*f).friends[i] = make(map[int]bool)
		(*f).friends[i][j] = true
	}
}

func friendshipsFab() *friendships {
	return &friendships{make(map[int]map[int]bool)}
}
