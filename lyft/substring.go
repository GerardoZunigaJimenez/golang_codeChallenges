package lyft

import (
	"sort"
)

//mibSubstring("abc","abc")

func minSubstring(fetch, input string) string {
	fMap, fKeys := createMap(fetch)
	iSet, iKeys := createSetPosition(input)

	if len(fKeys) != len(iKeys) {
		return ""
	}
	for i, _ := range fKeys {
		if fKeys[i] != iKeys[i] {
			return ""
		}
	}

	bgnW := len(input) - 1
	endW := 0
	for _, order := range iSet {
		lttr := order.Key
		qty, ok := fMap[lttr]
		if !ok {
			continue
		}
		positions := order.Positions
		// the letters from the fetch string doesn't have
		// the same number of iterations on the input string
		lenPsts := len(positions)
		if lenPsts < qty {
			return ""
		}
		//New End Position
		nbgn := positions[0]
		nEnd := positions[lenPsts-1]

		if nEnd <= bgnW {
			bgnW = nEnd
		}
		if nbgn >= endW {
			endW = nbgn
		}
	}

	return input[bgnW : endW+1]
}

func createMap(s string) (map[string]int, []string) {
	rMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		lttr := s[i : i+1]
		rMap[lttr] = rMap[lttr] + 1
	}
	var keys []string
	for k := range rMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return rMap, keys
}

type eSet struct {
	Key       string
	Order     int
	Positions []int
}

type sortByOrder []eSet

func (a sortByOrder) Len() int           { return len(a) }
func (a sortByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByOrder) Less(i, j int) bool { return a[i].Order < a[j].Order }

func createSetPosition(s string) ([]eSet, []string) {
	rMap := make(map[string]eSet)

	for i, o := 0, 1; i < len(s); i++ {
		lttr := s[i : i+1]
		set, ok := rMap[lttr]
		if !ok {
			rMap[lttr] = eSet{
				Key:       lttr,
				Order:     o,
				Positions: []int{i},
			}
			o++
		} else {
			rMap[lttr] = eSet{
				Key:       set.Key,
				Order:     set.Order,
				Positions: append(set.Positions, i),
			}
		}
	}

	var rs []eSet
	var keys []string
	for k, v := range rMap {
		rs = append(rs, v)
		keys = append(keys, k)
	}

	sort.Sort(sortByOrder(rs))
	sort.Strings(keys)
	return rs, keys
}
