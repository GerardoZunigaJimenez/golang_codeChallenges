package main

import (
	"fmt"
	"sort"
)

// func 2 parameters, []integer and & integer
// return the position of the numbers, that the sum is equals to the input intger
// the 2 position should be different

func main() {
	s := []int{94, 1, -2, 3, -5, 8, 10, 7, 6,11,0}
	target := 7
	p1, p2 := LocatePositionNumberFoSum2(s, target)

	if p1 > -1 && p2 > -1 {
		fmt.Println("The position numbers that sum the taget", target, "are ", s[p1]," - ",s[p2])
	} else {
		fmt.Println("there is not a possible answer for the request")
	}

}

func LocatePositionNumberFoSum(s []int, target int) (int, int) {
	if len(s) > 1 {
		for k, _ := range s {
			for kk:=k+1; kk<len(s); kk++ {
				if k == kk {continue}
				//happy path
				if s[k]+s[kk] == target {
					return k, kk
				}
			}
		}
	}

	// Final scenario, there not possible solution
	return -1, -1
}

func LocatePositionNumberFoSum2(s []int, target int) (int, int) {
	sort.Ints(s)
	if len(s) > 1 {
		for k, v := range s {
			subs := target - v
			for kk := k +1; k<len(s)-1; kk++{
				if s[kk] == subs{
					return k,kk
				}
			}
		}
	}

	// Final scenario, there not possible solution
	return -1, -1
}