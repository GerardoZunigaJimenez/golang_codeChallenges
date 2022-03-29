package main

import "fmt"

func main(){
	m := map[string]string{"field":"value"}

	if v,ok := m["fieldd"]; ok{
		fmt.Printf("Existio '%s' %t", v, ok)
	} else {
		fmt.Printf("No Existio '%s' %t", v, ok)
	}

}
