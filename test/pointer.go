package main

import (
	"fmt"
	"sync"
)

var lock sync.RWMutex

type Counter struct {
	T0    int64
	Count int64
	mux   sync.Mutex
}

func main() {
	m := map[string]*Counter{"aaa": {T0: 111, Count: 222}}
	change(m)
	fmt.Println(m["aaa"].T0)

	//for i:=1;i<=30;i++{
	//	go change(m)
	//}
	//fmt.Println(m["aaa"].T0)
}

func change(mp map[string]*Counter) {
	for _, cnt := range mp {
		println(&cnt.mux)
		cnt.mux.Lock()
		defer cnt.mux.Unlock()
		println(&cnt.mux)
		*cnt = Counter{
			T0:    cnt.T0 + 3,
			Count: 0,
		}
		println(&cnt.mux)
	}
	fmt.Println(mp["aaa"].T0)
}
