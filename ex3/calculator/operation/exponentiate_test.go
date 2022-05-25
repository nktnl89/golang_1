package operation

import "testing"

func TestExponentiate(t *testing.T) {
	testCases := []struct {
		name        string
		firstValue  float32
		secondValue float32
		result      float32
	}{
		{
			"Степень 2 от 2",
			2.0,
			2.0,
			4.0,
		},
		{
			"Степень 0 от 2",
			2.0,
			0.0,
			1.0,
		},
		{
			"Степень -1 от 2",
			2.0,
			-1.0,
			0.0,
		},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.name, func(t *testing.T) {
			res := Exponentiate(cse.firstValue, cse.secondValue)

			if res != cse.result {
				t.Fatalf("result incorrect, got %.2f, expected %.2f", res, cse.result)
			}
		})
	}

}
