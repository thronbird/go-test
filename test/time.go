package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//t := time.Now()
	//print(t.Month())
	i, err := strconv.ParseInt("1600229553", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm)
}
