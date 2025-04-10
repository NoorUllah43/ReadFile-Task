package main

import (
	"fmt"
	"os"
	"time"
	"readFile-task/Combinefunc"
	"readFile-task/Separatefuncs"
	
)

func main() {
	file, ferr := os.ReadFile("file.txt")
	if ferr != nil {
		panic(ferr)
	}

	str := string(file)

	
	Separatefuncs.NumCharacter(str)
	Separatefuncs.NumLines(str)
	Separatefuncs.NumDigits(str)
	Separatefuncs.NumParagraph(str)
	Separatefuncs.NumPunctuation(str)
	Separatefuncs.NumSpaces(str)
	Separatefuncs.NumSpecialCharacter(str)
	Separatefuncs.NumSymboles(str)
	Separatefuncs.NumVowels(str)
	Separatefuncs.NumWord(str)


	timeStart := time.Now()

	Combinefunc.Combine(str)
	

	time.Sleep(time.Second)

	takenTime := time.Since(timeStart)
	fmt.Printf("time: %v \n", takenTime)

}



// Number of lines----------------------------------------------------------------------------
