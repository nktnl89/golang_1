package search

type MySimpleSearch struct{}

func (s MySimpleSearch) FindSimpleNumbers(sequence []int) []int {
	simpleNumbers := []int{1, 2, 3}
	for _, num := range sequence {
		isSimple := true
		for _, simple := range simpleNumbers {
			if num%simple == 0 && simple != 1 {
				isSimple = false
				break
			}
		}
		if isSimple {
			simpleNumbers = append(simpleNumbers, num)
		}
	}
	return simpleNumbers
}
