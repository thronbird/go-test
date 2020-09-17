package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//t := time.Now()
	//print(t.Month())
	i, err := strconv.ParseInt("1600236899", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm)

	//j:= 9898932
	//print(j/1000)

	k := 872
	print(k % 100)
}
