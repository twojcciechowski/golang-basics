package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	f5()
}

func f1() {
	c := make(chan int) //can be also as buffered channel if make(chan int, 100)
	println("capacity", cap(c))

	go func(x chan int) {
		time.Sleep(10 * time.Second)
		x <- 1
	}(c)

	a := <-c
	println("message", a)
}

func f2() {
	println("Hello")
	var wg sync.WaitGroup

	for x := 0; x < 1000; x++ {
		wg.Add(1)

		go func(x int) {
			time.Sleep(1 * time.Second)
			println("after sleep: ", x)
			wg.Done()
		}(x)
	}

	println("Tom")

	wg.Wait()

	println("ending")
}

var UsersArray = []string{}
var mu sync.Mutex

func f3() {
	println("Hello")
	var wg sync.WaitGroup

	for x := 0; x < 1000; x++ {
		wg.Add(1)

		go processData(x, &wg)
	}

	println("Tom")

	wg.Wait()

	println("ending")

	println("UsersArray lenght: ", len(UsersArray))
	//for _, v := range UsersArray {
	//	println(v)
	//}
}

func processData(x int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	println("after sleep: ", x)
	defer func() {
		mu.Unlock()
	}()
	mu.Lock()
	UsersArray = append(UsersArray, fmt.Sprintf("user%v", x))

	wg.Done()
}

func f4() {

	count := 12

	c := make(chan int, 12)
	s := make(chan string, 12)
	defer func() {
		close(c)
		close(s)
	}()

	//goroutines
	for x := 0; x < 10; x++ {

		go func(id int, x chan int, y chan string) {

			//println("capacity: ", cap(x))
			//reading from chanel x
			for i := range x {

				//fmt.Printf("id: %v, i: %v \n", id, i)

				//convert i to string
				z := "abcd" + strconv.Itoa(i)

				//sending z to chanel y
				y <- z

			}
		}(x, c, s)
	}

	//sending messages
	for x := 0; x < count; x++ {
		c <- x
	}

	valS := []string{}
	for x := 0; x < count; x++ {
		v := <-s
		valS = append(valS, v)
	}

	println(len(valS))
}

func f5() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "key1", "val1")

	ctx, cancelFunc := context.WithTimeout(ctx, 1*time.Second)

	defer cancelFunc()
	out := make(chan string)
	runWithTimeout(ctx, out)

	select {

	case <-ctx.Done():
		{
			log.Printf("Processing to long \n")
			log.Print(ctx.Err())
		}
	case s := <-out:
		{
			log.Printf("Processing ended with result: %s \n", s)
		}
	}
}

func runWithTimeout(ctx context.Context, out chan string) {
	log.Printf("Starting key %v, val: %v \n", "key1", ctx.Value("key1"))

	go func() {
		log.Printf("Staring processing\n")

		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		sec := r1.Float64() * 7

		duration := time.Duration(sec) * time.Second
		log.Printf("Processing will take : %v , %v \n", sec, duration)
		time.Sleep(duration)
		log.Printf("Processing finished\n")
		out <- "Ready"
	}()

}
