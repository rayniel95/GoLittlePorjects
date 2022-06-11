package securepath

import "fmt"

func warning(table [][]int, cell [2]int) bool {
	return (cell[0]+1 < 8 && cell[1]-1 >= 0 && table[cell[0]+1][cell[1]-1] == -1) ||
		(cell[0]-1 >= 0 && cell[1]+1 < 8 && table[cell[0]-1][cell[1]+1] == -1) ||
		(cell[0]-1 >= 0 && cell[1]-1 >= 0 && table[cell[0]-1][cell[1]-1] == -1) ||
		(cell[0]+1 < 8 && cell[1]+1 < 8 && table[cell[0]+1][cell[1]+1] == -1)
}

func bestPath(table [][]int, cell [2]int, number int) {
	movRows := [8]int{0, -1, -1, -1, 0, 1, 1, 1}
	movCols := [8]int{-1, -1, 0, 1, 1, 1, 0, -1}
	vomita := false
	for index := 0; index < 8; index++ {
		newCell := [2]int{cell[0] + movRows[index], cell[1] + movCols[index]}
		if isValid(newCell) && !warning(table, newCell) && (table[newCell[0]][newCell[1]] > number+1 ||
			table[newCell[0]][newCell[1]] == 0 || table[newCell[0]][newCell[1]] == -1) {

			if table[newCell[0]][newCell[1]] == -1 {
				vomita = true
			}
			table[newCell[0]][newCell[1]] = number + 1
			bestPath(table, newCell, number+1)
			if vomita {
				table[newCell[0]][newCell[1]] = -1
				vomita = false
			}
		}
	}
}

func isValid(cell [2]int) bool {
	return cell[0] >= 0 && cell[0] < 8 && cell[1] < 8 && cell[1] >= 0
}

func path(table [][]int) int {
	table[7][0] = 1
	bestPath(table, [2]int{7, 0}, 1)
	printMatrix(table)
	return table[0][7]
}

func printMatrix(table [][]int) {
	for index := 0; index < len(table); index++ {
		fmt.Println(table[index])
	}
	fmt.Println()
}
