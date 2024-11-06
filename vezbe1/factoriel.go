package main

import "fmt"

func factorielIterative(n int) int {
	if n < 0 {
		fmt.Println("Negative number provided.")
		return -1
	}

	if n == 0 || n == 1 {
		return 1
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

func factorielRecursive(n int) int {
	if n < 0 {
		fmt.Println("Negative number provided.")
		return -1
	}

	if n == 0 || n == 1 {
		return 1
	}

	return n * factorielRecursive(n-1)
}
