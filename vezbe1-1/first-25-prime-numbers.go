package main

import "fmt"

func firstNPrimeNumbers(n int) {
	fmt.Println("---------------------CALCULATE FIRST N PRIME NUMBERS---------------------")

	count := 0
	num := 2

	for count < n {
		if isPrimeNumber(num) {
			count++
			fmt.Println("Prime number:", num)
		}
		num++
	}

}
