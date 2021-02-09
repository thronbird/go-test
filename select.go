package main

import (
	"fmt"
)

func main() {
	//select {
	//case <- time.Tick(1* time.Second):
	//	fmt.Println("ssss") // "0" "2" "4" "6" "8"
	//}
	ch := make(chan int, 1)
	//for i := 0; i < 10; i++ {
	select {
	case x := <-ch:
		fmt.Println(x) // "0" "2" "4" "6" "8"
		//case ch <- 2:
		//default:
		//	print("x")
	}
	//}
}
