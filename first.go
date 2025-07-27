package main

import "fmt"

func main() {

	nums := []int{2, 11, 23, -2, 3, 5, 6, 0, 1}
	target := 9

	// --------------------------------------- My Code ------------------------------

	numbersMap := map[int]int{}
	resultSlice := []int{}

	for index, item := range nums {
		if _, found := numbersMap[target-item]; found {

			resultSlice = append(resultSlice, index)
			resultSlice = append(resultSlice, numbersMap[target-item])
		} else {
			numbersMap[item] = index
		}
	}
	fmt.Println(resultSlice)

}
