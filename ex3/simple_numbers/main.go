package main

import (
	"fmt"
	"golang_1/ex3/simple_numbers/search"
	"os"
)

type SimpleNumberSearch interface {
	FindSimpleNumbers(sequence []int) []int
}

func main() {
	var a int

	fmt.Print("Введите число, до которого нужно найти все простые числа: ")
	fmt.Scanln(&a)

	if a < 0 {
		fmt.Println("Работает только для положительных чисел")
		os.Exit(1)
	}

	seq := generateSequanceNumbers(a)
	searchImpls := []SimpleNumberSearch{
		search.MySimpleSearch{},
		search.StandardSearch{},
	}
	for _, s := range searchImpls {
		fmt.Println("Найдены простые числа:")
		if a == 1 {
			fmt.Println([1]int{1})
		} else if a == 2 {
			fmt.Println([2]int{1, 2})
		} else if a == 3 {
			fmt.Println([3]int{1, 2, 3})
		} else {
			fmt.Println(s.FindSimpleNumbers(seq))
		}
	}
}

func generateSequanceNumbers(limit int) []int {
	var result []int
	for i := 4; i <= int(limit); i++ {
		result = append(result, i)
	}
	return result
}
