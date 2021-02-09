package main

import (
	"fmt"
	"time"
)

/**
https://medium.com/golangspec/blocks-in-go-2f68768868f6
https://medium.com/golangspec/scopes-in-go-a6042bb4298c
https://medium.com/golangspec/labels-in-go-4ffd81932339

Besides blocks which are explicitly marked there a few implicit ones:
universe block contains all source code,
package block contains all package’s source code (package can be spread over several files in a single directory),
file block contains file’s source code,
for statement is in its own implicit block:
*/
func main() {
	{ // start outer block
		a := 1
		fmt.Println(a)
		{ // start inner block
			b := 2
			fmt.Println(b)
		} // end inner block
	} // end outer block
	tick := time.Tick(100 * time.Millisecond)

LOOP:
	for {
		select {
		case <-tick:
			i := 0
			fmt.Println("tick", i)
			break LOOP
		default:
			// "i" is undefined here
			fmt.Println("sleep")
			time.Sleep(30 * time.Millisecond)
		}
	}
	// "i" is undefined here

	goto X
X:
	{
		/**
		Labels are not block scoped (more about blocks and scoping in “Blocks in Go” and “Scopes in Go”) so it’s impossible to redeclare label inside nested block:
		*/
		//X:
	}
	{
		v := 1
		{
			fmt.Println(v)
		}
		fmt.Println(v)
	}
	// “undefined: v” compilation error
	// fmt.Println(v)

	v := "outer"
	fmt.Println(v)
	{
		v := "inner"
		fmt.Println(v)
		{
			fmt.Println(v)
		}
	}
	{
		fmt.Println(v)
	}
	fmt.Println(v)
}
