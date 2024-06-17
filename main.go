package main

import (
	"fmt"
)

type Matrix4x4 [4][4]int

func (m1 Matrix4x4) Add(m2 Matrix4x4) Matrix4x4 {
	var result Matrix4x4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return result
}

func (m1 Matrix4x4) Subtract(m2 Matrix4x4) Matrix4x4 {
	var result Matrix4x4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = m1[i][j] - m2[i][j]
		}
	}
	return result
}

func (m1 Matrix4x4) Multiply(m2 Matrix4x4) Matrix4x4 {
	var result Matrix4x4
	for i := 0; i < 4; i++ {
		for j := 0; i < 4; j++ {
			sum := 0
			for k := 0; k < 4; k++ {
				sum += m1[i][k] * m2[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

func (m Matrix4x4) Print() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	m1 := Matrix4x4{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	m2 := Matrix4x4{
		{16, 15, 14, 13},
		{12, 11, 10, 9},
		{8, 7, 6, 5},
		{4, 3, 2, 1},
	}

	fmt.Println("Matrix 1:")
	m1.Print()

	fmt.Println("Matrix 2:")
	m2.Print()

	fmt.Println("Addition:")
	resultAdd := m1.Add(m2)
	resultAdd.Print()

	fmt.Println("Subtraction:")
	resultSubtract := m1.Subtract(m2)
	resultSubtract.Print()

	fmt.Println("Multiplication:")
	resultMultiply := m1.Multiply(m2)
	resultMultiply.Print()
}
