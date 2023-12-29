package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Printf("redirect: %s \n", req.URL)
			return nil
		},
	}
	_, err := client.Get("http://ya.ru")
	if err != nil {
		log.Fatal(err)
		return
	}
}
