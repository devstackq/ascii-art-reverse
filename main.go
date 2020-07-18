package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Ascii struct {
	letter []string
}

func main() {
	//ascii create letters
	alpahabet := asciiLibs()
	hms := make(map[string]string)
	alphaNum := 0

	for _, letter := range alpahabet {
		hms[letter] = string(rune(32 + alphaNum))
		alphaNum++
	}
	var arrBytesFile []byte

	if os.Args[1][0:10] == "--reverse=" {
		file, _ := os.Open(os.Args[1][10:])
		arrBytesFile, _ = ioutil.ReadAll(file)
	}

	splittedWord := strings.Split(string(arrBytesFile), "\n")
	argStructWord := &Ascii{}
	argSecondStructWord := &Ascii{}

	var firstWord []string
	var secondWord []string

	wordNewLine := false
	firstLen := 0
	secondLen := 0
	if len(splittedWord) > 9 {
		wordNewLine = true
		firstWord = splittedWord[0:8]
		secondWord = splittedWord[8:16]
		firstLen = len(firstWord[0])
		secondLen = len(secondWord[0])
	}

	lenArg := len(splittedWord[0])

	var word, word1, word2 string
	if wordNewLine {
		//idx 9 - 17, return this letter, and concat first
		word1 = compareASCII(findLetterByIndex(firstLen, firstWord, argStructWord), hms)
		word1 += "\n"
		word2 = compareASCII(findLetterByIndex(secondLen, secondWord, argSecondStructWord), hms)
		word = word1 + word2
	} else {
		word = compareASCII(findLetterByIndex(lenArg, splittedWord, argStructWord), hms)
	}
	fmt.Print(word)
}

func findLetterByIndex(lenArg int, splittedArg []string, argStructWord *Ascii) *Ascii {

	count := 0
	notFind := false
	start := 0
	countSpace := 0

	for column := 0; column < lenArg; column++ {

		for row := 0; row < 8; row++ {
			if splittedArg[row][column] == 32 {
				count++
			} else {
				notFind = true
				break
			}
			if count == 8 {
				find := createLetterByIndex(start, column, splittedArg)
				for _, v := range find.letter {

					if v == " \n \n \n \n \n \n \n \n" {
						countSpace++
					}

					argStructWord.letter = append(argStructWord.letter, v)
				}
				start = column + 1
				count = 0
				if countSpace == 6 {
					argStructWord.letter = append(argStructWord.letter, "      \n      \n      \n      \n      \n      \n      \n      \n")
					countSpace = 0
				}
				break
				//if line, count 6 8 newline, count 6, index + 6, found = append(found, space)
			}
		}
		if notFind {
			//next column
			count = 0
		}
	}
	return argStructWord
}

func createLetterByIndex(start, end int, inputSpl []string) *Ascii {

	ascii := Ascii{}
	res := ""
	for line := 0; line < 8; line++ {
		res += inputSpl[line][start : end+1]
		res += "\n"
	}
	ascii.letter = append(ascii.letter, res)
	return &ascii
}

func compareASCII(word *Ascii, hms map[string]string) string {

	resik := ""
	for _, arg := range word.letter {
		for k, v := range hms {
			if arg == k {
				resik += v
				break
			}
		}
	}
	return resik
}

func asciiLibs() []string {

	libsLetters, _ := os.Open("standard.txt")
	libs, _ := ioutil.ReadAll(libsLetters)
	splitLiba := strings.Split(string(libs), "\n")
	acc := ""
	var arrASCIILetters []string
	temp := 0
	for k := 0; k < len(splitLiba)-1; k += 9 {
		for t := 0; t < 8; t++ {
			acc += splitLiba[temp]
			temp++
			if t > 0 {
				acc += "\n"
			}
		}
		//last line
		acc += splitLiba[temp] + "\n"
		arrASCIILetters = append(arrASCIILetters, acc)
		temp++
		acc = ""
	}
	return arrASCIILetters
}
