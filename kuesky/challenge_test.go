package main

import (
	"fmt"
	"testing"
)

//dfadfad(
//dfasopadfad
//dfadfad
func Test_searchByMatrix(t *testing.T) {
	matrix := []string{"dfadfad", "fdadfadf", "dfdsopadf", "fdadfadf"}
	fmt.Println(searchByMatrix(matrix, "sopa"))
}

func Test_SearchWordInMatrixHorizontal(t *testing.T) {
	matrix := []string{"dfadfad", "fdadfadf", "dfdsopadf", "fdadfadf"}
	fmt.Println(searchByMatrix(matrix, "sopa"))
}

func Test_SearchWordInMatrixVertical(t *testing.T) {
	matrix := []string{"abcdfs",
		"abcdfo",
		"abcdfp",
		"abcdfa",
	}
	fmt.Println(SearchWordInMatrix(matrix, "sopa"))
}

func Test_transposeVertial(t *testing.T) {
	matrix := []string{"abcdfs",
		"abcdfo",
		"abcdfa",
		"abcdfp",
	}

	fmt.Println(transposeVertical(matrix))

}
