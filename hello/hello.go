package main

import "fmt"

func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // круглые скобки в конце означают, что функция вызывается
	return "Казалось бы"
}

func intuitive() string {
	value := "Казалось бы"
	defer func() { value = "На самом деле" }()
	return value
}

func main() {
	value := unintuitive()

	fmt.Println(value)

	value = intuitive()

	fmt.Println(value)

	var str string
	str = "Hello, world!"

	fmt.Println(string(str[0]))
	fmt.Println("Hello, world!")
}
