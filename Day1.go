package main

func computeResultDay1_1() int {
	var input []int
	var err error
	input, err = readIntegers("/Users/jonathangildevall/Documents/Programming/AdventOfCodeGO/input/day1.txt")
	if err != nil {
		return -1
	}
	frequency := 0
	for i := 0; i < len(input); i++ {
		frequency += input[i]
	}
	return frequency
}

func computeResultDay1_2() int {
	var input []int
	var err error
	input, err = readIntegers("/Users/jonathangildevall/Documents/Programming/AdventOfCodeGO/input/day1.txt")
	if err != nil {
		return -1
	}
	var frequencies map[int]struct{}
	frequencies = make(map[int]struct{})
	frequency := 0

	for {
		for i := 0; i < len(input); i++ {
			frequency += input[i]
			if _, ok := frequencies[frequency]; ok {
				return frequency
			} else {
				frequencies[frequency] = struct{}{}
			}
		}
	}
}
