package main

import (
	"fmt"
	"os"
)

func main() {
	file, ferr := os.ReadFile("file.txt")
	if ferr != nil {
		panic(ferr)
	}
	
	str := string(file)

	numCharacter(str)
	numWord(str)
	numLines(str)
	numSpaces(str)
	numVowels(str)
	numDigits(str)
	numSpecialCharacter(str)
	numSymboles(str)
	numPunctuation(str)
	

}





// Number of character
func numCharacter(str string)  {
	character := 0
	for i := 0; i < len(str); i++ {
		character++
	}
	fmt.Printf("Number of character: %v\n",character)
	
}

// Number of lines
func numLines(str string)  {
	Lines := 0
	for _,ch := range str {
		if ch == rune('\n') {
			Lines++
		}
		
	}
	fmt.Printf("Number of Lines: %v\n",Lines)
	
}

// Number of vowels
func numVowels(str string)  {
	vowels := 0
	for _,ch := range str {
		if ch == rune('a') || ch == rune('e') || ch == rune('i') || ch == rune('o') || ch == rune('u') {
			vowels++
		}
		
	}
	fmt.Printf("Number of vowels: %v\n",vowels)
	
}

// Number of digits
func numDigits(str string)  {
	digits := 0
	for _,ch := range str {
		if ch == rune('0') || ch == rune('1') || ch == rune('2') || ch == rune('3') || ch == rune('4') || ch == rune('5') || ch == rune('6') || ch == rune('7') || ch == rune('8') || ch == rune('9') {
			digits++
		}
		
	}
	fmt.Printf("Number of digits: %v\n",digits)
	
}

// Number of SpecialCharacter
func numSpecialCharacter(str string)  {
	SpecialCharacter := 0
	for _,ch := range str {
		if ch == rune('!') || ch == rune('@') || ch == rune('#') || ch == rune('$') || ch == rune('%') || ch == rune('^') || ch == rune('&') || ch == rune('*') || ch == rune('~') {
			SpecialCharacter++
		}
		
	}
	fmt.Printf("Number of SpecialCharacter: %v\n",SpecialCharacter)
	
}


// Number of symbol
func numSymboles(str string)  {
	symboles := 0
	for _,ch := range str {
		if ch == rune('[') || ch == rune(']') || ch == rune('{') || ch == rune('}') || ch == rune('(') || ch == rune(')') || ch == rune('|') || ch == rune('\\') || ch == rune('/') || ch == rune('+') || ch == rune('=') || ch == rune('<') || ch == rune('>'){
			symboles++
		}
		
	}
	fmt.Printf("Number of SpecialCharacter: %v\n",symboles)
	
}

// Number of symbol
func numPunctuation(str string)  {
	punctuation := 0
	for _,ch := range str {
		if ch == rune(';') || ch == rune(':') || ch == rune('\'') || ch == rune('"') || ch == rune(',') || ch == rune('.') || ch == rune('?') {
			punctuation++
		}
		
	}
	fmt.Printf("Number of SpecialCharacter: %v\n",punctuation)
	
}


// Number of word
func numWord(str string)  {
	words := 0
	for _,ch := range str {
		if ch == rune(' ') || ch == rune('\n') {
			words++
		}
		
	}
	fmt.Printf("Number of words: %v\n",words)
	
}

// Number of spaces
func numSpaces(str string)  {
	spaces := 0
	for _,ch := range str {
		if ch == rune(' ') {
			spaces++
		}
		
	}
	fmt.Printf("Number of spaces: %v\n",spaces)
	
}