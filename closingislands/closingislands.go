package closingislands

func cerca(table [][]int) int {
	MaxUint := ^uint(0)
	MaxInt := int(MaxUint >> 1)

	best := MaxInt
	for x1 := 0; x1 < len(table); x1++ {
		for x2 := x1; x2 < len(table); x2++ {
			for y1 := 0; y1 < len(table[0]); y1++ {
				for y2 := y1; y2 < len(table[0]); y2++ {
					area := (x2 - x1 + 1) * (y2 - y1 + 1)
					if area < best && allIn(table, x1, x2, y1, y2) {
						best = area
					}
				}
			}
		}
	}
	return best
}

func allIn(table [][]int, x1, x2, y1, y2 int) bool {
	cantity := make(map[int]bool)
	for firstIndex := 0; firstIndex < len(table); firstIndex++ {
		for secondIndex := 0; secondIndex < len(table[0]); secondIndex++ {
			cantity[table[firstIndex][secondIndex]] = true
		}
	}
	_, okZero := cantity[0]
	numberIslands := len(cantity)
	if okZero {
		numberIslands--
	}
	inside := make(map[int]bool)
	for firstIndex := x1; firstIndex <= x2; firstIndex++ {
		for secondIndex := y1; secondIndex <= y2; secondIndex++ {
			inside[table[firstIndex][secondIndex]] = true
		}
	}
	_, okZero = inside[0]
	numberInside := len(inside)
	if okZero {
		numberInside--
	}
	return numberInside == numberIslands
}
