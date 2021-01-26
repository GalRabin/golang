package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg2 sync.WaitGroup


func main() {
	wg2.Add(5)
	var x  int64 = 0
	for j := 0; j < 5; j++ {
		go func() {
			for i := 0; i < 4; i++ {
				fmt.Println(atomic.LoadInt64(&x))
				runtime.Gosched()
				atomic.AddInt64(&x, 1)
			}
			wg2.Done()
		}()
	}


	wg2.Wait()
	fmt.Println(x)
}