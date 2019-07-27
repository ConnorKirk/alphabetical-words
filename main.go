package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const src = "https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt"

func main() {
	file, err := os.Open("words_alpha.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var alphabetical []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if isAlphabetical(word) {
			alphabetical = append(alphabetical, word)
		}
	}

	fmt.Printf("There are %v alphabetical words\n", len(alphabetical))
	var b bytes.Buffer
	for _, word := range alphabetical {
		b.WriteString(word + "\n")
	}

	err = ioutil.WriteFile("alphabetical_words.txt", b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func isAlphabetical(word string) bool {
	for i := range word[:len(word)-1] {
		if word[i] > word[i+1] {
			return false
		}
	}
	return true
}
