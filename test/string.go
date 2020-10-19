package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := int64(1600657834)

	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)
}
