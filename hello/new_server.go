package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func WriteHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "1")
	fmt.Fprint(w, "2")
	w.Write([]byte("3"))
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	// собираем данные
	subj := Subj{"Milk", 50}
	// кодируем в JSON
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "application/json")
	// устанавливаем код 200
	w.WriteHeader(http.StatusOK)
	// пишем тело ответа
	w.Write(resp)
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ===============\r\n"
	if err := req.ParseForm(); err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	res.Write([]byte(body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, WriteHandle)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
