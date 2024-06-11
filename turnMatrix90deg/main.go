package main

import (
	"fmt"
)

/*Создание матрицы*/
func createMatrix(x, y, status int) [][]int {
	res := make([][]int, y)

	for i := range res {
		res[i] = make([]int, x, x)
	}

	if status == 1 { //Если статус равен 1, то читаем значения матрицы
		for i, row := range res {
			for k := range row {
				fmt.Scan(&res[i][k])
			}
		}
	}
	return res
}

/*Печать матрицы*/
func printMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {
		for k := 0; k < len(m[i]); k++ {
			if k+1 == len(m[i]) { //Если это последний элемент в строке, то печать без пробела на конце
				fmt.Printf("%v", m[i][k])
			} else {
				fmt.Printf("%v ", m[i][k])
			}
		}
		fmt.Print("\n")
	}
}

/*Поворот матрицы*/
func turn90deg(matrix [][]int) [][]int {
	resX := len(matrix) - 1
	resY := len(matrix[0]) - 1
	res := createMatrix(resX+1, resY+1, -1)
	for i := range matrix {
		resCollumn := resX - i
		for k := resY; k >= 0; k-- {
			res[resY-k][resCollumn] = matrix[i][resY-k]
		}
	}
	return res
}

func main() {
	var (
		x, y int
	)
	fmt.Scanf("%v %v", &x, &y)

	matrix := createMatrix(y, x, 1)

	matrix2 := turn90deg(matrix)

	printMatrix(matrix2)
}
