package lyft

import (
	"sort"
)

type eSet struct {
	Key       string
	Order     int
	Positions []int
	Quantity  int
}

//mibSubstring("abc","abc")
func minSubstring(fetch, input string) string {
	fSet, fKeys := createSetPosition(fetch)
	iMap, iKeys := createMapPosition(input)

	fKeysLng := len(fKeys)
	iKeysLng := len(iKeys)

	// The letters from fetch are not existing from fetch to input
	if fKeysLng != iKeysLng {
		return ""
	}
	// The letters from fetch are not existing from fetch to input
	for i, _ := range fKeys {
		if fKeys[i] != iKeys[i] {
			return ""
		}
	}

	bgnW := len(input) - 1
	endW := 0
	for i := 0; i < fKeysLng; i++ {
		fObj := fSet[i]
		fLttr := fObj.Key
		fQty := fObj.Quantity

		iObj := iMap[fLttr]
		iPositions := iObj.Positions
		iObjPstsLng := len(iPositions)
		if iObjPstsLng < fQty {
			return ""
		}
		//New End Position
		fBgn := iPositions[iObj.Quantity-1]
		fEnd := iPositions[0]

		if fBgn < bgnW {
			bgnW = fBgn
		}
		if fEnd > endW {
			endW = fEnd
		}
	}

	return input[bgnW : endW+1]
}

type sortByOrder []eSet

func (a sortByOrder) Len() int           { return len(a) }
func (a sortByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByOrder) Less(i, j int) bool { return a[i].Order < a[j].Order }

type sortByQuantity []eSet

func (a sortByQuantity) Len() int           { return len(a) }
func (a sortByQuantity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByQuantity) Less(i, j int) bool { return a[i].Quantity < a[j].Quantity }

func createSetPosition(s string) (rs []eSet, keys []string) {
	rMap := make(map[string]eSet)

	for i, o := 0, 1; i < len(s); i++ {
		lttr := s[i : i+1]
		set, ok := rMap[lttr]
		if !ok {
			newSet := eSet{
				Key:       lttr,
				Order:     o,
				Positions: []int{i},
				Quantity:  1,
			}
			rMap[lttr] = newSet
			o++
		} else {
			rMap[lttr] = eSet{
				Key:       set.Key,
				Order:     set.Order,
				Positions: append(set.Positions, i),
				Quantity:  set.Quantity + 1,
			}
		}
	}

	for _, v := range rMap {
		rs = append(rs, v)
		keys = append(keys, v.Key)
	}

	sort.Sort(sortByQuantity(rs))
	sort.Strings(keys)
	return rs, keys
}

func createMapPosition(s string) (rm map[string]eSet, keys []string) {
	rm = make(map[string]eSet)

	for i, o := 0, 1; i < len(s); i++ {
		lttr := s[i : i+1]
		set, ok := rm[lttr]
		if !ok {
			rm[lttr] = eSet{
				Key:       lttr,
				Order:     o,
				Positions: []int{i},
				Quantity:  1,
			}
			keys = append(keys, lttr)
			o++
		} else {
			rm[lttr] = eSet{
				Key:       set.Key,
				Order:     set.Order,
				Positions: append(set.Positions, i),
				Quantity:  set.Quantity + 1,
			}
		}
	}

	sort.Strings(keys)
	return rm, keys
}
