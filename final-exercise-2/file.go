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
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	studentList := []Student{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		student := Student{}
		err := json.Unmarshal([]byte(scanner.Text()), &student)
		if err != nil {
			log.Fatal(err)
		}
		studentList = append(studentList, student)
	}

	file.Close()

	return studentList
}

func addStudent(filename string) {
	newStudent := Student{}

	fmt.Println("Please enter student index:")
	_, err := fmt.Scanf("%s", &newStudent.Index)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter student firstname:")
	_, err = fmt.Scanf("%s", &newStudent.Firstname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter student lastname:")
	_, err = fmt.Scanf("%s", &newStudent.Lastname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please enter student rating:")
	_, err = fmt.Scanf("%g", &newStudent.Rating)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)
	studentBytes, err := json.Marshal(newStudent)
	if err != nil {
		log.Fatal(err)
	}

	_, err = writer.WriteString(string(studentBytes) + "\n")
	if err != nil {
		log.Fatal(err)
	}

	writer.Flush()
	file.Close()
}
