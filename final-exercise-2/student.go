package main

import "strconv"

type IStudent interface {
	toString() string
}

type Student struct {
	Index     string
	Firstname string
	Lastname  string
	Rating    float64
}

func (student Student) toString() string {
	return "Student: " + student.Index + "/" + student.Firstname + "/" + student.Lastname + "/" + strconv.FormatFloat(student.Rating, 'f', -1, 64)
}
