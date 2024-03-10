package main

import "regexp"

func CountSmileys(arr []string) int {
	validSmileyRegex := regexp.MustCompile(`[:;][-~]?[)D]`)
	count := 0
	for _, s := range arr {
		if validSmileyRegex.MatchString(s) {
			count++
		}
	}
	return count
}
