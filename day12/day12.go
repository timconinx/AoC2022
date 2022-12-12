package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type (
	gridspace struct {
		letter  string
		painted bool
	}
	coordinate struct {
		x int
		y int
	}
)

var letters = make(map[string]int)

func (gs gridspace) canFlowTo(other gridspace) bool {
	if other.painted {
		return false
	}
	if letters[other.letter] > letters[gs.letter]+1 {
		return false
	}
	return true
}

func (c coordinate) getNeighbours(maxx, maxy int) []coordinate {
	result := []coordinate{}
	// left
	if c.x > 0 {
		result = append(result, coordinate{x: c.x - 1, y: c.y})
	}
	// right
	if c.x < maxx {
		result = append(result, coordinate{x: c.x + 1, y: c.y})
	}
	// up
	if c.y > 0 {
		result = append(result, coordinate{x: c.x, y: c.y - 1})
	}
	// down
	if c.y < maxy {
		result = append(result, coordinate{x: c.x, y: c.y + 1})
	}
	return result
}

func (c coordinate) equals(other coordinate) bool {
	if c.x == other.x && c.y == other.y {
		return true
	}
	return false
}

func (c coordinate) String() string {
	return "(" + strconv.Itoa(c.x) + "," + strconv.Itoa(c.y) + ")"
}

func main() {
	for i, l := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		letters[l] = i
	}
	var grid = [][]gridspace{}
	var streams [][]coordinate
	var starts = []coordinate{}
	var end coordinate
	var maxx, maxy int
	var path []coordinate

	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := strings.Split(reader.Text(), "")
		if len(grid) == 0 {
			maxx = len(text) - 1
		}
		gridline := []gridspace{}
		for i, l := range text {
			if l == "S" || l == "a" {
				starts = append(starts, coordinate{x: i, y: maxy})
				l = "a"
			}
			if l == "E" {
				end = coordinate{x: i, y: maxy}
				l = "z"
			}
			gridline = append(gridline, gridspace{letter: l, painted: false})
		}
		grid = append(grid, gridline)
		maxy++
	}
	maxy--

	min := 500
	for _, start := range starts { // extra for-loop for part b
		newgrid := [][]gridspace{}
		for _, gridline := range grid {
			newgridline := []gridspace{}
			for _, gs := range gridline {
				newgridline = append(newgridline, gridspace{letter: gs.letter, painted: false})
			}
			newgrid = append(newgrid, newgridline)
		}
		grid = newgrid
		path = []coordinate{}
		streams = [][]coordinate{}
		streams = append(streams, []coordinate{start})
		var pathFound bool
		var noNeighbours bool
		for {
			if pathFound || noNeighbours {
				break
			}
			newStreams := [][]coordinate{}
			noNeighbours = true
			for _, stream := range streams {
				lc := stream[len(stream)-1]
				neighbours := lc.getNeighbours(maxx, maxy)
				noNeighbours = noNeighbours && (len(neighbours) == 0)
				for _, n := range neighbours {
					newStream := make([]coordinate, len(stream))
					copy(newStream, stream)
					if grid[lc.y][lc.x].canFlowTo(grid[n.y][n.x]) {
						grid[n.y][n.x].painted = true
						newStream = append(newStream, n)
						newStreams = append(newStreams, newStream)
						if n.equals(end) {
							pathFound = true
							path = newStream
							break
						}
					}
				}
				if pathFound {
					break
				}
			}
			streams = newStreams
		}
		if pathFound {
			if len(path)-1 < min {
				min = len(path) - 1
			}
		}
	}
	println("minimal number of steps is " + strconv.Itoa(min))
}
