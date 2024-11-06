package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func initializeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func getAllStudents(filename string) []Student {
	students := []Student{}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		student := Student{}
		err := json.Unmarshal([]byte(scanner.Text()), &student)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}

	file.Close()
	return students
}

func addStudent(filename string) {
	student := Student{}

	fmt.Println("Please enter a unique student index:")
	_, err := fmt.Scanf("%s", &student.Index)
	if err != nil {
		log.Fatal(err)
	}

	students := getAllStudents(filename)
	if !isIndexUnique(students, student.Index) {
		fmt.Println("Entered index is not unique!")
		return
	}

	fmt.Println("Please enter students firstname:")
	_, err = fmt.Scanf("%s", &student.Firstname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter students lastname:")
	_, err = fmt.Scanf("%s", &student.Lastname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter students rating:")
	_, err = fmt.Scanf("%g", &student.Rating)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)
	studentJson, err := json.Marshal(student)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.WriteString(string(studentJson) + "\n")
	writer.Flush()

	file.Close()
}

func removeStudent(filename string) {
	var enteredIndex string
	fmt.Println("Please enter student index:")
	_, err := fmt.Scanf("%s", &enteredIndex)
	if err != nil {
		log.Fatal(err)
	}

	var retVal []Student
	students := getAllStudents(filename)
	for _, student := range students {
		if student.Index != enteredIndex {
			retVal = append(retVal, student)
		}
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	newWriter := bufio.NewWriter(file)
	for _, student := range retVal {
		jsonStudent, err := json.Marshal(student)
		if err != nil {
			log.Fatal(err)
		}
		_, err = newWriter.WriteString(string(jsonStudent) + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	newWriter.Flush()
	file.Close()
}
