package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число, до которого нужно рассчитать последовательность Фибоначчи:")
	fmt.Scanln(&num)

	for i := 0; i <= num; i++ {
		fmt.Printf("%d: %d\n", i, fibonachiNum(i))
	}
}

func fibonachiNum(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fibonachiNum(num-1) + fibonachiNum(num-2)
}
