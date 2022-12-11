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
	return ""
}
