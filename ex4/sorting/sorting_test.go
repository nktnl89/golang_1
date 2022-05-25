package sorting

import "github.com/stretchr/testify/assert"
import "testing"

func TestSortInt(t *testing.T) {
	// given:
	a := []int{5, 4, 3, 2, 1}

	// when:
	SortInt(&a)

	// then:
	assert.Equal(t, 1, a[0], "Первый элемент должен быть 1")
	assert.Equal(t, 2, a[1], "Первый элемент должен быть 2")
	assert.Equal(t, 3, a[2], "Первый элемент должен быть 3")
	assert.Equal(t, 4, a[3], "Первый элемент должен быть 4")
	assert.Equal(t, 5, a[4], "Первый элемент должен быть 5")
}
