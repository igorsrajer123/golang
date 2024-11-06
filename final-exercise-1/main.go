package main

import "fmt"

func main() {
	filename := "students.txt"
	initializeFile(filename)

	for {
		chosenOption := initializeMenu()
		switch chosenOption {
		case 1:
			addStudent(filename)
		case 2:
			students := getAllStudents(filename)
			printStudents(bubbleSortStudents(students))
		case 3:
			removeStudent(filename)
		case 4:
			fmt.Println("AAA" + string(chosenOption))
		case 5:
			fmt.Println("AAA" + string(chosenOption))
		default:
			return
		}
	}
}
