package main

import (
	"fmt"
)

func getChosenOperation(operationKey int) string {
	var operation string

	switch operationKey {
	case 1:
		operation = "PLUS"
	case 2:
		operation = "MINUS"
	case 3:
		operation = "MULTIPLY"
	case 4:
		operation = "DIVIDE"
	default:
		operation = "REMAINDER"
	}

	return operation
}

func calculator() {
	fmt.Println("---------------------CALCULATE FUNCTION---------------------")
	var operationKey, x, y, result int

	fmt.Println("Please enter an operation key:")
	n1, err1 := fmt.Scanf("%d", &operationKey)
	if n1 != 1 || err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println("OPERATION CHOSEN:", getChosenOperation(operationKey))

	fmt.Println("Please enter the first number:")
	n2, err2 := fmt.Scanf("%d", &x)
	if n2 != 1 || err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("Please enter the second number:")
	n3, err3 := fmt.Scanf("%d", &y)
	if n3 != 1 || err3 != nil {
		fmt.Println(err3)
		return
	}

	switch operationKey {
	case 1:
		result = x + y
	case 2:
		result = x - y
	case 3:
		result = x * y
	case 4:
		result = x / y
	default:
		result = x % y
	}

	fmt.Println("RESULT: ", result)
}
