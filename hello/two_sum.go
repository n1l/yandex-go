package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, val := range nums {
		if j, ok := hash[target-val]; ok {
			return []int{i, j}
		}

		hash[val] = i
	}

	return []int{}
}

func main() {

	fmt.Println("jj")
}
