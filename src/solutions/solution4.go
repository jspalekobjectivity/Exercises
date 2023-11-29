package solutions

import "fmt"

type Solution4 struct {
}

//should be test
func (s Solution4) Solve() {
	fmt.Println("\n===== Solution 4 =====")

	fmt.Printf("Are 'below' and 'elbow' an anagram: %v \n", AreAnagram("below", "elbow"))

	fmt.Printf("Are 'stereo' and 'rodeos' an anagram: %v \n", AreAnagram("stereo", "rodeos"))

	fmt.Printf("Are 'Tatooine' and 'Mustafar' an anagram: %v \n", AreAnagram("Tatooine", "Mustafar"))

	fmt.Printf("Are '' and ' ' an anagram: %v \n", AreAnagram("", " "))

}

func AreAnagram(source, target string) bool {
	if len(source) != len(target) {
		return false
	}

	sourceMap := make(map[rune]int)
	targetMap := make(map[rune]int)

	for _, letter := range source {
		sourceMap[letter]++
	}

	for _, letter := range target {
		targetMap[letter]++
	}

	// see maps package to compere maps
	for letter, sourceCount := range sourceMap {
		if targetCount, ok := targetMap[letter]; !ok || sourceCount != targetCount {
			return false
		}
	}

	return true
}
