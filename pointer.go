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
	//change2(m)
	//fmt.Println(m["aaa"].T0)
	var wg sync.WaitGroup

	for i := 1; i <= 30; i++ {
		wg.Add(1)
		go change(m, &wg)
	}
	wg.Wait()
	fmt.Println(m["aaa"].T0)
}

func change(mp map[string]*Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, cnt := range mp {
		cnt.mux.Lock()
		defer cnt.mux.Unlock()

		*cnt = Counter{
			T0:    cnt.T0 + 3,
			Count: 0,
		}
	}
	fmt.Println(mp["aaa"].T0)
}

func change2(mp map[string]*Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, cnt := range mp {
		cnt.mux.Lock()
		defer cnt.mux.Unlock()

		cnt.T0 = cnt.T0 + 3
		cnt.Count = 0
	}
	//fmt.Println(mp["aaa"].T0)
}
