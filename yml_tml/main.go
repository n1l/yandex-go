package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

type Data struct {
	ID     int    `toml:"id"`
	Name   string `toml:"name"`
	Values []byte `toml:"values"`
}

const yamlData = `
id: 101
name: Gopher
values:
- 11
- 22
- 33
`

func main() {
	var data Data
	err := yaml.Unmarshal([]byte(yamlData), &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	dataBytes, err := toml.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(dataBytes))
}
