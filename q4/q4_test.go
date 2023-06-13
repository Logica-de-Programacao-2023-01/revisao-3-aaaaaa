package q4

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			nums:   []int{2, 2, 1},
			result: 2,
		},
		{
			nums:   []int{4, 1, 2, 1, 2},
			result: 0,
		},
		{
			nums:   []int{1},
			result: 0,
		},
		{
			nums:   []int{1, 1, 2},
			result: 2,
		},
		{
			nums:   []int{1, 1, 2, 2, 3},
			result: 4,
		},
		{
			nums:   []int{1, 1, 2, 2, 3, 3, 4},
			result: 6,
		},
	}

	for _, test := range tests {
		result := SingleNumber(test.nums)
		if result != test.result {
			t.Errorf("SingleNumber(%v) => %d,want %d", test.nums, result, test.result)
		}
	}
}
