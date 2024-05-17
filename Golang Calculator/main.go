package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		fmt.Println("Welcome to My Console Calculator.")
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
		fmt.Print("Enter the first number: ")
		num1, err := getUserInput()
		if err != nil {
			fmt.Println("Invalid Input, Please Try Again")
			continue
		}
		var result float64
		var operation string

		for {
			fmt.Print("Select an Operation (+, -, *, /, or =): ")
			fmt.Scanln(&operation)

			if operation == "=" {
				break
			} else if operation == "+" || operation == "-" || operation == "*" || operation == "/" {
				fmt.Print("Input Next Number: ")
				num2, err := getUserInput()
				if err != nil {
					fmt.Println("Invalid Input, Please Try Again")
					continue
				}

				switch operation {
				case "+":
					result = add(num1, num2)
				case "-":
					result = subtract(num1, num2)
				case "*":
					result = multiply(num1, num2)
				case "/":
					result, err = divide(num1, num2)
					if err != nil {
						fmt.Println(err)
						continue
					}
				}

				num1 = result
			} else {
				fmt.Println("Invalid operation")
				continue
			}
		}

		fmt.Println("Result:", result)

		var restart string
		fmt.Print("Do you want to calculate again? (yes/no): ")
		fmt.Scan(&restart)

		if restart != "yes" {
			break
		}

	}
}

func getUserInput() (float64, error) {
	var input string
	fmt.Scanln(&input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return x / y, nil
}
