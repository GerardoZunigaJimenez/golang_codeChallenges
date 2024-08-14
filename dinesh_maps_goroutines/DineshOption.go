package main

import (
	"fmt"
	"os"
	"os/signal"
)

type KeyValue struct {
	Key   string
	Value string
}

var output chan KeyValue

var quit = make(chan os.Signal, 1)

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
	finalMap := map[string]string{}
	output = make(chan KeyValue, len(m))
	go removeMapDuplicatesWithChannel(m)

	for {
		select {
		case kv := <-output:
			finalMap[kv.Key] = kv.Value
		case <-quit:
			close(output)
			close(quit)
			break
		default:
			//nothing
		}
	}

	fmt.Println("Final Map", finalMap)
}

func removeMapDuplicatesWithChannel(m map[string]string) {
	m2 := map[string]string{}

	for k, v := range m {
		if _, ok := m2[v]; ok {
			output <- KeyValue{k, v}
			continue
		}
		m2[v] = k
	}
	signal.Stop(quit)
}
