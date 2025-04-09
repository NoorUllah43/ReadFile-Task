package main

import (
	"fmt"
	"os"
	"time"
	"sync"
)

func main() {
	timeStart := time.Now()

	file, ferr := os.ReadFile("file.txt")
	if ferr != nil {
		panic(ferr)
	}

	str := string(file)

	chunkelen := [4]int{}
	for i := 0; i < 4; i++ {
		if i == 0 {
			chunkelen[i] = len(str) / 4
		} else {
			chunkelen[i]= chunkelen[i-1]+ chunkelen[0]

		}
	}

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func()  {
			defer wg.Done()
			if i==0 {
				Chunke(str,0,chunkelen[i])
			}else {
				go Chunke(str, chunkelen[i-1], chunkelen[i])
				
			}
			
		}()
	}
	
	// go Chunke(str, 0, chunkelen[0])
	// go Chunke(str, chunkelen[0], chunkelen[1])
	// go Chunke(str, chunkelen[1], chunkelen[2])
	// go Chunke(str, chunkelen[2], chunkelen[3])

	// received := <-ch

	// fmt.Println("Received data from channel:", received)

	time.Sleep(time.Second)

	takenTime := time.Since(timeStart)
	fmt.Printf("time: %v \n", takenTime)

}

func Chunke(str string, start int, length int) {

	chunkeData := make(map[string]int)

	for i := start; i < length; i++ {
		chunkeData["character"]++
		if str[i] == (' ') {
			chunkeData["spaces"]++
		}
		if str[i] == (' ') || str[i] == ('\n') {
			chunkeData["words"]++
		}
		if str[i] == (';') || str[i] == (':') || str[i] == ('\'') || str[i] == ('"') || str[i] == (',') || str[i] == ('.') || str[i] == ('?') {
			chunkeData["punctuation"]++
		}
		if str[i] == ('[') || str[i] == (']') || str[i] == ('{') || str[i] == ('}') || str[i] == ('(') || str[i] == (')') || str[i] == ('|') || str[i] == ('\\') || str[i] == ('/') || str[i] == ('+') || str[i] == ('=') || str[i] == ('<') || str[i] == ('>') {
			chunkeData["symboles"]++
		}
		if str[i] == ('!') || str[i] == ('@') || str[i] == ('#') || str[i] == ('$') || str[i] == ('%') || str[i] == ('^') || str[i] == ('&') || str[i] == ('*') || str[i] == ('~') {
			chunkeData["specialstr[i]aracters"]++
		}
		if str[i] == ('0') || str[i] == ('1') || str[i] == ('2') || str[i] == ('3') || str[i] == ('4') || str[i] == ('5') || str[i] == ('6') || str[i] == ('7') || str[i] == ('8') || str[i] == ('9') {
			chunkeData["digits"]++
		}
		if str[i] == ('a') || str[i] == ('e') || str[i] == ('i') || str[i] == ('o') || str[i] == ('u') {
			chunkeData["vowels"]++
		}
		if str[i] == ('.') {
			chunkeData["lines"]++
		}
		if str[i] == ('\n') {
			chunkeData["paragraph"]++
		}
	}

	// ch <- chunkeData
	fmt.Println(chunkeData)
	time.Sleep(time.Second)


}
