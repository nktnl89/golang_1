package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Print("Введите число, до которого нужно найти все простые числа: ")
	fmt.Scanln(&a)

	fmt.Println("Найдены простые числа:")
	if a < 0 {
		fmt.Println("Работает только для положительных чисел")
	} else if a == 1 {
		fmt.Println([1]int{1})
	} else if a == 2 {
		fmt.Println([2]int{1, 2})
	} else if a == 3 {
		fmt.Println([3]int{1, 2, 3})
	} else {
		fmt.Println(findSimpleNumbers(generateSequanceNumbers(a)))
	}
}

func findSimpleNumbers(sequence []int) []int {
	simpleNumbers := []int{1, 2, 3}
	for _, num := range sequence {
		isSimple := true
		for _, simple := range simpleNumbers {
			if num%simple == 0 && simple != 1 {
				isSimple = false
				break
			}
		}
		if isSimple {
			simpleNumbers = append(simpleNumbers, num)
		}
	}
	return simpleNumbers
}

func generateSequanceNumbers(limit int) []int {
	var result []int
	for i := 4; i <= int(limit); i++ {
		result = append(result, i)
	}
	return result
}
