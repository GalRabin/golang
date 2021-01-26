package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg1 sync.WaitGroup

func incrementer() func(increment bool) int {
	x := 0
	return func(increment bool) int {
		if increment {
			x += 1
		}

		return x
	}
}

func main() {
	wg1.Add(2)
	inc := incrementer()
	lock := sync.Mutex{}
	go func() {
		for i := 0; i < 10; i++ {
			lock.Lock()
			fmt.Println(inc(false))
			runtime.Gosched()
			inc(true)
			lock.Unlock()
		}
		wg1.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			lock.Lock()
			fmt.Println(inc(false))
			runtime.Gosched()
			inc(true)
			lock.Unlock()
		}
		wg1.Done()
	}()

	wg1.Wait()

	fmt.Println(inc(false))
}