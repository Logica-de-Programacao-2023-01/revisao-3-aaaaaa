package q2

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs   []string
		result string
	}{
		{
			strs:   []string{"flower", "flow", "flight"},
			result: "fl",
		},
		{
			strs:   []string{"dog", "racecar", "car"},
			result: "",
		},
		{
			strs:   []string{"", "racecar", "car"},
			result: "",
		},
		{
			strs:   []string{"", "racecar", ""},
			result: "",
		},
		{
			strs:   []string{"", "", ""},
			result: "",
		},
		{
			strs:   []string{"froolic", "froolic", "fraodf", "frloadsf", "frlozcv"},
			result: "fr",
		},
		{
			strs:   []string{"a"},
			result: "a",
		},
		{
			strs:   []string{"a", "a"},
			result: "a",
		},
		{
			strs:   []string{"apple", "apply", "apollo", "abracadabra"},
			result: "a",
		},
		{
			strs:   []string{"qwerty", "hello"},
			result: "",
		},
		{
			strs:   []string{"qwerty", "qwerty"},
			result: "qwerty",
		},
		{
			strs:   []string{"helloworld", "hell", "hello"},
			result: "hell",
		},
	}

	for _, test := range tests {
		result := LongestCommonPrefix(test.strs)
		if result != test.result {
			t.Errorf("longestCommonPrefix(%v) => %v,want %v", test.strs, result, test.result)
		}
	}
}
