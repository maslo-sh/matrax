package utils

import (
	"fmt"
	"strings"
)

type Matrix struct {
	Elements   [][]int
	Columns    int
	Rows       int
	Rank       int
	IsDiagonal bool
}

func (self *Matrix) GetTransposed() *Matrix {
	var tempMatr = make([][]int, self.Columns)
	for i := 0; i < self.Columns; i++ {
		tempMatr[i] = make([]int, self.Rows)
	}

	for i := 0; i < self.Columns; i++ {
		for j := 0; j < self.Rows; j++ {
			tempMatr[i][j] = self.Elements[j][i]
		}
	}

	return &Matrix{
		Elements:   tempMatr,
		Columns:    len(tempMatr[0]),
		Rows:       len(tempMatr),
		Rank:       0,
		IsDiagonal: false,
	}

}

func (self *Matrix) SetElement(row int, col int, val int) {
	self.Elements[row][col] = val
}

func (self *Matrix) String() string {
	var sb strings.Builder
	for _, row := range self.Elements {
		sb.WriteString(fmt.Sprintf("%v\n", row[0:]))
	}

	return sb.String()
}

func (self *Matrix) updateProperties() {
	self.Rows = len(self.Elements)
	self.Columns = len(self.Elements[0])
}

func NewMatrix(height int, width int) *Matrix {
	var matr = make([][]int, height)
	for i := 0; i < width; i++ {
		matr[i] = make([]int, width)
	}

	return &Matrix{
		Elements:   matr,
		Columns:    width,
		Rows:       height,
		Rank:       0,
		IsDiagonal: true,
	}
}

func (self *Matrix) InsertColumn(column []int, index int) {
	for i, val := range column {
		self.SetElement(i, index, val)
	}
}

func (self *Matrix) InsertRow(row []int, index int) {
	self.Elements[index] = row
}
