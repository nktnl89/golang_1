package main

import (
	"fmt"
	"math"
)

func main() {
	var circleSquare float64
	fmt.Println("Введите площадь круга:")
	fmt.Scanln(&circleSquare)

	var radius = math.Sqrt(circleSquare / math.Pi)
	var length = 2 * math.Pi * radius
	fmt.Printf("Диаметр окружности = %f\n", radius*2)
	fmt.Printf("Длина окружности = %f\n", length)
}
