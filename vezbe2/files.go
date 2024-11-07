package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func openAndReadFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		student := Student{}
		err = json.Unmarshal([]byte(scanner.Text()), &student)
		fmt.Println(student)
	}

	return file
}

// Write text to file
func writeToFile(file *os.File) {
	student := Student{Firstname: "pera", Lastname: "peric", Year: "2029", Rating: 6.7}
	writer := bufio.NewWriter(file)
	jsonStudent, err := json.Marshal(student)
	if err != nil {
		log.Fatal(err)
	}
	_, err = writer.WriteString(string(jsonStudent) + "\n")

	writer.Flush()
}
