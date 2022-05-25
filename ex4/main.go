package main

import (
	"bufio"
	"fmt"
	"golang_1/ex4/sorting"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers = []int{}

	for scanner.Scan() {
		a := scanner.Text()
		if a == "stop" {
			break
		}
		newNumber, _ := strconv.Atoi(a)
		numbers = append(numbers, newNumber)
	}

	sorting.SortInt(&numbers)

	fmt.Println("Результат: ", numbers)
}
