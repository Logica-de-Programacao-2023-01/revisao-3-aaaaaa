package q1

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			result: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			result: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			result: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 3},
			target: 6,
			result: []int{0, 2},
		},
	}

	for _, test := range tests {
		result := TwoSum(test.nums, test.target)
		if len(result) != len(test.result) {
			t.Errorf("twoSum(%v, %d) => %v,want %v", test.nums, test.target, result, test.result)
		}
		for i, v := range result {
			if v != test.result[i] {
				t.Errorf("twoSum(%v, %d) => %v,want %v", test.nums, test.target, result, test.result)
			}
		}
	}
}
