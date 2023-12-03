package main

import "fmt"

func main() {
	var a int = 5
	p := &a

	fmt.Println(a, p) //a=5 p=0xc0000b2008
}
