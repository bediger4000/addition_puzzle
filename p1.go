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

// solve is an O(n^2) solution
func solve(k int, numberList []int) bool {
	for mIdx, m := range numberList {
		for nIdx, n := range numberList {
			if mIdx == nIdx {
				continue
			}
			if m+n == k {
				return true
			}
		}
	}

	return false
}
