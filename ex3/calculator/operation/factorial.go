package operation

// Factorial counts factorial of a parameter
func Factorial(a float32) float32 {
	var fact = 1
	for i := 2; i <= int(a); i++ {
		fact = fact * i
	}
	return float32(fact)
}
