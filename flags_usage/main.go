package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.1"

var Usage = func() {
	fmt.Fprintf(flag.CommandLine.Output(), "Version: %v\nUsage of %s:\n", version, os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	// допишите код
	flag.Parse()
}
