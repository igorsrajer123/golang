package main

import (
	"bufio"
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
		fmt.Println(scanner.Text())
	}

	return file
}

// Write text to file
func writeToFile(file *os.File, text string) {
	writer := bufio.NewWriter(file)
	_, err := writer.WriteString(text)
	if err != nil {
		log.Fatal(err)
	}

	writer.Flush()
}
