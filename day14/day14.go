package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	STONE string = "stone"
	SAND  string = "sand"
)

type (
	coordinate struct {
		x int
		y int
	}
	gridspace struct {
		filled bool
		with   string
	}
)

func (c coordinate) isZero() bool {
	return c.x == 0 && c.y == 0
}

func (c coordinate) updateBounds(left, right, maxy int) (int, int, int) {
	var l, r, m = left, right, maxy
	if c.x < left {
		l = c.x
	}
	if c.x > right {
		r = c.x
	}
	if c.y > maxy {
		m = c.y
	}
	return l, r, m
}

func (from coordinate) betweenCoords(to coordinate) []coordinate {
	var result = []coordinate{}
	if from.y == to.y {
		if from.x < to.x {
			for i := from.x; i <= to.x; i++ {
				result = append(result, coordinate{x: i, y: from.y})
			}
		} else {
			for i := to.x; i <= from.x; i++ {
				result = append(result, coordinate{x: i, y: from.y})
			}
		}
	} else {
		if from.y < to.y {
			for i := from.y; i <= to.y; i++ {
				result = append(result, coordinate{x: from.x, y: i})
			}
		} else {
			for i := to.y; i <= from.y; i++ {
				result = append(result, coordinate{x: from.x, y: i})
			}
		}
	}
	return result
}

func copyGrid(oldgrid [][]gridspace, xoffset int, maxy int) [][]gridspace {
	var newgrid [][]gridspace
	newlength := len(oldgrid[0]) + int(math.Abs(float64(xoffset)))
	for i := 0; i <= maxy; i++ {
		var newline = make([]gridspace, newlength)
		if i < len(oldgrid) {
			var x int
			if xoffset < 0 {
				x = 0 - xoffset
			}
			for j := 0; j < len(oldgrid[0]); j++ {
				newline[x] = oldgrid[i][j]
				x++
			}
		}
		newgrid = append(newgrid, newline)
	}
	return newgrid
}

func printGrid(grid [][]gridspace) {
	for line, gridline := range grid {
		print(strconv.Itoa(line))
		if line < 10 {
			print("  ")
		} else {
			print(" ")
		}
		for _, gs := range gridline {
			switch gs.filled {
			case true:
				switch gs.with {
				case STONE:
					print("#")
				case SAND:
					print("o")
				}
			case false:
				print(".")
			}
		}
		println()
	}
}

// warning! Everything is upside-down!
func main() {
	var grid [][]gridspace
	var left, right, maxy int
	var previouscoordinate = coordinate{}
	var nograins int

	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := reader.Text()
		var startcoordinate = coordinate{}
		for {
			if parts := regexp.MustCompile(`^(\d+),(\d+)`).FindStringSubmatch(text); parts != nil {
				var thiscoordinate = coordinate{}
				var err error
				thiscoordinate.x, err = strconv.Atoi(parts[1])
				if err != nil {
					panic(err)
				}
				thiscoordinate.y, err = strconv.Atoi(parts[2])
				if err != nil {
					panic(err)
				}
				if previouscoordinate.isZero() {
					grid = make([][]gridspace, thiscoordinate.y+1)
					for i := 0; i <= thiscoordinate.y; i++ {
						grid[i] = make([]gridspace, 1)
					}
					grid[thiscoordinate.y][0] = gridspace{filled: true, with: STONE}
					left, right = thiscoordinate.x, thiscoordinate.x
					maxy = thiscoordinate.y
				} else {
					var xoffset int
					if left > thiscoordinate.x {
						xoffset = thiscoordinate.x - left
					} else if right < thiscoordinate.x {
						xoffset = thiscoordinate.x - right
					}
					left, right, maxy = thiscoordinate.updateBounds(left, right, maxy)
					grid = copyGrid(grid, xoffset, maxy)
					if !startcoordinate.isZero() {
						for _, c := range thiscoordinate.betweenCoords(startcoordinate) {
							grid[c.y][c.x-left].filled = true
							grid[c.y][c.x-left].with = STONE
						}
					}
				}
				startcoordinate = thiscoordinate
				text = strings.TrimLeft(text, parts[0])
				text = strings.TrimLeft(text, " -> ")
				previouscoordinate = thiscoordinate
			} else {
				break
			}
		}
	}
	sandpoint := 500
	//part 2
	grid = copyGrid(grid, 0-left, maxy)
	grid = copyGrid(grid, 500, maxy)
	left = 0
	right += 500
	grid = append(grid, make([]gridspace, right-left+1))
	grid = append(grid, make([]gridspace, right-left+1))
	for i := 0; i < right-left+1; i++ {
		grid[maxy+2][i].filled = true
		grid[maxy+2][i].with = STONE
	}

	// bring on the sand!!
	for {
		var sandx, sandy = sandpoint, 0
		for i := 0; i <= maxy+1; i++ {
			if grid[i][sandx-left].filled {
				if sandx > left && grid[i][sandx-left-1].filled {
					if sandx < right && grid[i][sandx-left+1].filled {
						break
					} else {
						sandx = sandx + 1
					}
				} else {
					sandx = sandx - 1
				}
			}
			sandy = i
		}
		//part 1
		/*if sandy == maxy {
			break
		}*/
		if sandy == 0 {
			nograins++
			grid[sandy][sandx-left].filled = true
			grid[sandy][sandx-left].with = SAND
			break
		}

		nograins++
		grid[sandy][sandx-left].filled = true
		grid[sandy][sandx-left].with = SAND
	}
	println("left is " + strconv.Itoa(left))
	println("right is " + strconv.Itoa(right))
	println("number of sand grains is " + strconv.Itoa(nograins))
	//printGrid(grid)
}
