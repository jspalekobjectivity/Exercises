package solutions

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Solution8 struct {
}

func (s Solution8) Solve() {
	fmt.Println("\n===== Solution 8 =====")

	dogs := []Dog{
		{"Włodek", 3},
		{"Okruszek", 4},
		{"Alojzy", 2},
		{"Burek", 11},
		{"Azor", 7},
	}

	fmt.Println("Dogs before birthday: ", dogs)

	var waitGroup sync.WaitGroup

	for i := 0; i < len(dogs); i++ {
		waitGroup.Add(1)
		i := i

		go func() {
			defer waitGroup.Done()
			birthdayWorker(&dogs, i)
		}()
	}

	waitGroup.Wait()

	fmt.Println("Dogs after birthday: ", dogs)
}

func birthdayWorker(dogsSlice *[]Dog, index int) {
	timeOffsetMin := 200
	timeOffsetMax := 5000
	randomTimeOffset := rand.Intn((timeOffsetMax - timeOffsetMin + 1) - timeOffsetMin)

	fmt.Printf("Worker %d starting with offset %d \n", index, randomTimeOffset)

	// Simulate some heavier calculations for dramatic effect
	time.Sleep(time.Duration(randomTimeOffset) * time.Millisecond)
	(*dogsSlice)[index].Age += 1

	fmt.Printf("Worker %d done\n", index)
}
