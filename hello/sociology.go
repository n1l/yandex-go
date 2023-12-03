package main

import "fmt"

func main() {
	var yearOfBirth int

	fmt.Println("Введите год рождения: ")
	_, err := fmt.Scanf("%d", &yearOfBirth)

	if err != nil {
		fmt.Println("Unexpected error.")
		fmt.Println(err)
	}

	switch {
	case yearOfBirth >= 1946 && yearOfBirth <= 1964:
		fmt.Println("Привет, бумер!")
	case yearOfBirth >= 1965 && yearOfBirth <= 1980:
		fmt.Println("Привет, представитель X!")
	case yearOfBirth >= 1981 && yearOfBirth <= 1996:
		fmt.Println("Привет, миллениал!")
	case yearOfBirth >= 1997 && yearOfBirth <= 2012:
		fmt.Println("Привет, зумер!")
	case yearOfBirth >= 2013:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Здравствуйте!")
	}
}
