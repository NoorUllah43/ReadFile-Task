package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	timeStart := time.Now()

	file, ferr := os.ReadFile("file.txt")
	if ferr != nil {
		panic(ferr)
	}

	str := string(file)

	// ch := make(chan map[string]int)
	// dividedFile := []int{5318003, 10636006, 15954009, 21272012}

	// for i := 0; i < len(dividedFile); i++ {
	// 	if i == 0 {
	// 		go Chunke(str, 0, dividedFile[i])
	// 	}else {
	// 		go Chunke(str, dividedFile[i-1], dividedFile[i])
	// 	}

	// }
	go Chunke(str,0,5318003 )
	go Chunke(str,5318003,10636006 )
	go Chunke(str,10636006,15954009 )
	go Chunke(str,15954009,21272012 )

	// received := <-ch

	// fmt.Println("Received data from channel:", received)

	time.Sleep(time.Second)

	takenTime := time.Since(timeStart)
	fmt.Printf("time: %v \n", takenTime)

}

func Chunke(str string, start int, length int) {

	data := make(map[string]int)

	for i := start; i < length; i++ {
		data["character"]++
		if str[i] == (' ') {
			data["spaces"]++
		}
		if str[i] == (' ') || str[i] == ('\n') {
			data["words"]++
		}
		if str[i] == (';') || str[i] == (':') || str[i] == ('\'') || str[i] == ('"') || str[i] == (',') || str[i] == ('.') || str[i] == ('?') {
			data["punctuation"]++
		}
		if str[i] == ('[') || str[i] == (']') || str[i] == ('{') || str[i] == ('}') || str[i] == ('(') || str[i] == (')') || str[i] == ('|') || str[i] == ('\\') || str[i] == ('/') || str[i] == ('+') || str[i] == ('=') || str[i] == ('<') || str[i] == ('>') {
			data["symboles"]++
		}
		if str[i] == ('!') || str[i] == ('@') || str[i] == ('#') || str[i] == ('$') || str[i] == ('%') || str[i] == ('^') || str[i] == ('&') || str[i] == ('*') || str[i] == ('~') {
			data["specialstr[i]aracters"]++
		}
		if str[i] == ('0') || str[i] == ('1') || str[i] == ('2') || str[i] == ('3') || str[i] == ('4') || str[i] == ('5') || str[i] == ('6') || str[i] == ('7') || str[i] == ('8') || str[i] == ('9') {
			data["digits"]++
		}
		if str[i] == ('a') || str[i] == ('e') || str[i] == ('i') || str[i] == ('o') || str[i] == ('u') {
			data["vowels"]++
		}
		if str[i] == ('.') {
			data["lines"]++
		}
		if str[i] == ('\n') {
			data["paragraph"]++
		}
	}

	// ch <- data
	fmt.Println(data)

}
