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
	close(channel)
}

func Consumer(channel <-chan int) {
	number := <-channel
	fmt.Println("Consumed number: ", number)
}
