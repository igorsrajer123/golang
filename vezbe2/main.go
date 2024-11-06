package main

import (
	"math/cmplx"
)

func findArgAndModulusOfZ(z complex128) (float64, float64) {
	return cmplx.Phase(z), cmplx.Abs(z)
}

func main() {
	// z := 2 + 2i
	// arg, modul := findArgAndModulusOfZ(z)
	// fmt.Println("Complex number module and arg: ", arg, modul)

	// var args = os.Args
	// fmt.Println("Length of args:", len(args))

	// for each loop - 1st element is index of element, second is an
	// for _, arg := range args {
	// 	fmt.Println(arg)
	// }

	// Structure exercise
	// triangle := Triangle{a: 5, b: 6, c: 3}
	// fmt.Println("Triangle Surface is:", calculateTriangleSurface(triangle))

	// File O/I
	// file := openAndReadFile("nekinovifajl.txt")
	// writeToFile(file, "hehe neki text!\n")
	// defer file.Close()

	// Sorting algorithm
	// sampleArray := []int{11, 14, 3, 8, -18, 43}
	// fmt.Println(sampleArray)
	// fmt.Println(bubbleSort(sampleArray))

	// Maps exercise
	// students := []Student{{firstname: "Aleksandar", lastname: "Srajer", year: "2024", rating: 7.55}}
	// students = append(students, newStudent)
	students := initializeStudentsMap()
	for true {

		value := displayMenu()
		switch value {
		case 1:
			displayStudents(students)
		case 2:
			addStudentToMap(students)
		case 3:
			removeStudentFromMap(students)
		case 4:
			return
		default:
			displayMenu()
		}
	}
}
