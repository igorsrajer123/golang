package main

func bubbleSortStudents(students []Student) []Student {
	for i := 0; i < len(students)-1; i++ {
		for j := 0; j < len(students)-i-1; j++ {
			if students[j].Rating < students[j+1].Rating {
				students[j], students[j+1] = students[j+1], students[j]
			}
		}
	}
	return students
}
