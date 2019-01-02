package main

import (
	"strconv"
	"strings"
)

func computeResultDay3_1() int {

	var claims []string
	var err error
	claims, err = readLines("/Users/jonathangildevall/Documents/Programming/AdventOfCodeGO/input/day3.txt")
	if err != nil {
		return -1
	}
	fabric := make([][]int, 1000)
	for i := 0; i < len(fabric); i++ {
		temp := make([]int, 1000)
		fabric[i] = temp
	}
	overlapping := 0
	for i := 0; i < len(claims); i++ {
		indexAt := strings.Index(claims[i], "@")
		indexComma := strings.Index(claims[i], ",")
		indexColon := strings.Index(claims[i], ":")
		indexX := strings.Index(claims[i], "x")

		id, _ := strconv.Atoi(claims[i][1 : indexAt-1])
		x, _ := strconv.Atoi(claims[i][indexAt+2 : indexComma])
		y, _ := strconv.Atoi(claims[i][indexComma+1 : indexColon])
		width, _ := strconv.Atoi(claims[i][indexColon+2 : indexX])
		heigth, _ := strconv.Atoi(claims[i][indexX+1 : len(claims[i])])
		for tX := x; tX < x+width; tX++ {
			for tY := y; tY < y+heigth; tY++ {
				if fabric[tX][tY] == 0 {
					fabric[tX][tY] = id
				} else if fabric[tX][tY] == -1 {
				} else {
					fabric[tX][tY] = -1
					overlapping++
				}
			}
		}
	}
	return overlapping
}

func computeResultDay3_2() int {

	var claims []string
	var err error
	claims, err = readLines("/Users/jonathangildevall/Documents/Programming/AdventOfCodeGO/input/day3.txt")
	if err != nil {
		return -1
	}
	fabric := make([][]string, 1000)
	for i := 0; i < len(fabric); i++ {
		temp := make([]string, 1000)
		fabric[i] = temp
	}
	overlapped := make([]bool, len(claims))
	for i := 0; i < len(claims); i++ {
		indexAt := strings.Index(claims[i], "@")
		indexComma := strings.Index(claims[i], ",")
		indexColon := strings.Index(claims[i], ":")
		indexX := strings.Index(claims[i], "x")

		id := claims[i][1 : indexAt-1]
		x, _ := strconv.Atoi(claims[i][indexAt+2 : indexComma])
		y, _ := strconv.Atoi(claims[i][indexComma+1 : indexColon])
		width, _ := strconv.Atoi(claims[i][indexColon+2 : indexX])
		heigth, _ := strconv.Atoi(claims[i][indexX+1 : len(claims[i])])
		for tX := x; tX < x+width; tX++ {
			for tY := y; tY < y+heigth; tY++ {
				if fabric[tX][tY] == "" {
					fabric[tX][tY] = id
				} else {
					//When 0 is overlapped
					if len(strings.Split(fabric[tX][tY], " ")) == 1 {
						overlappedId, _ := strconv.Atoi(fabric[tX][tY])
						overlapped[overlappedId-1] = true
					}
					intId, _ := strconv.Atoi(id)
					overlapped[intId-1] = true
					fabric[tX][tY] = fabric[tX][tY] + " " + id

				}
			}
		}
	}
	for i := 0; i < len(overlapped); i++ {
		if !overlapped[i] {
			return i + 1
		}
	}
	return 0
}
