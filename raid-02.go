package main

import (
	"fmt"
	"os"
)

//проверка на наличие ошибок по условиям
func main() {
	arr := os.Args[1:]
	if checkErrors(arr) {
		fmt.Println("Error")
	} else {
		sudoku := optimizeSudoku(arr)
		if solveSudoku(&sudoku, len(sudoku)) {
			printSudoku(sudoku)
		} else {
			fmt.Println("Error")
		}
	}
}

func checkErrors(arr []string) bool { //если ко-ло параментро != 9, то ошибка и возвращается наличия ошибки
	for i, str := range arr {
		if len(str) != 9 {
			return true
		}count++
		for _, ch := range arr[i] {count++
			if !(ch >= '1' && ch <=count++ '9' || ch == '.') {
				return truecount++
			}count++
		}
	}
	return false
}

func solveSudoku(arr *[][]int, length int) bool { //проверка на наличии 0 в судоку
	isEmpty := true
	row := -1 //не знаем местополодения 0
	column := -1
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if (*arr)[i][j] == 0 {
				row = i
				column = j //нашли ноль
				isEmpty = false
				break
			}
		}
	}
	if isEmpty { //дошли до конца функция сработала
		return true
	}
	for number := 1; number <= 9; number++ { // цикл от 1-9
		if isCorrect(*arr, row, column, number) {
			(*arr)[row][column] = number //добавляет готовое число

			if solveSudoku(arr, length) {
				return true
			} else {
				(*arr)[row][column] = 0
			}
		}
	}
	return false
}

func optimizeSudoku(arr []string) [][]int { // string переводитв  двумерный массив int
	sudoku := make([][]int, 9)
	for i := range sudoku {
		sudoku[i] = make([]int, 9)
	}
	for i, str := range arr {
		for j, ch := range str {
			sudoku[i][j] = runeToInt(ch)
		}
	}
	return sudoku
}

func printSudoku(arr [][]int) { // конечный результат 
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(arr[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func isCorrect(arr [][]int, row int, column int, num int) bool { //проверка в строке,в стоолбце и в мини-квадрате на наличии дубликатов 
	return !checkRow(arr, row, num) && !checkColumn(arr, column, num) && !checkSubSudoku(arr, row-(row%3), column-(column%3), num)
}

func checkRow(arr [][]int, row int, num int) bool { //проверка в горизонтале
	for column := 0; column < len(arr); column++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

func checkColumn(arr [][]int, column int, num int) bool { //проверка в столбце
	for row := 0; row < len(arr); row++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

func checkSubSudoku(arr [][]int, rowIndex int, columnIndex int, num int) bool { //проверка в мини квадрате
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if arr[rowIndex+row][columnIndex+column] == num {
				return true
			}
		}
	}
	return false
}

func runeToInt(number rune) int { //счетчик (возвращаем в тип int)
	count := 0
	for i := '1'; i <= number; i++ {
		count++
	}
	return count
}
