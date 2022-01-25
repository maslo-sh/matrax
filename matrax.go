package main

import (
	"fmt"
	"matrax/utils"
)

func main() {

	matrix1 := NewMatrix(2, 2)
	matrix1.Elements[0][0] = 1
	matrix1.Elements[0][1] = 2
	matrix1.Elements[1][0] = 3
	matrix1.Elements[1][1] = 1444

	matrix2 := NewMatrix(2, 2)
	matrix2.Elements[0][0] = 10
	matrix2.Elements[0][1] = 22
	matrix2.Elements[1][0] = 33
	matrix2.Elements[1][1] = 1
	fmt.Println(matrix2.String())

	newMatr := utils.Multiply(*matrix1, *matrix2)
	fmt.Println(newMatr.String())

	m1 := *NewMatrix(2, 5)
	m2 := *NewMatrix(5, 2)

	m1.InsertRow([]int{1, 2, 3, 4, 5}, 0)
	m1.InsertRow([]int{1, 2, 22, 14, 15}, 1)
	m2.InsertColumn([]int{10, 22, 33, 1, 1}, 0)
	m2.InsertColumn([]int{100, 122, 233, 11, 1}, 1)

	newMatr1 := utils.Multiply(m1, m2)
	fmt.Println(fmt.Sprintf("%v * %v = %v", m1.String(), m2.String(), newMatr1.String()))
	fmt.Println(newMatr1.String())
}

func NewMatrix(height int, width int) *utils.Matrix {
	var matr = make([][]int, height)
	for i, _ := range matr {
		matr[i] = make([]int, width)
	}

	return &utils.Matrix{
		Elements:   matr,
		Columns:    width,
		Rows:       height,
		Rank:       0,
		IsDiagonal: true,
	}
}
