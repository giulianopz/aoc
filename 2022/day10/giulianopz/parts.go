package main

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	noop = "noop"
	add  = "addx"
)

func PartOne(input []string) string {

	addRegex, _ := regexp.Compile(add + "\\s(-?\\d+)")

	offset := 40
	var cycleNum int
	var lastObserved int

	X := 1
	var signalStrengt int

	compute := func(cycleNum int) {
		if cycleNum == 20 || cycleNum == (lastObserved+offset) {
			lastObserved = cycleNum
			signalStrengt += (X * cycleNum)
		}
	}

	for _, line := range input {

		if line == noop {
			cycleNum++
			compute(cycleNum)

		} else if addRegex.MatchString(line) {
			cycleNum++
			compute(cycleNum)

			cycleNum++
			compute(cycleNum)

			matches := addRegex.FindAllStringSubmatch(line, -1)
			value, _ := strconv.Atoi(matches[0][1])
			X += value
		}
	}
	return fmt.Sprintf("%d", signalStrengt)
}

func interesting(cycleNum int) bool {
	return cycleNum == 20 ||
		cycleNum == 60 || cycleNum == 100 || cycleNum == 140 || cycleNum == 180 || cycleNum == 220

}

func PartTwo(input []string) string {

	CRT := [6][40]string{}
	X := 1

	addRegex, _ := regexp.Compile(add + "\\s(-?\\d+)")

	draw := func(cycleNum, X int) {
		var i int
		if (cycleNum % 40) != 0 {
			i = cycleNum / 40
		} else {
			i = (cycleNum / 40) - 1
		}
		var j int
		if (cycleNum % 40) != 0 {
			j = cycleNum%40 - 1
		} else {
			j = 40 - 1
		}
		//fmt.Printf("drawing in array %d at position %d when cycle num is %d\n", i, j, cycleNum)

		if j == X || j == (X-1) || j == (X+1) {
			CRT[i][j] = "#"
		} else {
			CRT[i][j] = "."
		}
	}

	var cycleNum int
	for _, line := range input {

		if line == noop {
			cycleNum++
			draw(cycleNum, X)

		} else if addRegex.MatchString(line) {
			cycleNum++
			draw(cycleNum, X)

			cycleNum++
			draw(cycleNum, X)

			matches := addRegex.FindAllStringSubmatch(line, -1)
			value, _ := strconv.Atoi(matches[0][1])
			X += value
		}
	}

	image := "\n"
	for _, pixels := range CRT {
		for _, p := range pixels {
			image += p
		}
		image += "\n"
	}
	return image
}
