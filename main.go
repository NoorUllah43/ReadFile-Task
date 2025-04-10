package main

import (
	"fmt"
	"os"
	"readFile-task/Combinefunc"
	"time"
	// "readFile-task/Separatefuncs"
	"sync"
)

func main() {

	// reading the file
	file, ferr := os.ReadFile("../file.txt")
	if ferr != nil {
		panic(ferr)
	}

	// convert the buffer into string
	str := string(file)

	// make the chunking size and store it in chunklen
	// chunks := [4]int{}
	totalChunks := 4
	chunks := make([]int, totalChunks)

	for i := 0; i < totalChunks; i++ {
		if i == 0 {
			chunks[0] = len(str) / totalChunks
		} else {
			chunks[i] = chunks[i-1] + chunks[0]

		}
	}

	fmt.Printf("\n\n ------------- using gorutines ----------------\n")

	var wg sync.WaitGroup

	
	ch1 := make(chan map[string]int, 1)
	ch2 := make(chan map[string]int, 1)
	ch3 := make(chan map[string]int, 1)
	ch4 := make(chan map[string]int, 1)

	rutineTime := time.Now()
	wg.Add(4)

	go Chunke(str, 0, chunks[0], &wg, ch1)
	go Chunke(str, chunks[0], chunks[1], &wg, ch2)
	go Chunke(str, chunks[1], chunks[2], &wg, ch3)
	go Chunke(str, chunks[2], chunks[3], &wg, ch4)

	time.Sleep(time.Second)

	takenTimeRutine := time.Since(rutineTime)

	wg.Wait()

	data := make(map[string]int)

	receivedData1 := <-ch1
	receivedData2 := <-ch2
	receivedData3 := <-ch3
	receivedData4 := <-ch4

	// add all receiving data and store it in data map
	for key, value := range receivedData1 {
		data[key] = value + receivedData2[key] + receivedData3[key] + receivedData4[key]
	}

	fmt.Println(data)
	fmt.Printf("execution time of gorutines: %v \n", takenTimeRutine)

	fmt.Printf("\n\n ------------- using combine function ----------------\n")

	timeStart := time.Now()

	Combinefunc.Combine(str)

	time.Sleep(time.Second)

	takenTime := time.Since(timeStart)
	fmt.Printf("execution time of combine fucntion: %v \n", takenTime)

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

}

func Chunke(str string, start int, length int, wg *sync.WaitGroup, ch chan map[string]int) {
	defer wg.Done()

	chunkeData := make(map[string]int)

	for i := start; i < length; i++ {
		// chunkeData["character"]++
		if str[i] == (' ') {
			chunkeData["spaces"]++
		}
		if str[i] == ('.') {
			chunkeData["lines"]++
		}
		if str[i] == ('\n') {
			chunkeData["paragraph"]++
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
	}

	ch <- chunkeData
	// fmt.Println(chunkeData)

}
