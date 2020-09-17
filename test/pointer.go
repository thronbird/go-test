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
	CounterMap := map[string]*Counter{"aaa:": {T0: 111, Count: 222}}
	change(CounterMap)
	fmt.Println(CounterMap)
}

func change(mp map[string]*Counter) {
	for _, cnt := range mp {
		//cnt.mux.Lock()
		//defer cnt.mux.Unlock()
		*cnt = Counter{
			T0:    cnt.T0 + 3,
			Count: 0,
		}
	}
	fmt.Println(mp)
}
