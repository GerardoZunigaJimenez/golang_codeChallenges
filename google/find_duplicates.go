// To execute Go code, please declare a func main() in a package "main"
/**
You are given a list of elements. Each element has a unique id and 3 properties. Two elements are considered as duplicate if they share any of the 3 properties. Please write a function that takes the input and return all the duplicates.


Sample of inputs:
E1: id1, p1, p2, p3
E2: id2, p1, p4, p5
E3: id3, p6, p7, p8

Example #1:
input: [[id1, p1, p2, p3], [id2, p1, p4, p5], [id3, p6, p7, p8]]

Expected outcome: [[id1, id2], [id3]]

Example #2:
input : [[id1, p1, p2, p3], [id2, p1, p4, p5], [id3, p6, p7, p8], [id4, p9, p2, p10]]

Expected outcome: [["id1", "id2", "id4"], ["id3"]]

Note that order of id does not matter in the outcome. thus answer  [[id2, id1], id3] is also accepted. Also it is guaranteed that each element won’t have duplicate property (ex. E4: id4, p1, p1, p2)


elements := [["id1", "p1", "p2", "p3"], ["id2", "p1", "p4", "p5"], ["id3", "p6", "p7", "p8"], ["id4", "p9", "p2", "p10"]]
function find_duplicate(elements Array<Array<string>>) {

}

*/

package google

import "fmt"

func main() {
	fmt.Println(findDuplicate([][]string{[]string{"id1", "p1", "p2", "p3"}, []string{"id2", "p1", "p4", "p5"}, []string{"id3", "p6", "p7", "p8"}}))
	fmt.Println(findDuplicate([][]string{[]string{"id1", "p1", "p2", "p3"}, []string{"id2", "p1", "p4", "p5"}, []string{"id3", "p6", "p7", "p8"}, []string{"id4", "p6"}, []string{"id6", "p10", "p12", "p20"}, []string{"id7", "p8"}}))
}

func findDuplicate(elements [][]string) (r [][]string) {
	ids := map[string]bool{}
	properties := map[string][]string{}

	for _, e := range elements {
		// id is first position of e
		id := e[0]
		ids[id] = false
		for i := 1; i < len(e); i++ {
			if s, ok := properties[e[i]]; ok {
				properties[e[i]] = append(s, id)
			} else {
				properties[e[i]] = []string{id}
			}
		}
	}

	duplicated := []string{}
	for _, v := range properties {
		if len(v) > 1 {
			for _, vv := range v {
				duplicated = append(duplicated, vv)
				if _, ok := ids[vv]; ok {
					delete(ids, vv)
				}
			}

		}
	}
	r = append(r, duplicated)
	for k, _ := range ids {
		r = append(r, []string{k})
	}

	return r
}

func findDuplicate2(elements [][]string) (r [][]string) {
	ids := map[string]bool{}
	idsWithDuplicates := []string{}
	properties := map[string][]string{}

	for _, e := range elements {
		// id is first position of e
		id := e[0]
		ids[id] = false
		for i := 1; i < len(e); i++ {
			if s, ok := properties[e[i]]; ok {
				properties[e[i]] = append(s, id)
				idsWithDuplicates = append(properties[e[i]])
			} else {
				properties[e[i]] = []string{id}
				delete(ids, id)
			}
		}
	}

	r = append(r, idsWithDuplicates)
	for id, _ := range ids {
		r = append(r, []string{id})
	}

	return r
}

func findDuplicate3(elements [][]string) (r [][]string) {
	ids := map[string]bool{}
	properties := map[string]map[string]bool{}
	for _, e := range elements {
		// id is first position of e
		id := e[0]
		ids[id] = false
		for i := 1; i < len(e); i++ {
			if sM, ok := properties[e[i]]; ok {
				sM[id] = false
				properties[e[i]] = sM
				for kkk, _ := range sM {
					delete(ids, kkk)
				}
			} else {
				sM = map[string]bool{id: false}
				properties[e[i]] = sM
			}
		}
	}

	for _, m := range properties {
		r = append(r)
		if len(m) > 1 {
			var s []string
			for p, _ := range m {
				s = append(s, p)
			}
			r = append(r, s)
		}
	}
	for id, _ := range ids {
		r = append(r, []string{id})
	}

	return r
}
