package utils

import (
	"strings"
)

func VerifyIfContainsInvalidCharacters(word string) bool{
	invalidcharacter := []string{"(", ")", "!",".",",",";",":"}

	var containsInvalid bool

	for i:=0; i < len(invalidcharacter) ; i++ {
		
		if(strings.Contains(word, invalidcharacter[i])){
			containsInvalid = true
			break;
		}
	}

	return containsInvalid
}