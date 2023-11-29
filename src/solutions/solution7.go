package solutions

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Solution7 struct {
}

func (s Solution7) Solve() {
	fmt.Println("\n===== Solution 7 =====")

	var waitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		i := i

		go func() {
			defer waitGroup.Done()
			worker(i)
		}()
	}

	waitGroup.Wait()
}

func worker(id int) {
	timeOffsetMin := 200
	timeOffsetMax := 2000
	randomTimeOffset := rand.Intn((timeOffsetMax - timeOffsetMin + 1) - timeOffsetMin)

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(randomTimeOffset) * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}
