package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Set up the array of numbers from command line
	k, _ := strconv.Atoi(os.Args[1])
	var numberList []int
	for _, str := range os.Args[2:] {
		n, err := strconv.Atoi(str)
		if err == nil {
			numberList = append(numberList, n)
		}
	}

	cannit := solve(k, numberList)
	if cannit {
		fmt.Printf("two numbers from list add up to %d\n", k)
	}
}

// solve makes one pass over numberList and finds
// 2 numbers that add up to k.
func solve(k int, numberList []int) bool {
	listmap := make(map[int]bool)
	for _, m := range numberList {
		n := k - m
		if listmap[n] {
			return true
		}
		// putting current number m into the map *after*
		// checking for addition solution so that if m+m == k,
		// this func doesn't give the wrong answer.
		listmap[m] = true
	}
	return false
}
