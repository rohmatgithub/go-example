package tests

import (
	"fmt"
	"math"
	"testing"
)

func TestRandom(t *testing.T) {
	// InsertRandomWord()
	matrix := [][]int32{
		{112, 42, 83, 119}, // 00, 01, 02, 03
		{56, 125, 56, 49},  // 10, 11, 12, 13
		{15, 78, 101, 43},  // 20, 21, 22, 23
		{62, 98, 114, 108}, // 30, 31, 32, 33
	}

	// q := 1
	// n := 2

	matrixNew := make([][]int32, len(matrix))
	for i := range matrix {
		matrixNew[i] = make([]int32, len(matrix[i]))
		copy(matrixNew[i], matrix[i])
	}

	// // transpose the matrix
	// for i, v := range matrix {
	// 	for j, v2 := range v {
	// 		matrixNew[j][i] = v2
	// 	}
	// }

	// rotate the matrix
	for i, v := range matrix {
		for j, v2 := range v {
			matrixNew[len(matrix)-1-i][j] = v2
		}
	}

	for i, v := range matrixNew {
		fmt.Printf("%d : %v\n", i, v)
	}
}

func TestString(t *testing.T) {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())
}

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func startTest(input []string, target string) (result [][]string) {

	return
}
