package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var t = string2Time("")
	println(t.IsZero())

	println(time.Now().Unix())
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

	//print(time.Now().Unix()/1000)
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	fmt.Println(timestamp) // prints: 1436773875771421417

	fmt.Println(time.Now().Unix())
}

func string2Time(t string) time.Time {
	timeTemplate1 := "2006-01-02T15:04:05Z"
	stamp, _ := time.ParseInLocation(timeTemplate1, t, time.UTC)
	return stamp
}
