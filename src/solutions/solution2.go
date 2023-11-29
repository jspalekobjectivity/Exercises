package solutions

import (
	"errors"
	"fmt"
)

type Solution2 struct {
}

func (s Solution2) Solve() {
	fmt.Println("\n===== Solution 2 =====")

	slice1 := []int{1, 2, 3}
	var slice2 []int

	fmt.Println("Slice before appending: ", slice1, "Length: ", len(slice1))

	if err := appendSlice(&slice1, 4); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Slice after appending: ", slice1, "Length: ", len(slice1))

	fmt.Println("Testing for nil value slice: ")

	if err := appendSlice(&slice2, 4); err != nil {
		fmt.Println(err)
	}
}

func appendSlice(slice *[]int, value int) error {
	if *slice == nil {
		return errors.New("invalid slice pointer")
	}

	if len(*slice)%2 == 1 {
		*slice = append(*slice, value)
	}

	return nil
}
