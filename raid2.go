package ma1

import "fmt"

var sudoku = [9][9]int{}

func main1(){
  IsValid(sudoku, 0)
  Display(sudoku)
}

// func Display(sudoku[9][9] int){
//   var x, y int

//   for x = 0; x < 9; x++ {
//         fmt.Println("")
//         if(x == 3 || x == 6){
//           fmt.Println(" ")
//         }
//       for y = 0; y < 9; y++ {
//         if(y == 3 || y == 6){
//           fmt.Print("|")
//         }
//          fmt.Print(sudoku[x][y])
//       }

//    }

// }

func AbsentOnLine(k int, sudoku [9][9]int, x int) bool {
  var y int
    for y=0; y < 9; y++ {
        if (sudoku[x][y] == k){
            return false
          }
        }
    return true
}

func AbsentOnRow(k int, sudoku [9][9]int, y int) bool {
  var x int
  for x=0; x < 9; x++{
       if (sudoku[x][y] == k){
           return false
         }
       }
   return true
}

func AbsentOnBloc(k int, sudoku [9][9]int, x int, y int) bool {
  var firstX, firstY int;
  firstX =  x-(x%3)
  firstY =  y-(y%3)
  for x = firstX; x < firstX+3; x++ {
        for y = firstY; y < firstY+3; y++ {
            if (sudoku[x][y] == k){
                return false
              }
        }
      }
    return true

}

func IsValid(sudoku [9][9]int, position int) bool {

  if (position == 9*9){
        return true
      }

      var x, y, k int

    x = position/9
    y = position%9

    if (sudoku[x][y] != 0){
        return IsValid(sudoku, position+1)
      }

    for k=1; k <= 9; k++ {
        if (AbsentOnLine(k,sudoku,x) && AbsentOnRow(k,sudoku,y) && AbsentOnBloc(k,sudoku,x,y)){
            sudoku[x][y] = k

            if (IsValid(sudoku, position+1)){
                return true
              }
        }
    }
    sudoku[x][y] = 0
    return false
}