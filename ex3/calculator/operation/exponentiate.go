package operation

func Exponentiate(a, b float32) float32 {
	var exp float32

	if b < 0 {
		return exp
	}

	if b == 0 {
		return 1.0
	}

	for i := 1; i < int(b); i++ {
		exp = exp + a*a
	}
	return exp
}
