package Combinefunc

import "fmt"

func Combine(str string)  {
	data := make(map[string]int)

	for _, ch := range str {
		
		if ch == rune(' ') {
			data["spaces"]++
		}
		if ch == rune(' ') || ch == rune('\n') {
			data["words"]++
		}
		if ch == rune(';') || ch == rune(':') || ch == rune('\'') || ch == rune('"') || ch == rune(',') || ch == rune('.') || ch == rune('?') {
			data["punctuation"]++
		}
		if ch == rune('[') || ch == rune(']') || ch == rune('{') || ch == rune('}') || ch == rune('(') || ch == rune(')') || ch == rune('|') || ch == rune('\\') || ch == rune('/') || ch == rune('+') || ch == rune('=') || ch == rune('<') || ch == rune('>') {
			data["symboles"]++
		}
		if ch == rune('!') || ch == rune('@') || ch == rune('#') || ch == rune('$') || ch == rune('%') || ch == rune('^') || ch == rune('&') || ch == rune('*') || ch == rune('~') {
			data["specialCharacters"]++
		}
		if ch == rune('0') || ch == rune('1') || ch == rune('2') || ch == rune('3') || ch == rune('4') || ch == rune('5') || ch == rune('6') || ch == rune('7') || ch == rune('8') || ch == rune('9') {
			data["digits"]++
		}
		if ch == rune('a') || ch == rune('e') || ch == rune('i') || ch == rune('o') || ch == rune('u') {
			data["vowels"]++
		}
		if ch == rune('\n') {
			data["paragraph"]++
		}
		if ch == rune('.') {
			data["lines"]++
		}
		if ch == rune('\n') {
			data["paragraph"]++
		} 
	}

	fmt.Println(data)
}