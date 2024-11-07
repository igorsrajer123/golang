package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum([]int{2, 2, 2}, c) // 6
	go sum(s, c)              // 12 Redosled pozivanja gorutina je bitan?
	go sum(s[len(s)/2:], c)   // -5
	go sum([]int{5, 5, 5}, c) // 15
	x, y, z, q := <-c, <-c, <-c, <-c

	fmt.Println(x, y, z, q)
}
