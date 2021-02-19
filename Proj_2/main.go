package main

import "fmt"

func main() {
	var num1, num2 int
	var operator string
	fmt.Println("Enter the first Number")
	fmt.Scanln(&num1)
	fmt.Println("Enter the Second Number")
	fmt.Scanln(&num2)
	fmt.Println("Enter the Operator")
	fmt.Scanln(&operator)

	output := 0

	switch operator {
	case "+":
		output = num1 + num2
		break
	case "-":
		output = num1 - num2
		break
	case "/":
		output = num1 / num2
		break
	case "*":
		output = num1 * num2
		break
	case "%":
		output = num1 % num2
		break
	}

	fmt.Printf("%d %s %d = %d\n", num1, operator, num2, output)
}
