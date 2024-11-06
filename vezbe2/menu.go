package main

import (
	"fmt"
	"log"
)

func displayMenu() int {
	fmt.Println("Please choose one of the following options:")
	fmt.Println("1. Display current students.")
	fmt.Println("2. Add new student.")
	fmt.Println("3. Remove student by index.")
	fmt.Println("4. Exit program.")

	var chosenValue int
	_, err := fmt.Scanf("%d", &chosenValue)
	if err != nil {
		log.Fatal(err)
	}

	return chosenValue
}
