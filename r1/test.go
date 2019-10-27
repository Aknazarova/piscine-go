package main

import (
	"fmt"
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if Check(args) {
		sudoku := [9][9]rune{}
		revsud := [9][9]rune{}
		sudoku = FillBoard(sudoku, args)

		if SolveSudoku(&sudoku) {
			revsud = FillBoard(revsud, args)
			RevSudoku(&revsud)
		}

		if SolveSudoku(&sudoku) && sudoku == revsud {
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					if j != 8 {
						z01.PrintRune(rune(sudoku[i][j]))
						z01.PrintRune(' ')
					} else {
						z01.PrintRune(rune(sudoku[i][j]))
					}
				}
				z01.PrintRune('\n')
			}
		} else {
			fmt.Println("Error")
		}
	}
}
// в противном случае возвращает Error
func Check(args []string) bool {
	if len(args) != 9 {
		fmt.Println("Error")
		return false
	}

	for r := 0; r < len(args); r++ {
		if len(args[r]) != 9 {
			fmt.Println("Error")
			return false
		}
	}
	for r := 0; r < len(args); r++ {
		for _, char := range args[r] {
			if char == 47 || char == 48 {
				fmt.Println("Error")
				return false
			} else if char < 46 || char > 57 {
				fmt.Println("Error")
				return false
			}
		}
	}
	return true
}

func FillBoard(sudoku [9][9]rune, args []string) [9][9]rune {
	for i := range args {
		for j := range args[i] {
			sudoku[i][j] = rune(args[i][j])
		}
	}
	return sudoku
}

func EmptyCell(sudoku *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == '.' {
				return true
			}
		}
	}
	return false
}

func IsValid(sudoku *[9][9]rune, i int, j int, w rune) bool {
	for l := 0; l < 9; l++ {
		if w == sudoku[l][i] {
			return false
		}
	}

	for k := 0; k < 9; k++ {
		if w == sudoku[j][k] {
			return false
		}
	}

	d := i / 3
	e := j / 3

	for m := 3 * d; m < 3*(d+1); m++ {
		for n := 3 * e; n < 3*(e+1); n++ {
			if w == sudoku[n][m] {
				return false
			}
		}
	}
	return true
}


func SolveSudoku(sudoku *[9][9]rune) bool {
	if !EmptyCell(sudoku) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == '.' {
				for x := '1'; x <= '9'; x++ {
					if IsValid(sudoku, i, j, x) {
						sudoku[i][j] = x
						if SolveSudoku(sudoku) {
							return true
						}
					}
					sudoku[i][j] = '.'
				}
				return false
			}
		}
	}
	return false
}


func RevSudoku(revsud *[9][9]rune) bool {
	if !EmptyCell(revsud) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if revsud[i][j] == '.' {
				for x := '9'; x >= '1'; x-- {
					if IsValid(revsud, i, j, x) {
						revsud[i][j] = x
						if RevSudoku(revsud) {
							return true
						}
					}
					revsud[i][j] = '.'
				}
				return false
			}
		}
	}
	return false
}