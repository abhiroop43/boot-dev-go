package main

func concurrentFib(n int) []int {
	ch := make(chan int)
	retVal := make([]int, 0)
	go fibonacci(n, ch)
	for item := range ch {
		retVal = append(retVal, item)
	}

	return retVal
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
