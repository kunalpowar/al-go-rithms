package main

import (
	"flag"
	"fmt"
	"sync"
)

var n = flag.Int("n", 10, "-n=10 print all primes smaller than n")

func main() {
	flag.Parse()

	if *n < 2 {
		flag.Usage()
		return
	}

	if *n == 2 {
		fmt.Println(2)
		return
	}

	var (
		markCh = make(chan int)
		done   = make(chan bool)

		marked = make(map[int]bool)
	)

	wg := sync.WaitGroup{}

	for i := 2; i <= *n; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mark(num, markCh)
		}(i)
	}

	go func() {
		defer wg.Done()
		for {
			select {
			case m := <-markCh:
				marked[m] = true
			case <-done:
				break
			}
		}
	}()

	wg.Wait()

	for i := 2; i <= *n; i++ {
		if !marked[i] {
			fmt.Println(i)
		}
	}
}

func mark(num int, ch chan int) {
	i := 2
	multiple := num * i
	for {
		if multiple > *n {
			break
		}

		ch <- multiple
		i++
		multiple = num * i
	}
}
