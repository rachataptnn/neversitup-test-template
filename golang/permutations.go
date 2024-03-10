package main

import "sort"

func Permutations(input string) []string {
	result := []string{}
	permute([]rune(input), 0, &result)
	sort.Strings(result) // Sorting to remove duplicates
	return result
}

func permute(s []rune, index int, result *[]string) {
	if index == len(s)-1 {
		*result = append(*result, string(s))
		return
	}
	seen := make(map[rune]bool)
	for i := index; i < len(s); i++ {
		if seen[s[i]] {
			continue
		}
		seen[s[i]] = true
		s[index], s[i] = s[i], s[index]
		permute(s, index+1, result)
		s[index], s[i] = s[i], s[index]
	}
}
