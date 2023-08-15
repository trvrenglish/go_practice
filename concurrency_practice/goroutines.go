package main

import (
	"fmt"
	"time"
)

func main() {
	Exec()
}

// Count function never finishes, so it will only ever count sheep
func Goroutine() {
	c := make(chan string)
	go Count("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func Count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}

func Channel() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	var msg string

	for i := 0; i < 2; i++ {
		msg = <-c
		fmt.Println(msg)
	}
}

func Select() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func Exec() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go Worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func Worker(jobs chan int, results chan int) {
	for n := range jobs {
		results <- Fib(n)
	}
}

func Fib(n int) int {
	if n <= 1 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}
