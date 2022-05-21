package search

import "math/big"

type StandardSearch struct{}

func (s StandardSearch) FindSimpleNumbers(sequence []int) []int {
	simpleNumbers := []int{1, 2, 3}
	for _, num := range sequence {
		bigNum := big.NewInt(int64(num))
		isSimple := bigNum.ProbablyPrime(1)
		if isSimple {
			simpleNumbers = append(simpleNumbers, num)
		}
	}
	return simpleNumbers
}
