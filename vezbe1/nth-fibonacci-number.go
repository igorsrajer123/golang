package main

func nthFibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	first, second := 0, 1
	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}

	return second
}

func nthFibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return nthFibonacciRecursive(n-1) + nthFibonacciRecursive(n-2)
}
