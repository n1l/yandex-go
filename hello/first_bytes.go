package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	bodyArr := make([]byte, 512)
	io.ReadAtLeast(response.Body, bodyArr, 512)
	fmt.Println(len(bodyArr))
	fmt.Println(string(bodyArr))
}
