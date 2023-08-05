package main

import (
	"fmt"
)

/**
* 两个goroutine交替打印字母和数字
 */

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-ch1:
				fmt.Print(i)
				i++
				ch2 <- true
			}
		}
	}()

	go func() {
		j := 'A'
		for {
			select {
			case <-ch2:
				fmt.Printf("%c", j)
				j++
				if j > 'Z' {
					ch3 <- true
					return
				}
				ch1 <- true
			}
		}
	}()
	ch1 <- true
	<-ch3
}
