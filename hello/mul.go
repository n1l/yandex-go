package main

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch num := a.(type) {
	case int:
		return num * b
	case string:
		return strings.Repeat(num, b)
	case fmt.Stringer:
		return strings.Repeat(num.String(), b)
	default:
		return nil
	}
}

func main() {
	fmt.Println(Mul(5, 7))
	fmt.Println(Mul("5", 7))
}
