package solutions

import (
	"fmt"
	"sync"
)

type Solution10 struct {
}

func (s Solution10) Solve() {
	numberPipe := make(chan int)
	textPipe := make(chan string)

	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	go func() {
		defer waitGroup.Done()
		TextProducer(textPipe, []string{"A", "B", "C", "D", "E"})
	}()

	go func() {
		defer waitGroup.Done()
		NumberProducer(numberPipe, []int{1, 2, 3, 4, 5})
	}()
	var out <-chan interface{}

	go func() {
		defer waitGroup.Done()
		out = Merge(textPipe, numberPipe)
	}()

	waitGroup.Wait()

	var textResults []string
	var numberResults []int

	for data := range out {

		switch t := data.(type) {
		case int:
			castedValue, ok := data.(int)
			fmt.Println(t)
			if ok {
				numberResults = append(numberResults, castedValue)
			}
		case string:
			castedValue, ok := data.(string)

			if ok {
				textResults = append(textResults, castedValue)
			}
		}
	}
}

func NumberProducer(pipe chan<- int, dataToProduce []int) {
	for _, data := range dataToProduce {
		pipe <- data
	}

	close(pipe)
}

func TextProducer(pipe chan<- string, dataToProduce []string) {
	for _, data := range dataToProduce {
		pipe <- data
	}

	close(pipe)
}

func Merge(textPipe <-chan string, numberPipe <-chan int) <-chan interface{} {
	out := make(chan interface{})

	select {
	case number := <-numberPipe:
		out <- number
	case text := <-textPipe:
		out <- text
	}

	return out
}
