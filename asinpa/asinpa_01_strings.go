package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println( stringifUserIDsOriginal([]string{"abcd","zxy","1z5s6","fders"}) )
	fmt.Println( stringifUserIDsFor([]string{"abcd","zxy","1z5s6","fders"}) )
	fmt.Println( stringifUserIDs([]string{"abcd","zxy","1z5s6","fders"}) )
}

func stringifUserIDs(IDs []string) (s string){
	return fmt.Sprintf("[\"%s\"]",strings.Join(IDs, "\",\""))
}

func stringifUserIDsFor(IDs []string) string{
	s := "["
	for _,id := range IDs{
		s += fmt.Sprintf("\"%s\",",id)
	}
	s = s[:len(s)-1]
	s+="]"
	return s
}

func stringifUserIDsOriginal(IDs []string) string{
	return "[\"" + strings.Join(IDs, "\",\"") + "\"]"
}
