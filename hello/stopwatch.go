package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	value  time.Time
	splits []time.Duration
}

func (sw *Stopwatch) Start() {
	sw.value = time.Now()
}

func (sw *Stopwatch) SaveSplit() {
	sw.splits = append(sw.splits, time.Now().Sub(sw.value))
}

func (sw Stopwatch) GetResults() []time.Duration {
	return sw.splits
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())

}
