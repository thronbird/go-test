package main

import "sync"

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			sendRpc(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func sendRpc(i int) {
	print(i)
}
