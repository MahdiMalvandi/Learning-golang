package main

import "fmt"

func main() {
	var value string

	fmt.Print("Type a String: ")
	fmt.Scan(&value)

	fmt.Println(ReverseString(value))
}


func ReverseString(str string) string {
	result := []rune(str)
	stringLength := len(str)
	fIndex := 0
	lIndex := stringLength - 1

	midIndex := stringLength / 2
	for fIndex != lIndex {
		if fIndex > midIndex {
			break
		}
		runes := []rune(str)
		result[fIndex] = runes[lIndex]
		result[lIndex] = runes[fIndex]
		fIndex++
		lIndex--
	}

	return string(result)
}
