package main

import (
	"fmt"
	"log"
)

type Student struct {
	Firstname string
	Lastname  string
	Year      string
	Rating    float64
}

func initializeStudentsMap() map[string]Student {
	igor := Student{Firstname: "Igor", Lastname: "Srajer", Year: "2024", Rating: 8.85}
	students := make(map[string]Student)
	students["ra78/2024"] = igor

	return students
}

func displayStudents(students map[string]Student) {
	fmt.Println("These are the current students:")
	for key, student := range students {
		fmt.Printf("Student index %s: %v\n", key, student)
	}
}

func addStudentToMap(studentMap map[string]Student) {
	newStudent := Student{}

	fmt.Println("Please enter new student F:")
	_, err := fmt.Scanf("%s", &newStudent.Firstname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter new student lastname:")
	_, err = fmt.Scanf("%s", &newStudent.Lastname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter new student year:")
	_, err = fmt.Scanf("%s", &newStudent.Year)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter new student rating:")
	_, err = fmt.Scanf("%g", &newStudent.Rating)
	if err != nil {
		log.Fatal(err)
	}

	studentMap["ra79/2024"] = newStudent
}

func removeStudentFromMap(studentMap map[string]Student) {
	fmt.Println("Please enter student index:")
	var index string
	_, err := fmt.Scanf("%s", &index)
	if err != nil {
		log.Fatal(err)
	}

	delete(studentMap, index)
}
