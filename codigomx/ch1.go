package codigomx

func CheckZeros(input []int) (result []int) {
	result = make([]int, 0)
	zeroS := make([]int, 0)
	for _, v := range input {
		if v == 0 {
			zeroS = append(zeroS, 0)
			continue
		}
		result = append(result, v)
	}
	result = append(zeroS, result...)

	return result
}
