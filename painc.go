package main

import (
	"time"
)

func main() {
	ch := make(chan int, 1)
	var i int
	go func() {
		for {
			select {
			case j := <-ch:
				print(j)
			default:
				i++
				ch <- i
				time.Sleep(1 * time.Second)
			}
		}
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
		}
	}()
	//debug.PrintStack()
	time.Sleep(1 * time.Second)
}
