package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

func (address *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", address.Host, address.Port)
}

func (address *NetAddress) Set(flagValue string) error {
	values := strings.Split(flagValue, ":")
	if len(values) != 2 {
		panic("address should be in '{host}:{port} format.'")
	}
	port, err := strconv.Atoi(values[1])
	if err != nil {
		panic(err)
	}
	address.Host = values[0]
	address.Port = port
	return nil
}

func main() {
	addr := new(NetAddress)
	// если интерфейс не реализован,
	// здесь будет ошибка компиляции
	_ = flag.Value(addr)
	// проверка реализации
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}
