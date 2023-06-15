package q5

import (
	"strings"
)

//Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
//caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
//alfanuméricos incluem letras e números.
//
//Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	var carac = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", ":", " ", ","}
	for _, n := range carac {
		s = strings.ReplaceAll(s, n, "")
	}
	var strInvert = ""
	for i := len(s) - 1; i >= 0; i-- {
		strInvert += string(s[i])
	}
	if strInvert == s {
		return true
	}
	return false
}
