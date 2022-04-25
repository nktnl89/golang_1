package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число, до которого нужно рассчитать последовательность Фибоначчи:")
	fmt.Scanln(&num)

	var fibonachiHash = make(map[int]int)
	fibonachiHash[0] = 0
	fibonachiHash[1] = 1

	for i := 0; i <= num; i++ {
		fmt.Printf("%d: %d\n", i, fibonachiNum(i, fibonachiHash))
	}

	fmt.Println(fibonachiHash)
}

func fibonachiNum(num int, hash map[int]int) int {
	hashValue, ok := hash[num]
	if ok {
		return hashValue
	} else {
		hash[num] = fibonachiNum(num-1, hash) + fibonachiNum(num-2, hash)
	}
	return hash[num]
}
