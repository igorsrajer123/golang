package main

import "fmt"

func initializeMenu() int {
	fmt.Println("Please chose one of the following options:")
	fmt.Println("1. Add new student")
	fmt.Println("2. List students")
	fmt.Println("3. Remove student by index")
	fmt.Println("4. Generate random numbers for students")
	fmt.Println("5. Show avg. students rating")
	fmt.Println("6. Exit program")

	var chosenOption int
	fmt.Scanf("%d", &chosenOption)

	return chosenOption
}
