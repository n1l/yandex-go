package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type generator struct {
	rnd rand.Source // Генератор случайных чисел. Вообще rand.Rand уже реализует интерфейс io.Reader, но для примера мы реализуем его самостоятельно.
}

// New — обратите внимание, что мы возвращаем generator, присвоенный интерфейсу io.Reader, сама структура generator неэкспортируемая.
// Мы скрыли внутри пакета все детали.
func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

// Read — реализация io.Reader
func (g *generator) Read(bytes []byte) (n int, err error) { // error это тип ошибки, подробнее мы рассмотрим его в следующем разделе.
	for i := 0; i+8 < len(bytes); i += 8 {
		binary.LittleEndian.PutUint64(bytes[i:i+8], uint64(g.rnd.Int63()))
	}
	return len(bytes), nil
}

func main() {

	// создаём генератор случайных чисел
	generator := New(time.Now().UnixNano()) // в качестве затравки передаём ему текущее время, и при каждом запуске оно будет разным.

	buf := make([]byte, 16)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf) // единственный доступный метод, но он нам и нужен.
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}

}
