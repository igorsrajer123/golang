package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"os"
)

func add_v1(x int, y int) int {
	return x + y
}

func add_v2(x, y int) int {
	return x + y + x
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	y = sum - x
	return
}

var c, python, java = 1, 2, "no"
var pascal = "pascal"

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1 // << operacija znaci da smo kreirali binarnu vrednost sa 1 na pocetku i 64 nule iza nje
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("Hello there PACKAGES:", rand.Intn(10))
	fmt.Println("Math squirting:", math.Sqrt(9))
	fmt.Println("Printing exported PI:", math.Pi)

	fmt.Println("A + B V1:", add_v1(5, 5))
	fmt.Println("A + B V2:", add_v2(1, 2))

	a, b := swap("I love", "you")
	fmt.Println("Swap functin call [A]:", a, ", And another one [B]: ", b)

	first, second := split(16)
	fmt.Println("Naked return function [X]:", first, "Naked return function [Y]:", second)

	golang := 5
	delphi, sql := true, "SQL"
	fmt.Println("Variable declaration: ", c, python, java, golang, nil, pascal, delphi, sql)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	sale := os.Args
	fmt.Println("SALE: ", sale)

	var pera bool
	fmt.Println("PERA: ", pera)

	i := math.Round(42.6)
	f := float64(i)
	u := uint(i)
	cpx := 5 + 12i
	fmt.Printf("\nint: %v \nfloat64: %v \nuint: %v \ncomplex type: %T and value: %v", i, f, u, cpx, cpx)

	const igor = "世界"
	fmt.Println(igor)
}
