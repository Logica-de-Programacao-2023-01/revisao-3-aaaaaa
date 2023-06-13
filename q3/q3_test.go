package q3

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n      int
		result int
	}{
		{
			n:      1,
			result: 1,
		},
		{
			n:      2,
			result: 2,
		},
		{
			n:      3,
			result: 3,
		},
		{
			n:      4,
			result: 5,
		},
		{
			n:      5,
			result: 8,
		},
		{
			n:      6,
			result: 13,
		},
		{
			n:      7,
			result: 21,
		},
		{
			n:      8,
			result: 34,
		},
		{
			n:      9,
			result: 55,
		},
		{
			n:      10,
			result: 89,
		},
	}

	for _, test := range tests {
		result := ClimbStairs(test.n)
		if result != test.result {
			t.Errorf("ClimbStairs(%d) => %d,want %d", test.n, result, test.result)
		}
	}
}
