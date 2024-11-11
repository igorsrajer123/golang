package main

import "fmt"

func main() {

	filename := "studenti.txt"
	for {
		initializeFile(filename)

		selectedOption := initializeMenu()
		switch selectedOption {
		case 1:
			students := getAllStudents(filename)
			for i := range students {
				fmt.Println(students[i].toString())
			}
		case 2:
			addStudent(filename)
		default:
			return
		}
	}
}
