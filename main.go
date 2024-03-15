package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter expression (例： 5 + 3), or type 'exit' to quit:")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			fmt.Println("Exiting calculator.")
			os.Exit(0)
		}

		result, err := evaluateExpression(input)
		if err != nil {
			fmt.Println("エラー:", err)
		} else {
			fmt.Println("結果:", result)
		}
	}
}

func evaluateExpression(input string) (float64, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid expression format")
	}

	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
	}

	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
	}

	operator := parts[1]
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
}
