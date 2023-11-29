package solutions

import (
	"fmt"
	"time"
)

type Solution6 struct {
}

func (s Solution6) Solve() {
	fmt.Println("\n===== Solution 6 =====")

	for i := 0; i < 10; i++ {
		go fmt.Println("Id: ", i)
	}
	time.Sleep(5 * time.Second)
}
