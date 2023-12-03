package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a, len(a), cap(a)) // [1 2 3 4] 4 4

	b := a[2:3] // b = [3]
	b = append(b, 7)
	fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	fmt.Println(b, len(b), cap(b)) // [3 7] 2 2
	b = append(b, 8, 9, 10)
	b[0] = 11
	fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	fmt.Println(b, len(b), cap(b)) // [11 7 8 9 10] 5 6
}
