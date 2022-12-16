package lyft

//mibSubstring("abc","abc")

func minSubstring(fetch, input string) string {
	fMap := createMap(fetch)
	iMap := createMapPosition(input)

	bgnW := len(input) - 1
	endW := 0
	for lttr, qty := range fMap {
		positions, ok := iMap[lttr]
		// the letters from the fetch string doesn't have
		// the same number of iterations on the input string
		lenPsts := len(positions)
		if !ok || lenPsts < qty {
			return ""
		}
		//New End Position
		nbgn := positions[lenPsts-1]
		nEnd := positions[lenPsts-qty]

		if nEnd < bgnW {
			bgnW = nEnd
		}
		if nbgn > endW {
			endW = nbgn
		}
	}

	return input[bgnW : endW+1]
}

func createMap(s string) map[string]int {
	rMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		rMap[s[i:i+1]] = rMap[s[i:i+1]] + 1
	}
	return rMap
}

func createMapPosition(s string) map[string][]int {
	rMap := make(map[string][]int)
	for i := 0; i < len(s); i++ {
		rMap[s[i:i+1]] = append(rMap[s[i:i+1]], i)
	}
	return rMap
}
