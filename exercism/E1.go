package main

import "fmt"

func IsAnagram(s string) bool{

	var s1, s2 string
	for _,v := range s{
		ss := string(v)
		if ss != " "{
			s1 += string(v)
		}
	}
	for i:= len(s1)-1; i>=0; i--{
		s2 += string(s1[i])
	}

	if s1 == s2{
		return true
	}

	return false
}

func main(){
	test1 := "A N A"  // Valid
	test2 := "GE ar do"  //Invalid
	test3 := "%3123GEse s" //Invalid
	test4 := "GEZE1 233 21EZEG"  //Valid

	fmt.Println("Is a valid Anagagram" , test1 , " - ", IsAnagram(test1))
	fmt.Println("Is a valid Anagagram" , test2 , " - ", IsAnagram(test2))
	fmt.Println("Is a valid Anagagram" , test3 , " - ", IsAnagram(test3))
	fmt.Println("Is a valid Anagagram" , test4 , " - ", IsAnagram(test4))

}