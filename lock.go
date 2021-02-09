package main

import "sync"

var l sync.Mutex

func main() {

	l.Lock()
	l.Unlock()
}
