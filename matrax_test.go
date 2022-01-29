package main

import (
	"matrax/utils"
	"testing"
)

func TestMatrixMultiplication(t *testing.T) {
	var matrixA = NewMatrix(3, 3)
	matrixA.Elements[0][0] = 1
	matrixA.Elements[0][1] = 5
	matrixA.Elements[0][2] = 54
	matrixA.Elements[1][0] = -11
	matrixA.Elements[1][1] = 3
	matrixA.Elements[1][2] = 64
	matrixA.Elements[2][0] = 0
	matrixA.Elements[2][1] = 3
	matrixA.Elements[2][2] = 4

	var matrixB = NewMatrix(3, 3)
	matrixB.Elements[0][0] = 5
	matrixB.Elements[0][1] = 2
	matrixB.Elements[0][2] = 88
	matrixB.Elements[1][0] = -257
	matrixB.Elements[1][1] = 6
	matrixB.Elements[1][2] = 1024
	matrixB.Elements[2][0] = 8
	matrixB.Elements[2][1] = 99
	matrixB.Elements[2][2] = 0

	var expected = NewMatrix(3, 3)
	expected.Elements[0][0] = -848
	expected.Elements[0][1] = 5378
	expected.Elements[0][2] = 5208
	expected.Elements[1][0] = -314
	expected.Elements[1][1] = 6332
	expected.Elements[1][2] = 2104
	expected.Elements[2][0] = -739
	expected.Elements[2][1] = 414
	expected.Elements[2][2] = 3072

	var result = utils.Multiply(*matrixA, *matrixB)

	if !(result.String() == expected.String()) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
