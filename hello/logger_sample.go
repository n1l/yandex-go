package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer

	var logger = log.New(&buf, "", 0)
	logger.Print("Hello, world!")
	logger.Print("Goodbye")
	fmt.Print(&buf)
}
