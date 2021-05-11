package main

import (
	"fmt"
	"reflect"
)

//https://blog.golang.org/laws-of-reflection#TOC_7.
func main() {
	var x float64 = 3.4
	var y = "ssssss"

	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
	fmt.Println("value:", reflect.ValueOf(y).String())
	println()

	{
		v := reflect.ValueOf(x)
		fmt.Printf("value is %7.1e\n", v.Interface()) // Interface() vs ValueOf

		fmt.Println("type:", v.Type())
		fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
		fmt.Println("value:", v.Float())
	}
	println()

	{
		v := reflect.TypeOf(x)
		fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	}

	println()
	{
		//to keep the API simple, the "getter" and "setter" methods of Value operate on the largest type that can hold the value: int64 for all the signed integers, for instance.
		var x uint8 = 'x'
		v := reflect.ValueOf(x)
		fmt.Println("type:", v.Type())                            // uint8.
		fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
		x = uint8(v.Uint())                                       // v.Uint returns a uint64.
	}

	{
		// Kind of a reflection object describes the underlying type, not the static type. In other words, the Kind cannot discriminate an int from a MyInt even though the Type can.
		type MyInt int
		var x MyInt = 7
		v := reflect.ValueOf(x)
		println(v.Kind() == reflect.Int)

	}

	{
		var x float64 = 3.4
		v := reflect.ValueOf(x)
		fmt.Println("settability of v:", v.CanSet())
		//v.SetFloat(7.1) // Error: will panic . we pass a copy of x to reflect.ValueOf, so the interface value created as the argument to reflect.ValueOf is a copy of x, not x itself. Thus, if the statement were allowed to succeed, it would not update x, even though v looks like it was created from x.

		p := reflect.ValueOf(&x) // Note: take the address of x.
		fmt.Println("type of p:", p.Type())
		fmt.Println("settability of p:", p.CanSet())

		v2 := p.Elem()
		fmt.Println("settability of v:", v2.CanSet())
		v2.SetFloat(7.1)
		fmt.Println(v2.Interface())
		fmt.Println(x)
	}

	{
		type T struct {
			//the field names of T are upper case (exported) because only exported fields of a struct are settable.
			A int
			B string
		}
		t := T{23, "skidoo"}
		s := reflect.ValueOf(&t).Elem()
		typeOfT := s.Type()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d: %s %s = %v\n", i,
				typeOfT.Field(i).Name, f.Type(), f.Interface())
		}
		s.Field(0).SetInt(77)
		s.Field(1).SetString("Sunset Strip")
		fmt.Println("t is now", t)
	}
}
