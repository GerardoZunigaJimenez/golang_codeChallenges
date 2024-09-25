package stringparser

import (
	"fmt"
	"testing"
)

func TestFindStringsUnderSlice(t *testing.T) {
	fmt.Println(FindStringsUnderSlice("{\"query\":[\"NW01\"],\"size\":500}"))
	fmt.Println(FindStringsUnderSlice("{\"query\":[\"NW01\",\"NW02\",\"NW\"],\"size\":500}"))
}
