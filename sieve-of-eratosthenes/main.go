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

	marked := new(sync.Map)
	wg := sync.WaitGroup{}

	for i := 2; i <= *n; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mark(num, marked)
		}(i)
	}

	wg.Wait()

	for i := 2; i <= *n; i++ {
		if _, present := marked.Load(i); !present {
			fmt.Println(i)
		}
	}
}

func mark(num int, markMap *sync.Map) {
	i := 2
	multiple := num * i
	for {
		if multiple > *n {
			break
		}

		markMap.Store(multiple, true)
		i++
		multiple = num * i
	}
}
