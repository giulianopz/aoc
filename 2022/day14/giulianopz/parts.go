package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func asPoint(xy string) *point {
	splitted := strings.Split(xy, ",")
	x, _ := strconv.Atoi(splitted[0])
	y, _ := strconv.Atoi(splitted[1])
	return &point{x: x, y: y}
}

const (
	sepS    = " -> "
	rockS   = "#"
	sourceS = "+"
	airS    = "."
	sandS   = "O"
)

var (
	s = &point{x: 500, y: 0}
)

func draw(grid [][]string, p *point, symbol string) {

	x := p.x
	if p.x > 400 {
		x -= 400
	}

	grid[p.y][x] = symbol
}

func blocked(grid [][]string, x, y int) bool {

	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return true
	}

	return grid[y][x] == rockS || grid[y][x] == sandS
}

func render(grid [][]string) {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Print("\n")
	}
}

func PartOne(input []string) string {

	grid := make([][]string, 170)
	for i := range grid {
		grid[i] = make([]string, 170)
		for j := range grid[i] {
			grid[i][j] = airS
		}
	}

	var lowest int
	for _, line := range input {

		coordinates := strings.Split(line, sepS)
		for i, j := 0, 1; i < len(coordinates) && j < len(coordinates); i, j = i+1, j+1 {

			f := asPoint(coordinates[i])
			s := asPoint(coordinates[j])

			if f.y > lowest {
				lowest = f.y
			}

			//draw along row
			if f.y == s.y {
				min := int(math.Min(float64(f.x), float64(s.x)))
				max := int(math.Max(float64(f.x), float64(s.x)))
				for x := min; x <= max; x++ {
					normalized := x - 400
					grid[f.y][normalized] = rockS
				}
			}

			//draw along column
			if f.x == s.x {
				min := int(math.Min(float64(f.y), float64(s.y)))
				max := int(math.Max(float64(f.y), float64(s.y)))
				for y := min; y <= max; y++ {
					normalized := f.x - 400
					grid[y][normalized] = rockS
				}
			}
		}
	}

	//render(grid)

	draw(grid, s, sourceS)

	var sandUnits int
	var sand *point
	var filled bool
	for !filled {

		if sand == nil {
			sandUnits++
			sand = &point{s.x - 400, s.y}
			draw(grid, sand, sandS)
		} else if !blocked(grid, sand.x, sand.y+1) {
			//down
			draw(grid, sand, airS)
			sand.y++
			draw(grid, sand, sandS)
		} else if !blocked(grid, sand.x-1, sand.y+1) {
			//left
			draw(grid, sand, airS)
			sand.x--
			sand.y++
			draw(grid, sand, sandS)
		} else if !blocked(grid, sand.x+1, sand.y+1) {
			// right
			draw(grid, sand, airS)
			sand.x++
			sand.y++
			draw(grid, sand, sandS)
		} else {
			sand = nil
		}
		if sand != nil && sand.y > lowest {
			//render(grid)
			filled = true
		}
	}

	return fmt.Sprintf("%d", sandUnits-1)
}

func PartTwo(input []string) string {
	grid := make([][]string, 170)
	for i := range grid {
		grid[i] = make([]string, 170)
		for j := range grid[i] {
			grid[i][j] = airS
		}
	}

	var floor int
	var lowest int
	for _, line := range input {

		coordinates := strings.Split(line, sepS)
		for i, j := 0, 1; i < len(coordinates) && j < len(coordinates); i, j = i+1, j+1 {

			f := asPoint(coordinates[i])
			s := asPoint(coordinates[j])

			if f.y > lowest {
				lowest = f.y
			}

			//draw along row
			if f.y == s.y {
				min := int(math.Min(float64(f.x), float64(s.x)))
				max := int(math.Max(float64(f.x), float64(s.x)))
				for x := min; x <= max; x++ {
					normalized := x - 400
					grid[f.y][normalized] = rockS
				}
			}

			//draw along column
			if f.x == s.x {
				min := int(math.Min(float64(f.y), float64(s.y)))
				max := int(math.Max(float64(f.y), float64(s.y)))
				for y := min; y <= max; y++ {
					normalized := f.x - 400
					grid[y][normalized] = rockS
				}
			}
		}
	}

	floor = lowest + 2
	for i := 0; i < len(grid[floor]); i++ {
		grid[floor][i] = rockS
	}
	//render(grid)

	draw(grid, s, sourceS)

	var sandUnits int
	var sand *point
	var filled bool
	for !filled {

		//render(grid)
		if sand == nil {
			sandUnits++
			sand = &point{s.x - 400, s.y}
			draw(grid, sand, sandS)
		} else if !blocked(grid, sand.x, sand.y+1) {
			//down
			draw(grid, sand, airS)
			sand.y++
			draw(grid, sand, sandS)
		} else if !blocked(grid, sand.x-1, sand.y+1) {
			//left
			draw(grid, sand, airS)
			sand.x--
			sand.y++
			draw(grid, sand, sandS)
		} else if !blocked(grid, sand.x+1, sand.y+1) {
			// right
			draw(grid, sand, airS)
			sand.x++
			sand.y++
			draw(grid, sand, sandS)
		} else {
			sand = nil
		}
		if grid[s.y][s.x-400] == sandS && grid[s.y+1][s.x-400] == sandS &&
			grid[s.y+1][s.x-401] == sandS && grid[s.y+1][s.x-399] == sandS {
			render(grid)
			filled = true
		}
	}

	return fmt.Sprintf("%d", sandUnits)
}
