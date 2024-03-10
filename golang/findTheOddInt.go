package main

import "strconv"

func FindTheOddInt(nums []string) (oddCount int) {
	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err == nil && n%2 != 0 {
			oddCount++
		}
	}
	return oddCount
}
