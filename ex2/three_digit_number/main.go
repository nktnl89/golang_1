package main

import "fmt"

func main() {
	var number int16
	fmt.Println("Введите 3-х значное число:")
	fmt.Scanln(&number)

	fmt.Printf("Количество сотен: %d\n", number/100)
	fmt.Printf("Количество десятков: %d\n", number/10%10)
	fmt.Printf("Количество единиц: %d\n", number%10)
}
