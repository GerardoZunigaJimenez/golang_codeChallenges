package codigomx

func GetNotDuplicatedValue(input []int) (r int) {
	dm := make(map[int]int)

	for _, v := range input {
		if _, ok := dm[v]; ok {
			delete(dm, v)
		} else {
			dm[v] = 1
		}
	}
	for k := range dm {
		r = k
		break
	}

	return r
}
