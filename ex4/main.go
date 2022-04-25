package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers = make([]int, 0, 10)

	for scanner.Scan() {
		newNumber, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, newNumber)
	}

	for i := 1; i < len(numbers); i++ {
		x := numbers[i]
		j := i
		for ; j >= 1 && numbers[j-1] > x; j-- {
			numbers[j] = numbers[j-1]
		}
		numbers[j] = x
	}
	fmt.Println("Результат: ", numbers)
}
