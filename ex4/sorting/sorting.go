package sorting

func SortInt(numbers *[]int) {

	for i := 1; i < len(*numbers); i++ {
		x := (*numbers)[i]
		j := i
		for ; j >= 1 && (*numbers)[j-1] > x; j-- {
			(*numbers)[j] = (*numbers)[j-1]
		}
		(*numbers)[j] = x
	}

}
