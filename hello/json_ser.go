package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	person := Person{
		Name:        "Aлекс",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(),
	}
	serializedPerson, error := json.Marshal(person)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(string(serializedPerson))
}
