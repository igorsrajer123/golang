package main

import "fmt"

func initializeMenu() int {
	fmt.Println("Please choose one of the following options:")
	fmt.Println("1. List all students")
	fmt.Println("2. Add new student")
	fmt.Println("3. Exit program")

	var selectedOption int
	fmt.Scanf("%d", &selectedOption)

	return selectedOption
}
