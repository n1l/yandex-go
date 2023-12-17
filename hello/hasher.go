package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func main() {
	str := "https://practicum.yandex.ru/"

	sum := md5.Sum([]byte(str))
	encoded := base64.StdEncoding.EncodeToString(sum[:])

	fmt.Printf(encoded[:8])
}
