package main

import (
	"fmt"
)

const (
	currentPos    = 'S'
	bestSignalPos = 'E'
)

type node struct {
	elevation rune
	char      string
	row       int
	col       int
	visited   bool
	steps     int
	previous  *node
}

type path []*node

type grid [][]*node

func PartOne(input []string) string {

	var source, destination *node
	g := make(grid, len(input))
	for i, line := range input {
		g[i] = make([]*node, len(line))
		for j, c := range line {
			g[i][j] = &node{elevation: c, row: i, col: j, char: string(c)}
			if c == currentPos {
				source = g[i][j]
			} else if c == bestSignalPos {
				destination = g[i][j]
			}
		}
	}

	steps := g.bfs(source, destination)

	return fmt.Sprintf("%d", steps)
}

func (g *grid) outside(r, c int) bool {
	return r < 0 || c < 0 || r >= len(*g) || c >= len(([][]*node)(*g)[0])
}

func (g *grid) neighbours(current *node) []*node {

	neigh := make([]*node, 0)
	//up
	if !g.outside(current.row-1, current.col) {
		up := ([][]*node)(*g)[current.row-1][current.col]
		if fit(current.elevation, up.elevation) {
			neigh = append(neigh, up)
		}
	}

	//down
	if !g.outside(current.row+1, current.col) {
		down := ([][]*node)(*g)[current.row+1][current.col]
		if fit(current.elevation, down.elevation) {
			neigh = append(neigh, down)
		}
	}

	//right
	if !g.outside(current.row, current.col+1) {
		right := ([][]*node)(*g)[current.row][current.col+1]
		if fit(current.elevation, right.elevation) {
			neigh = append(neigh, right)
		}
	}

	//left
	if !g.outside(current.row, current.col-1) {
		left := ([][]*node)(*g)[current.row][current.col-1]
		if fit(current.elevation, left.elevation) {
			neigh = append(neigh, left)
		}
	}

	return neigh
}

func fit(source, destination rune) bool {
	if source == currentPos {
		source = 'a'
	}
	if destination == bestSignalPos {
		destination = 'z'
	}

	if destination <= source || destination-source == 1 {
		return true
	}

	return false
}

func (g *grid) bfs(source, destination *node) int {

	source.visited = true
	source.steps = 0
	source.previous = nil

	q := make([]*node, 0)
	q = append(q, source)

	for len(q) != 0 {
		current := q[0]
		q = q[1:]
		for _, neigh := range g.neighbours(current) {
			if !neigh.visited {

				neigh.visited = true
				neigh.steps = current.steps + 1
				neigh.previous = current

				if neigh.elevation == destination.elevation {
					g.trace(neigh)
					return neigh.steps
				}
				q = append(q, neigh)
			}
		}
	}
	return 0
}

func (g *grid) trace(destination *node) {

	curr := destination
	for curr.previous != nil {
		fmt.Print(string(curr.elevation) + "<")
		curr = curr.previous
	}
}

func PartTwo(input []string) string {
	return ""
}
