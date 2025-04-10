package main

import (
	"fmt"
	"os"
	"time"
	"readFile-task/Combinefunc"
	// "readFile-task/Separatefuncs"
	"sync"
	
)

func main() {
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
			chunkelen[i] = chunkelen[i-1] + chunkelen[0]

		}
	}

	var wg sync.WaitGroup

	wg.Add(4)
	ch1 := make(chan map[string]int,1)
	ch2 := make(chan map[string]int,1)
	ch3 := make(chan map[string]int,1)
	ch4 := make(chan map[string]int,1)

	

	go Chunke(str, 0, chunkelen[0], &wg, ch1)
	go Chunke(str, chunkelen[0], chunkelen[1], &wg, ch2)
	go Chunke(str, chunkelen[1], chunkelen[2], &wg, ch3)
	go Chunke(str, chunkelen[2], chunkelen[3], &wg, ch4)
	

	wg.Wait()
	data := make(map[string]int)


	receivedData1 := <-ch1
	receivedData2 := <-ch2
	receivedData3 := <-ch3
	receivedData4 := <-ch4


	for key,value := range receivedData1 {
		data[key] = value + receivedData2[key] + receivedData3[key] + receivedData4[key]
	}

	fmt.Println(data)
	
	// Separatefuncs.NumCharacter(str)
	// Separatefuncs.NumLines(str)
	// Separatefuncs.NumDigits(str)
	// Separatefuncs.NumParagraph(str)
	// Separatefuncs.NumPunctuation(str)
	// Separatefuncs.NumSpaces(str)
	// Separatefuncs.NumSpecialCharacter(str)
	// Separatefuncs.NumSymboles(str)
	// Separatefuncs.NumVowels(str)
	// Separatefuncs.NumWord(str)


	timeStart := time.Now()

	Combinefunc.Combine(str)
	

	time.Sleep(time.Second)

	takenTime := time.Since(timeStart)
	fmt.Printf("time: %v \n", takenTime)

}


func Chunke(str string, start int, length int, wg *sync.WaitGroup, ch chan map[string]int) {
	defer wg.Done()

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
			chunkeData["specialCharacters"]++
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

	ch <- chunkeData
	// fmt.Println(chunkeData)

}