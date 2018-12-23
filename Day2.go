package main

import (
	"strings"
)

func computeResultDay2_1() int {
	var input []string
	var err error
	input, err = readLines("/Users/jonathangildevall/Documents/Programming/AdventOfCodeGO/input/day2.txt")
	if err != nil {
		return -1
	}
	nrOfPairs := 0
	nrOfTripples := 0
	for i := 0; i < len(input); i++ {
		temp := strings.Replace(input[i], input[i][:1], "", -1)
		tripples := 0
		pairs := 0
		for len(temp) > 0 {
			difference := len(input[i]) - len(temp)
			if difference == 3 {
				tripples++
			} else if difference == 2 {
				pairs++
			}
			input[i] = temp
			temp = strings.Replace(input[i], input[i][:1], "", -1)
		}
		if tripples > 0 {
			nrOfTripples++
		}
		if pairs > 0 {
			nrOfPairs++
		}
	}
	return nrOfPairs * nrOfTripples

}

func computeResultDay2_2() string {
	var input []string
	var err error
	input, err = readLines("/Users/jonathangildevall/Documents/Programming/AdventOfCodeGO/input/day2.txt")
	if err != nil {
		return ""
	}
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			difference := 0
			indexD := 0
			for k := 0; k < len(input[i])-1; k++ {
				if input[i][k:k+1] != input[j][k:k+1] {
					difference++
					indexD = k
					if difference > 1 {
						break
					}
				}
			}
			if difference == 1 {
				return input[i][:indexD] + input[i][indexD:]
			}
		}
	}
	return ""
}
