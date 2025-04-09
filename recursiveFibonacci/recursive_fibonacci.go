package recursiveFibonacci

func recursiveFibonacci(fib []int, limit int) []int {
	if fib == nil {
		fib = []int{1}
	}

	l := len(fib)
	if l >= 2 && fib[l-1] > limit {
		return fib[0 : l-1]
	}

	if l == 1 {
		fib = append(fib, 1)
	} else {
		fib = append(fib, fib[l-1]+fib[l-2])
	}

	return recursiveFibonacci(fib, limit)
}

func Fibonacci(limit int) []int {
	//adding 0
	r := recursiveFibonacci(nil, limit)
	r = append([]int{0}, r...)
	return r
}

func recursiveFibonacciMax(fib []int, limit int) []int {
	if fib == nil {
		fib = []int{0, 1}
	}

	l := len(fib)
	sum := fib[l-2] + fib[l-1]
	if sum > limit {
		return fib
	}

	fib = append(fib, sum)
	return recursiveFibonacciMax(fib, limit)
}
