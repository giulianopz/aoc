package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	rounds = 20
)

type Operation string

const (
	add   Operation = "+"
	times Operation = "*"
	old   string    = "old"
)

type monkey struct {
	items             []int
	inspectedItems    int
	operation         Operation
	operand           string
	divisor           int
	nextMonkeyIfTrue  int
	nextMonkeyIfFalse int
}

func PartOne(input []string) string {

	numRegex, _ := regexp.Compile(`(\d+)`)
	operationRegex, _ := regexp.Compile(`[+*]\s(\d+|old)`)

	var monkeyBusiness int

	monkeys := make([]*monkey, 0)
	var m *monkey
	for _, line := range input {

		if strings.Contains(line, "Monkey") {
			m = new(monkey)
			monkeys = append(monkeys, m)
		} else if strings.Contains(line, "Starting") {
			matches := numRegex.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				item, err := strconv.Atoi(match[1])
				if err != nil {
					log.Fatal(err)
				}
				m.items = append(m.items, item)
			}
		} else if strings.Contains(line, "Operation") {
			matches := operationRegex.FindAllStringSubmatch(line, -1)
			m.operand = matches[0][1]
			if strings.Contains(line, string(add)) {
				m.operation = add
			} else {
				m.operation = times
			}
		} else if strings.Contains(line, "divisible") {
			matches := numRegex.FindAllStringSubmatch(line, -1)
			d, err := strconv.Atoi(matches[0][1])
			if err != nil {
				log.Fatal(err)
			}
			m.divisor = d
		} else if strings.Contains(line, "If true") {
			matches := numRegex.FindAllStringSubmatch(line, -1)
			monkeyId, err := strconv.Atoi(matches[0][1])
			if err != nil {
				log.Fatal(err)
			}
			m.nextMonkeyIfTrue = monkeyId
		} else if strings.Contains(line, "If false") {
			matches := numRegex.FindAllStringSubmatch(line, -1)
			monkeyId, err := strconv.Atoi(matches[0][1])
			if err != nil {
				log.Fatal(err)
			}
			m.nextMonkeyIfFalse = monkeyId
		}
	}

	for i := 0; i < rounds; i++ {

		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]
			var worryLevel int
			for _, i := range m.items {

				m.inspectedItems += 1
				var op int
				if m.operand == old {
					op = i
				} else {
					num, err := strconv.Atoi(m.operand)
					if err != nil {
						log.Fatal(err)
					}
					op = num
				}

				switch m.operation {
				case add:
					worryLevel = i + op
				case times:
					worryLevel = i * op
				}

				worryLevel /= 3

				if (worryLevel % m.divisor) == 0 {
					monkeys[m.nextMonkeyIfTrue].items = append(monkeys[m.nextMonkeyIfTrue].items, worryLevel)
				} else {
					monkeys[m.nextMonkeyIfFalse].items = append(monkeys[m.nextMonkeyIfFalse].items, worryLevel)
				}
			}
			m.items = make([]int, 0)
		}
	}

	inspectedItemsPerMonkey := make([]int, 0)
	for _, m := range monkeys {
		inspectedItemsPerMonkey = append(inspectedItemsPerMonkey, m.inspectedItems)
	}

	sort.Slice(inspectedItemsPerMonkey, func(i, j int) bool {
		return inspectedItemsPerMonkey[i] > inspectedItemsPerMonkey[j]
	})

	monkeyBusiness = inspectedItemsPerMonkey[0] * inspectedItemsPerMonkey[1]

	return fmt.Sprintf("%d", monkeyBusiness)
}

func PartTwo(input []string) string {
	return ""
}
