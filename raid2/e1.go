package main
import (
	"fmt"
	"os"
)

func main() {

	board := parseBoard(os.Args)


	if backtrack(&board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}
func InputCheck(args []string) bool {
	if len(args) != 9 {
		fmt.Println("Error")
		return false
	}

	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			fmt.Println("Error")
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for _, k := range args[i] {
			if k == 47 || k == 48 {
				fmt.Println("Error")
				return false
			} else if k < 46 || k > 57 {
				fmt.Println("Error")
				return false
			}
		}
	}
	return true
}
func parseBoard (arg []string) [9][9] int {
	var board = [9][9] int {}
	for j:=1; j<=9;j++{
		for index, letter := range os.Args[j] {
			if letter == 46 {
				letter = 48
		
			}
			board [j-1][index] = int (letter -48)
		}
	}
	return board
} 
func printBoard(arg [9][9]int) {
	for i:=0; i<9; i++ {
		for j:=0; j<9;j++ {
			fmt.Print (arg [i][j])
			if j <8 {
				fmt.Printf(" ")
			}
		}
		fmt. Println()
	}
}
func backtrack(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 1; candidate <= 9; candidate++ {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}