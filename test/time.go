package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "01/02/2006 3:04:05 PM"
	t, err := time.Parse(layout, "02/28/2016 9:03:46 PM")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix())
}
