package main

import (
	"fmt"
)

type Set map[int]bool

func calculateWeights(s string) Set {
	weights := make(Set)
	i := 0
	for i < len(s) {
		j := i
		for j < len(s) && s[j] == s[i] {
			weight := int(s[i]-'a'+1) * (j - i + 1)
			weights[weight] = true
			j++
		}
		i = j
	}
	return weights
}

func weightedStrings(s string, queries []int) []string {
	weights := calculateWeights(s)
	results := make([]string, 0)

	for _, q := range queries {
		if weights[q] {
			results = append(results, "Yes")
		} else {
			results = append(results, "No")
		}
	}

	return results
}

func main() {
	inputString := "abbcccd"
	inputQueries := []int{1, 3, 9, 8}
	output := weightedStrings(inputString, inputQueries)
	fmt.Println(output)
}
