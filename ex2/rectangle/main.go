package main

import "fmt"

func main() {
	var width, height uint

	fmt.Println("Введите ширину:")
	fmt.Scanln(&width)

	fmt.Println("Введите высоту:")
	fmt.Scanln(&height)

	fmt.Printf("Площадь прямоугольника: %d\n", width*height)
}
