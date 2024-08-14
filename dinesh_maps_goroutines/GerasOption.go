package main

import (
	"fmt"
	"sync"
)

const maxNumberOfElementsPerGroup = 5

var out chan map[string]string
var group sync.WaitGroup

func main() {
	m := map[string]string{
		"0001": "jose",
		"0002": "maria",
		"0003": "pepe",
		"0004": "jose",
		"0005": "pepe",
		"0006": "jose",
		"0007": "Gerardo",
		"0008": "pepe",
		"0009": "jose",
		"0010": "pepe",
	}
	l := len(m)
	pages := l / maxNumberOfElementsPerGroup
	out = make(chan map[string]string, pages)


	counter := 0
	tempMap := map[string]string{}
	finalMap := map[string]string{}
	for k, v := range m {
		tempMap[k] = v
		counter++

		if counter == maxNumberOfElementsPerGroup {
			go removeMapDuplicatesToChannel(tempMap)
			counter = 0
			tempMap = map[string]string{}
		}
	}

	select {
	case f := <-out:
		for k, v := range f {
			finalMap[k] = v
		}
	}

	fmt.Println("Final Map", finalMap)

}

func removeMapDuplicates(m map[string]string) map[string]string {
	m2 := map[string]string{}

	for k, v := range m {
		if _, ok := m2[v]; ok {
			continue
		}
		m2[v] = k
	}

	return m2
}

func removeMapDuplicatesToChannel(m map[string]string) {
	out <- removeMapDuplicates(m)
}
