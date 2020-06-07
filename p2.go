package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
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
		listmap[m] = true
		n := k - m
		if listmap[n] {
			return true
		}
	}
	return false
}
