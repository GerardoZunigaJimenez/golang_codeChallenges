package main

import "fmt"

func main(){
	var input =	[]interface{}{1, {2,3,nil,4},"Gerardo",[]{nil},5}
	output := FlattenArray(input)
	fmt.Println(output)
}

func FlattenArray(s []interface{}) []interface{}{
	result := make ([]interface{}, 0,len(s))

	for _,v := range s{

		if v != nil || v != "" {
			result = append( result, v)
		}
	}

	return result
}
