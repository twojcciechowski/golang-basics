package main

import (
	"sync"
	"time"
)

func main() {
	c := make(chan int) //can be also as buffered channel if make(chan int, 100)
	println(cap(c))

	go func(x chan int) {
		x <- 1
	}(c)

	a := <-c
	println(a)

	println("Hello")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		println("after sleep")
		wg.Done()
	}()
	println("Tom")

	wg.Wait()
}

