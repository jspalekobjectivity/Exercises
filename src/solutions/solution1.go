package solutions

import "fmt"

//no need to create empty struct
type Solution1 struct {
}

// could be just a function
// should be test
func (s Solution1) Solve() {
	fmt.Println("===== Solution 1 =====")

	var slice = []int{1, 2, 3, 4, 5}

	fmt.Println("Three arguments: ", sum(1, 2, 3))
	fmt.Println("Seven arguments: ", sum(4, 5, 6, 7, 8, 9, 0))
	fmt.Println("A slice: ", sum(slice...))
}

func sum(nums ...int) int {
	var sum int

	for _, value := range nums {
		sum += value
	}

	return sum
}
