package main

import (
	"fmt"
	"runtime"
)

func isEven(value int) bool {
	if value%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var a = 5
	fmt.Printf("%v is an even number: %v\n", a, isEven(a))

	igor := runtime.GOOS
	switch igor {
	case "linux":
		fmt.Println("LINUX")
	default:
		fmt.Println(igor)
	}

	// calculator()
	// fmt.Println(isPrimeNumber(4))
	firstNPrimeNumbers(10)

}
