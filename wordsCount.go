package main

import (
	"fmt"
	"strings"
)

func main() {
	var value string
	fmt.Println("Type a Text:")
	fmt.Scan(value)
	wordCount, mapWord := WordCount(value)
	fmt.Printf("words count: %d", wordCount)
	fmt.Println(mapWord)
}

func WordCount(text string) (int, map[string]int) {
	wordsSlice := strings.Split(text, " ")
	wordCount := len(wordsSlice)
	mapWord := map[string]int{}

	for _, value := range wordsSlice {
		mapWord[value]++
	}

	return wordCount, mapWord
}
