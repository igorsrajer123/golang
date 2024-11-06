package main

import "fmt"

type Student struct {
	Index     string
	Firstname string
	Lastname  string
	Rating    float64
}

func isIndexUnique(students []Student, index string) bool {
	for _, student := range students {
		if index == student.Index {
			return false
		}
	}

	return true
}

func printStudents(students []Student) {
	for _, student := range students {
		fmt.Println(student)
	}
}
