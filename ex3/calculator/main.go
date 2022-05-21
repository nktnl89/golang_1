package main

import (
	"fmt"
	"golang_1/ex3/calculator/operation"
	"os"
)

func main() {
	var a, b float32
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+,-,*,/,**-возведение в степень,!-факториал): ")
	fmt.Scanln(&op)

	fmt.Printf("Результат выполнения операции: %.2f\n", processOperation(a, b, op))
}

func processOperation(a, b float32, op string) float32 {
	var res float32
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя")
			os.Exit(1)
		}
		res = a / b
	case "**":
		if b < 0 {
			fmt.Println("Степень должна быть больше 0")
			os.Exit(1)
		}
		res = operation.Exponentiate(a, b)
	case "!":
		res = operation.Factorial(a)
	default:
		fmt.Println("Операция выбрана неверно.")
		os.Exit(1)
	}
	return res
}
