package main

import "fmt"

func main() {
	var value string

	fmt.Print("Type a String: ")
	fmt.Scan(&value)

	if isOk := Palindrome(value); isOk {
		println(true)
	} else {
		println(false)
	}
}

func Palindrome(str string) bool {
	stringLength := len(str)
	fIndex := 0
	lIndex := stringLength - 1

	midIndex := stringLength / 2
	for fIndex != lIndex {
		if fIndex > midIndex {
			return true
		}
		if str[fIndex] == str[lIndex] {
			fIndex++
			lIndex--
		} else {
			return false
		}
	}

	return true
}
