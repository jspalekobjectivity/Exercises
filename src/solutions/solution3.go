package solutions

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

type Solution3 struct {
}

func (s Solution3) Solve() {
	fmt.Println("\n===== Solution 3 =====")

	dogs := []Dog{
		{"Włodek", 3},
		{"Okruszek", 4},
		{"Alojzy", 2},
	}

	dogsPtrs := []*Dog{
		{"Włodek", 3},
		{"Okruszek", 4},
		{"Alojzy", 2},
	}

	fmt.Println("Dogs before any changes: ", dogs)

	dogs = makeBirthday(dogs)
	fmt.Println("Dogs after birthday: ", dogs)

	dogs = sortDogByName(dogs)
	fmt.Println("Dogs sorted by name: ", dogs)

	sortDogPtrByName(dogsPtrs)
	fmt.Print("Dogs sorted by name using pointers:")
	for _, dog := range dogsPtrs {
		fmt.Printf("%v ", dog)
	}
	fmt.Println()

	dogs = sortDogByAge(dogs)
	fmt.Print("Dogs sorted by age: ", dogs)
}

type Dog struct {
	Name string
	Age  uint32
}

type ageSortedDogs []Dog

func (d ageSortedDogs) Len() int {
	return len(d)
}

func (d ageSortedDogs) Less(i, j int) bool {
	return d[i].Age < d[j].Age
}

func (d ageSortedDogs) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func makeBirthday(dogs []Dog) []Dog {
	for i := 0; i < len(dogs); i++ {
		dogs[i].Age += 1
	}

	return dogs
}

func sortDogByAge(dogs []Dog) []Dog {
	sort.Sort(ageSortedDogs(dogs))

	return dogs
}

func sortDogByName(dogs []Dog) []Dog {
	dogComparator := func(a, b Dog) int {
		return cmp.Compare(a.Name, b.Name)
	}

	slices.SortFunc(dogs, dogComparator)

	return dogs
}

func sortDogPtrByName(dogs []*Dog) {
	dogComparator := func(a, b *Dog) int {
		return cmp.Compare(a.Name, b.Name)
	}

	slices.SortFunc(dogs, dogComparator)
}
