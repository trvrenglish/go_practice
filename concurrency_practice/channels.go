package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func unbuffered() {
	unbuffered := make(chan int)

	// 1) blocks because no data
	a := <-unbuffered
	fmt.Println(a)
	// 2) blocks because no receiver
	unbuffered <- 1

	// 3) Synchronises
	go func() { <-unbuffered }()
	unbuffered <- 1
}

func buffered() {
	buffered := make(chan int, 1)

	// 4) Blocks because no data
	b := <-buffered
	fmt.Println(b)

	// 5) Success. There is room in the channel
	buffered <- 1

	// 6) Blocks because no room in the channel
	buffered <- 2
}

func closing() {
	c := make(chan int)
	close(c)
	fmt.Println(<-c) // receive and print
}

func TryReceive(c <-chan int) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true

	default: // Proceed when c is blocking
		return 0, true, false
	}
}

func TryReceiveWithoutTimeout(c <-chan int, duration time.Duration) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true

	case <-time.After(duration): // Blocks until duration is over
		return 0, true, false
	}
}

func Fanout(In <-chan int, OutA, OutB chan int) {
	for data := range In { // Receive until closed
		select { // Send to first non-blocking channel
		case OutA <- data:
		case OutB <- data:
		}
	}
}

func Turnout(InA, InB <-chan int, OutA, OutB chan int) {
	var data int
	var more bool

	for { // Receive until closed
		select { // Recieve from first non-blocking channel
		case data, more = <-InA:
		case data, more = <-InB:
		}
		if !more {
			return
		}
		select { // Send to first non-blocking
		case OutA <- data:
		case OutB <- data:
		}
	}
}

func TurnoutWithQuitChannel(Quit <-chan int, InA, InB, OutA, OutB chan int) {
	for { // Receive until closed
		select { // Receive from first non-blocking channel
		case _ = <-InA:
		case _ = <-InB:
		case <-Quit:
			close(InA)
			close(InB)

			Fanout(InA, OutA, OutB)
			Fanout(InB, OutA, OutB)
			return
		}
	}
}

type Spinlock struct {
	state *int32
}

const free = int32(0)

func (l *Spinlock) Lock() {
	for !atomic.CompareAndSwapInt32(l.state, free, 42) {
		runtime.Gosched()
	}
}

func (l *Spinlock) Unlock() {
	atomic.StoreInt32(l.state, free)
}

type TicketStore struct {
	ticket *uint64
	done   *uint64
	slots  []string // Imagine to be infinite
}

func (ts *TicketStore) Put(s string) {
	t := atomic.AddUint64(ts.ticket, 1) - 1 // Draw a ticket
	ts.slots[t] = s                         // Store your data
	for !atomic.CompareAndSwapUint64(ts.done, t, t+1) {
		runtime.Gosched()
	}
}

func (ts *TicketStore) GetDone() []string {
	return ts.slots[:atomic.LoadUint64(ts.done)+1]
}
