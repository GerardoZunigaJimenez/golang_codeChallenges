package main

import "fmt"

func main(){
	start:= 1612224000000//1612224000000
	end:=   1612483200000//1612483200000


	dates := map[int]int{
		99325657:1612493942000,
		99325655:1612406624000,
		99325654:1612233716000,
		99325656:1612233866000,
	}

	for k,v := range dates{
		isRange, text := inRange(start,end,v)
		fmt.Println( "the order ", k, " is in range ",isRange, text , " with the value" , v)
	}
		fmt.Println("\n\n\n")
}

func inRange(start, end, date int) (bool,string){
	var mayorStartDate, minorEndDate bool
	if date >= start {
		mayorStartDate = true
	}
	if date < end {
		minorEndDate = true
	}
	return mayorStartDate&&minorEndDate,fmt.Sprint(" with mayorStartDate ",mayorStartDate," and minorEndDate ",minorEndDate)
}