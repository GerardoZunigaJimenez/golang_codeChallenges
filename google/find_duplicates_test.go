package google

import (
	"fmt"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	//t.Parallel()
	tests := []struct {
		name   string
		input  [][]string
		output [][]string
	}{
		{
			name:   "id3",
			input:  [][]string{[]string{"id1", "p1", "p2", "p3"}, []string{"id2", "p1", "p4", "p5"}, []string{"id3", "p6", "p7", "p8"}},
			output: [][]string{[]string{"id1", "id2"}, []string{"id3"}},
		},
		{
			name:   "id3",
			input:  [][]string{[]string{"id1", "p1", "p2", "p3"}, []string{"id2", "p1", "p4", "p5"}, []string{"id3", "p6", "p7", "p8"}, []string{"id4", "p9", "p2", "p10"}},
			output: [][]string{[]string{"id1", "id2", "id4"}, []string{"id3"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//fmt.Println("findDuplicate1", findDuplicate(tt.input))
			fmt.Println("findDuplicate3", findDuplicate3(tt.input))
		})
	}
}
