package main

import s "Exercises/src/solutions"

func main() {
	// Created Solution interface for convenience when testing, comment out not needed solutions
	// Some shenanigans with goroutines will be seen when running all at once :)
	solutions := []s.Solution{
		s.Solution1{},
		s.Solution2{},
		s.Solution3{},
		s.Solution4{},
		s.Solution5{},
		s.Solution6{},
		s.Solution7{},
		s.Solution8{},
		s.Solution9{},
		s.Solution10{},
	}

	for _, solution := range solutions {
		solution.Solve()
	}
}
