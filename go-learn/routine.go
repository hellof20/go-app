package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello()
	}
	fmt.Println("main goroutine done!")
	wg.Wait()
}
