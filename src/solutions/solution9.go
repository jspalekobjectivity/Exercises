package solutions

import (
	"fmt"
	"sync"
)

type Solution9 struct {
}

func (s Solution9) Solve() {
	var waitGroup sync.WaitGroup
	channel := make(chan int)

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		Producer(channel)
	}()

	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			Consumer(channel)
		}()
	}

	waitGroup.Wait()
}

func Producer(channel chan<- int) {
	for i := 1; i < 11; i++ {
		channel <- i
	}

	//should be in defer
	close(channel)
}

// consumer should read from channel in loop
// key part of exercise was how to stop reading from channel in loop after channel is closed
func Consumer(channel <-chan int) {
	number := <-channel
	fmt.Println("Consumed number: ", number)
}
