package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	client := resty.New()

	client.
		// устанавливаем количество повторений
		SetRetryCount(3).
		// длительность ожидания между попытками
		SetRetryWaitTime(30 * time.Second).
		// длительность максимального ожидания
		SetRetryMaxWaitTime(90 * time.Second)

	var users []User
	url := "https://jsonplaceholder.typicode.com/users"

	_, err := client.R().SetResult(&users).Get(url)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range users {
		fmt.Printf("%s ", v.Username)
	}
}
