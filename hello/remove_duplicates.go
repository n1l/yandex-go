package main

import "fmt"

func removeDuplicates(input []string) []string {
	hash := make(map[string]bool, len(input))
	result := make([]string, 0)
	for _, val := range input {
		if _, ok := hash[val]; !ok {
			result = append(result, val)
		}
		hash[val] = true
	}

	return result
}

func main() {

	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	input = removeDuplicates(input)

	fmt.Print(input)
}
