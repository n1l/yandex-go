package main

import "fmt"

var Global = 5

func UseGlobal() {
	defer func(checkout int) {
		Global = checkout
	}(Global)

	Global = 42

	fmt.Println(Global)
}

func main() {
	fmt.Println(Global)

	UseGlobal()

	fmt.Println(Global)

}
