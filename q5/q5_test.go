package q5

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		result bool
	}{
		{
			s:      " ",
			result: true,
		},
		{
			s:      "a",
			result: true,
		},
		{
			s:      "aa",
			result: true,
		},
		{
			s:      "A man, a plan, a canal: Panama",
			result: true,
		},
		{
			s:      "Arara",
			result: true,
		},
		{
			s:      "A sacada da Casa",
			result: true,
		},
		{
			s:      "race a car",
			result: false,
		},
		{
			s:      "0P",
			result: false,
		},
		{
			s:      "a menina malina",
			result: false,
		},
		{
			s:      "andolaeta o ale adno",
			result: false,
		},
	}

	for _, test := range tests {
		result := IsPalindrome(test.s)
		if result != test.result {
			t.Errorf("IsPalindrome(%v) => %t,want %t", test.s, result, test.result)
		}
	}
}
