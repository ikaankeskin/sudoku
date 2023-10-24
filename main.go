package main

//Test #1 valid entry
//go run . "3...859.4" "...629..3" "...7...2." "5.6....9." "7.1.9.3.8" ".2....5.1" ".8...7..." "2..918..." "9.746...5" | cat -e
//Result
//3 6 2 1 8 5 9 7 4$
//4 7 8 6 2 9 1 5 3$
//1 9 5 7 3 4 8 2 6$
//5 3 6 8 4 1 7 9 2$
//7 4 1 5 9 2 3 6 8$
//8 2 9 3 7 6 5 4 1$
//6 8 3 2 5 7 4 1 9$
//2 5 4 9 1 8 6 3 7$
//9 1 7 4 6 3 2 8 5$

//Test #2 valid entry
//go run . "2.5..9..4" "......3.7" "7..856.1." "45.7....." "..9...1.." ".....2.85" ".2.418..6" "6.8......" "1..2..7.8" | cat -e
//2 1 5 3 7 9 8 6 4$
//9 8 6 1 2 4 3 5 7$
//7 3 4 8 5 6 2 1 9$
//4 5 2 7 8 1 6 9 3$
//8 6 9 5 4 3 1 7 2$
//3 7 1 6 9 2 4 8 5$
//5 2 7 4 1 8 9 3 6$
//6 4 8 9 3 7 5 2 1$
//1 9 3 2 6 5 7 4 8$

//Test #3 valid entry
//go run . "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79" | cat -e
//5 3 4 6 7 8 9 1 2$
//6 7 2 1 9 5 3 4 8$
//1 9 8 3 4 2 5 6 7$
//8 5 9 7 6 1 4 2 3$
//4 2 6 8 5 3 7 9 1$
//7 1 3 9 2 4 8 5 6$
//9 6 1 5 3 7 2 8 4$
//2 8 7 4 1 9 6 3 5$
//3 4 5 2 8 6 1 7 9$

//Test #4 Valid entry
//go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
//3 9 6 2 4 5 7 8 1$
//1 7 8 3 6 9 5 2 4$
//5 2 4 8 1 7 3 9 6$
//2 8 7 9 5 1 6 4 3$
//9 3 1 4 8 6 2 7 5$
//4 6 5 7 2 3 9 1 8$
//7 1 2 6 3 8 4 5 9$
//6 5 9 1 7 4 8 3 2$
//8 4 3 5 9 2 1 6 7$

//Test #5 Invalid entry, wrong number of arguments
//go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." ".1.63..59" ".59.7.83." "..359...7" | cat -e
//Error

//Test #6 Invalid entry, wrong number of cells in one of arguments
//go run . ".96.4...1" "1...6...4" "5.481.39." "..95..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
//Error

//Test #7 Invalid entry, wrong number of cells in one of arguments
//go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83.." "..359...7" | cat -e
//Error

//Test #8 Invalid entry, parameters contain invalid charachters
//go run . ".96.4.a.1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
//Error

//Test #9  Invalid entry, repeating characters in 2nd row
//go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
//Error$

//Test #9  Invalid entry, repeating characters in 5th column
//go run . ".96.4...1" "1...63..4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
//Error$

//Test #10 Invalid entry, repeating characters in last square
//go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359.8.7" | cat -e
//Error$

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	//Get input and save into slice.
	//[1:] means that we take all parameters except first one, which is a program name.
	slRows := os.Args[1:]

	//Validation
	//Validation #1: Check number of arguments.
	if len(slRows) != 9 {
		fmt.Println("Error")
		return
	}

	//Validation #2: Check that every argument has size of 9
	for i := 0; i < 9; i++ {
		if len(slRows[i]) != 9 {
			fmt.Println("Error")
			return
		}
	}

	//Validation #3: Check every argument contains either digit or dot
	for i := 0; i < 9; i++ {
		for j := 0; j < len(slRows[i]); j++ {
			ch := slRows[i][j]
			if !(ch >= '1' && ch <= '9') && ch != '.' {
				fmt.Println("Error")
				return
			}
		}
	}

	//Validation #4: Check for repating characters in input
	for i := 0; i < len(slRows); i++ {
		row := slRows[i]
		if HasRepeatingCharactersSudokuHelperFunction(row, '.') {
			fmt.Println("Error")
			return
		}
	}

	//We need three slices of strings: slice of rows, slice of columns and slice of 3x3 squares
	//We have slice of rows as input. Let's create slices of columns and squares
	//Create slice of columns
	slColumns := CalculateSliceOfColumnsSudokuHelperFunction(slRows)

	//Validation #5: Check for repating characters in columns
	for i := 0; i < len(slColumns); i++ {
		col := slColumns[i]
		if HasRepeatingCharactersSudokuHelperFunction(col, '.') {
			fmt.Println("Error")
			return
		}
	}

	//Create slice of local 3x3 sqares as slice of strings
	//This is how  local square is represented:
	//
	// 0 1 2
	// 3 4 5  => 012345678
	// 6 7 8
	//
	slSquares := CalculateSliceOfSquaresSudokuHelperFunction(slRows)

	//Validation #6: Check for repating characters in squares
	for i := 0; i < len(slSquares); i++ {
		sq := slSquares[i]
		if HasRepeatingCharactersSudokuHelperFunction(sq, '.') {
			fmt.Println("Error")
			return
		}
	}

	//Solution
	/*
		 *
		a) take a square 3x3 (local square)
		b) find all empty cells within local square
		c) run through digits 1 to 9
		d) check that digit ia not already present in local squre, if present - move on to next digit
		e) put current digit in each empty cell within local square and check respective row and column
		   for presence of same digit
		f) if we can put same digit in more than one cell in local square, then move to next digit,
		   otherwise update slice of rows, columns and squares
		g) repeat this until sudoku is solved
		h) sudoku is solved when total number of non-empty cells is 81
		*
	*/

	// Loop: is a lable.
	// It is used in combination with break at line 207 so program knows which loop to break.
	// In our case outer loop should be broken
Loop:
	for {
		//Use this variable to detect exit condition
		numberOfNonEmptyCells := 0

		//Detect if any solution found during this run
		//Each run through the loop should result in finding one valid digit
		//If no valid digit if fund we cold get stuck inside infinate loop
		//Assume that we woudn't find valid digit this run
		isSolutionFound := false

		//Run through local 3x3 squares
		for sqIndex := 0; sqIndex < 9; sqIndex++ {
			//Get values of a square
			sq := slSquares[sqIndex]

			//Find empty cells of square
			emptyCells := []int{}
			for cellIndex := 0; cellIndex < 9; cellIndex++ {
				//find empty cells in quare
				if !(sq[cellIndex] >= '1' && sq[cellIndex] <= '9') {
					emptyCells = append(emptyCells, cellIndex)
				}
			}
			//Check if any empty cells were found
			//Also check, maybe sudoku is already solved
			if len(emptyCells) == 0 {
				numberOfNonEmptyCells = numberOfNonEmptyCells + 9
				if numberOfNonEmptyCells >= 81 {
					break Loop
				}
			}

			//Run through digits
			for digit := '1'; digit <= '9'; digit++ {
				//Skip digits that already in the square.
				//Function GetIndexOfStringSudokuHelperFunction() returns index of a character or -1 if character not found.
				if GetIndexOfStringSudokuHelperFunction(sq, string(digit)) == -1 {

					fits := []int{} //Stores square indesies in which digit would fit
					//Try to put digit into each empty cell
					for i := 0; i < len(emptyCells); i++ {
						//actual index of empty cell in square
						emptyCellIndexLocal := emptyCells[i]
						//Convert local cellIndex into global cell index
						globalRowIndex := (sqIndex/3)*3 + emptyCellIndexLocal/3
						globalColumnIndex := (sqIndex%3)*3 + emptyCellIndexLocal%3
						//Check row and column
						//Current digit is not supposed to be present in neither column nor row
						if GetIndexOfStringSudokuHelperFunction(slRows[globalRowIndex], string(digit)) == -1 &&
							GetIndexOfStringSudokuHelperFunction(slColumns[globalColumnIndex], string(digit)) == -1 {
							//If digit can go in current cell, add index of that cell to 'fits'
							fits = append(fits, emptyCellIndexLocal)
						}
					}
					//Valid solusion would be when only one empty cell within local square can contain current digit
					//Otherewise we don't know in waht cell to put current digit
					if len(fits) == 1 {
						//Local valid cell in square
						index := fits[0]
						//Convert local valid index in square into global row and col values
						globRowIndex := (sqIndex/3)*3 + index/3
						globColIndex := (sqIndex%3)*3 + index%3

						//Update slRow
						newRow := ReplaceCharInStringSudokuHelperFunction(slRows[globRowIndex], digit, globColIndex)
						slRows[globRowIndex] = newRow
						//Update columns
						slColumns = CalculateSliceOfColumnsSudokuHelperFunction(slRows)
						//Update squares
						slSquares = CalculateSliceOfSquaresSudokuHelperFunction(slRows)

						//Valid digit was found this run so we set it to true
						isSolutionFound = true
					}
				}
			}
		}

		//This block make extra check for one missing digit in row and column
		//For example if we have a row like this: 156.87392 we can insert 4 as missing number
		//Check all rows
		for rIndex := 0; rIndex < 9; rIndex++ {
			numberOfDigit := GetNumberOfDigitFromStringSudokuHelperFunction(slRows[rIndex])
			if numberOfDigit == 8 {

				missingDigit, indx := FindMissingNumberAndIndexSudokuHelperFunction(slRows[rIndex])

				if missingDigit > 0 && indx > 0 {
					//Insert missing digit into row
					currentRow := slRows[rIndex]
					currentRowAsRunes := []rune(currentRow)
					currentRowAsRunes[indx] = missingDigit
					slRows[rIndex] = string(currentRow)
					//Update columns and squares
					slColumns = CalculateSliceOfColumnsSudokuHelperFunction(slRows)
					slSquares = CalculateSliceOfSquaresSudokuHelperFunction(slRows)

					//Valid digit was found this run so we set it to true
					isSolutionFound = true
				}
			}
		}

		//Check all columns
		for cIndex := 0; cIndex < 9; cIndex++ {
			numberOfDigit := GetNumberOfDigitFromStringSudokuHelperFunction(slColumns[cIndex])
			if numberOfDigit == 8 {
				missingDigit, indx := FindMissingNumberAndIndexSudokuHelperFunction(slColumns[cIndex])

				if missingDigit > 0 && indx > 0 {
					//Insert missing digit into column
					currentCol := slColumns[cIndex]
					currentColAsRunes := []rune(currentCol)
					currentColAsRunes[indx] = missingDigit
					slColumns[cIndex] = string(currentCol)

					//Update rows
					rowToUpdate := slRows[indx]
					rowToUpdateAsRunes := []rune(rowToUpdate)
					rowToUpdateAsRunes[cIndex] = missingDigit
					slRows[indx] = string(rowToUpdateAsRunes)

					//Update squares
					slSquares = CalculateSliceOfSquaresSudokuHelperFunction(slRows)

					//Valid digit was found this run so we set it to true
					isSolutionFound = true
				}
			}
		}

		//When no valid digit was found during this run, that means we would have to start guessing numbers.
		//This is a case when sudoku could have multiple solutions.
		//These kind of sudokus are considered to be invalid so we return error end exit ptogram.
		if isSolutionFound == false {
			fmt.Println("Error")
			return
		}
	}

	//Printing result
	for _, v := range slRows {
		slRunes := []rune(v)
		for i := 0; i < len(slRunes); i++ {
			z01.PrintRune(slRunes[i])
			if i == len(slRunes)-1 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(' ')
			}
		}
	}
}

//This function returns index of string inside another string, otherwise -1 is returned
func GetIndexOfStringSudokuHelperFunction(s string, toFind string) int {
	sliceS := []rune(s)
	sliceToFind := []rune(toFind)

	for i := 0; i < len(s); i++ {
		isFound := true
		for iToFind := 0; iToFind < len(toFind); iToFind++ {
			newIndex := i + iToFind
			if newIndex < len(sliceS) {
				if sliceS[newIndex] != sliceToFind[iToFind] {
					isFound = false
				}
			}
		}
		if isFound {
			return i
		}
	}
	return -1
}

//Replace a charachter at given position inside string
func ReplaceCharInStringSudokuHelperFunction(s string, ch rune, position int) string {
	slRunes := []rune(s)
	slRunes[position] = ch
	return string(slRunes)
}

//Calculate a slice of columns based on slice of rows
func CalculateSliceOfColumnsSudokuHelperFunction(slRows []string) []string {
	slColumns := []string{}
	for col := 0; col < 9; col++ {
		colString := ""
		for row := 0; row < 9; row++ {
			slRowRunes := []rune(slRows[row])
			colString = colString + string(slRowRunes[col])
		}
		slColumns = append(slColumns, colString)
	}
	return slColumns
}

//Calculate a slice of local 3x3 squares based on slice of rows
func CalculateSliceOfSquaresSudokuHelperFunction(slRows []string) []string {
	slSquares := []string{}
	for row := 0; row < 9; row = row + 3 {
		for col := 0; col < 9; col = col + 3 {

			square :=
				string(slRows[row][col]) + string(slRows[row][col+1]) + string(slRows[row][col+2]) +
					string(slRows[row+1][col]) + string(slRows[row+1][col+1]) + string(slRows[row+1][col+2]) +
					string(slRows[row+2][col]) + string(slRows[row+2][col+1]) + string(slRows[row+2][col+2])

			slSquares = append(slSquares, string(square))
		}
	}
	return slSquares
}

//Return true if there are repating charachters.
//Function will not check characters passed in 'ignore' parameter.
//In our case we ignore dots, otherwide all rows will be retun true
func HasRepeatingCharactersSudokuHelperFunction(s string, ignore rune) bool {
	if len(s) < 2 {
		return false
	}

	sl := SortRunesAscSudokuHelperFunction([]rune(s))
	for i := 1; i < len(sl); i++ {
		if sl[i-1] == sl[i] {
			if sl[i] != ignore {
				return true
			}
		}
	}
	return false
}

//Sort runes in ascending order.
func SortRunesAscSudokuHelperFunction(sliceInput []rune) []rune {
	if len(sliceInput) <= 1 {
		return sliceInput
	}
	isSorted := false
	for !isSorted {
		isSorted = true
		for index := 0; index < len(sliceInput)-1; index++ {
			item1 := sliceInput[index]
			item2 := sliceInput[index+1]
			if item1 > item2 {
				isSorted = false
				temp := sliceInput[index]
				temp1 := sliceInput[index+1]
				sliceInput[index] = temp1
				sliceInput[index+1] = temp
			}
		}
	}
	return sliceInput
}

//Calculate number of characters from '1' to '9' in string
func GetNumberOfDigitFromStringSudokuHelperFunction(s string) int {
	counter := 0
	for digitIndex := 0; digitIndex < len(s); digitIndex++ {
		ch := s[digitIndex]
		chAsRune := rune(ch)
		if chAsRune >= '1' && chAsRune <= '9' {
			counter++
		}
	}
	return counter
}

//Finds missing digit and returns that digit as rune and index
func FindMissingNumberAndIndexSudokuHelperFunction(s string) (int32, int) {
	var missingDidit int32
	//find what digit is missing
	for ch := '1'; ch <= '9'; ch++ {
		if GetIndexOfStringSudokuHelperFunction(s, string(ch)) == -1 {
			missingDidit = ch
			break
		}
	}
	if missingDidit == 0 {
		return 0, -1
	}

	//Find index of missing digit
	index := GetIndexOfStringSudokuHelperFunction(s, string('.'))
	if index == -1 {
		return 0, -1
	}

	return missingDidit, index
}
