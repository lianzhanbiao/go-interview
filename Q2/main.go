package main

import (
	"fmt"
	"time"
)

// 写一个简单的消费者生产者模型

func Producer(out chan<- int, value int) {
	for i := 0; ; i++ {
		out <- value
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Print(v, " ")
	}
}

func main() {
	messages := make(chan int, 64)
	go Producer(messages, 5)
	go Producer(messages, 3)
	go Consumer(messages)

	time.Sleep(2 * time.Second)
}
