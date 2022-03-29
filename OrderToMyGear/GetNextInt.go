package main

import (
"fmt"
"sort"
)

/*
You're running a pool of servers where the servers are numbered sequentially starting from 1. Over time, any given server might explode, in which case its server number is made available for reuse. When a new server is launched, it should be given the lowest available number.

Write a function which, given the list of currently allocated server numbers, returns the number of the next server to allocate. In addition, you should demonstrate your approach to testing that your function is correct. You may choose to use an existing testing library for your language if you choose, or you may write your own process if you prefer.

For example, your function should behave something like the following:

>> next_server_number([5, 3, 1])
2
>> next_server_number([5, 4, 1, 2])
3
>> next_server_number([3, 2, 1])
4
>> next_server_number([2, 3])
1
>> next_server_number([])
1
*/
func main(){
	fmt.Println(next_server_number([]int{5,3,1}))
	fmt.Println(next_server_number([]int{5,4,1,2}))  //2
	fmt.Println(next_server_number([]int{3,2,1}))
	fmt.Println(next_server_number([]int{2,3}))
	fmt.Println(next_server_number([]int{}))

	fmt.Println()

	fmt.Println(next_server_number2([]int{5,3,1}))
	fmt.Println(next_server_number2([]int{5,4,1,2}))  //2
	fmt.Println(next_server_number2([]int{3,2,1}))
	fmt.Println(next_server_number2([]int{2,3}))
	fmt.Println(next_server_number2([]int{}))

	fmt.Println()

	fmt.Println(next_server_number3([]int{5,3,1}))
	fmt.Println(next_server_number3([]int{5,4,1,2}))  //2
	fmt.Println(next_server_number3([]int{3,2,1}))
	fmt.Println(next_server_number3([]int{2,3}))
	fmt.Println(next_server_number3([]int{}))


	fmt.Println()

	fmt.Println(next_server_number4([]int{5,3,1}))
	fmt.Println(next_server_number4([]int{5,4,1,2}))  //2
	fmt.Println(next_server_number4([]int{3,2,1}))
	fmt.Println(next_server_number4([]int{2,3}))
	fmt.Println(next_server_number4([]int{}))
}


func next_server_number(b []int) int {
	s := len(b)

	//empty slice
	if s == 0 {
		return 1
	}

	sort.Ints( b )

	if b[0]>1 {
		return 1
	}

	for k, _ := range b{
		// We finish the array, we are in the last element
		if k == s-1 {
			break
		}
		//We are comparing the values of the current k and the next slice element,
		// to check if there is number between them
		if k < s && b[k]+1 == b[k+1]{
			continue
		} else {
			return b[k]+1
		}
		k++
	}

	//This scenario is for any other value of the sliced satisfies the value
	return b[s-1]+1
}

func next_server_number2(b []int) int {
	s := len(b)

	//empty slice
	if s == 0 {
		return 1
	}

	sort.Ints( b )

	if b[0]>1 {
		return 1
	}

	for k := 0; k < s-1; k++{
		//We are comparing the values of the current k and the next slice element,
		// to check if there is number between them
		if b[k]+1 != b[k+1]{
			return b[k]+1
		}
	}

	//This scenario is for any other value of the sliced satisfies the value
	return b[s-1]+1
}

func next_server_number3(b []int) int {
	s := len(b)

	sort.Ints( b )

	for i := 0; i<s; i++{
		if b[i] != (i+1){
			return (i+1)
		}
	}
	return s+1
}

func next_server_number4(servers []int) int {
	for i := 1; i <= len(servers); i++ {
		exists := contains(servers, i)

		if exists == false {
			return i
		}
	}

	return len(servers) + 1
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

