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
			channel := make(chan int)
			students := getAllStudents(filename)

			for i := 0; i < len(students); i++ {
				go generateRandomNumber(channel)
			}

			for i := range students {
				number := <-channel
				students[i].RandomNumber = number
			}

		case 5:
			fmt.Println("AAA" + string(chosenOption))
		default:
			return
		}
	}
}
